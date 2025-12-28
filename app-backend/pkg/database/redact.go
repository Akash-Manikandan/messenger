package database

import (
	"fmt"
	"regexp"
)

// RedactSQL redacts sensitive fields in SQL queries for logging
func RedactSQL(sql string) string {
	// Match INSERT statements with column names and values
	insertRegex := regexp.MustCompile(`(?s)INSERT INTO "[^"]*" \(([^)]+)\) VALUES \(([^)]+)\)`)
	sql = insertRegex.ReplaceAllStringFunc(sql, func(match string) string {
		parts := insertRegex.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}

		// Get column names
		colsPart := parts[1]
		valuesPart := parts[2]

		// Extract column names
		colRegex := regexp.MustCompile(`"([^"]+)"`)
		cols := colRegex.FindAllStringSubmatch(colsPart, -1)

		// Extract values (handle quoted strings properly)
		valRegex := regexp.MustCompile(`'([^']*)'|NULL|true|false|[0-9.]+`)
		vals := valRegex.FindAllString(valuesPart, -1)

		// Build map of column positions
		for i, col := range cols {
			if len(col) < 2 || i >= len(vals) {
				continue
			}

			colName := col[1]
			switch colName {
			case "password", "salt":
				// Replace the value at this position
				oldVal := vals[i]
				newVal := "'***REDACTED***'"
				valuesPart = regexp.MustCompile(regexp.QuoteMeta(oldVal)).ReplaceAllString(valuesPart, newVal)
			case "email":
				// Partial redaction for email
				oldVal := vals[i]
				emailMatch := regexp.MustCompile(`'([^@]{0,2})[^@]*@[^']*'`).FindStringSubmatch(oldVal)
				if len(emailMatch) > 1 {
					newVal := "'" + emailMatch[1] + "***@***'"
					valuesPart = regexp.MustCompile(regexp.QuoteMeta(oldVal)).ReplaceAllString(valuesPart, newVal)
				}
			}
		}

		return fmt.Sprintf("INSERT INTO \"users\" (%s) VALUES (%s)", colsPart, valuesPart)
	})

	// Redact UPDATE statements
	sql = regexp.MustCompile(`"password"='[^']*'`).ReplaceAllString(sql, `"password"='***REDACTED***'`)
	sql = regexp.MustCompile(`"salt"='[^']*'`).ReplaceAllString(sql, `"salt"='***REDACTED***'`)
	sql = regexp.MustCompile(`"email"='([^@]{0,2})[^@]*@[^']*'`).ReplaceAllString(sql, `"email"='$1***@***'`)

	return sql
}
