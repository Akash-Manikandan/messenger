package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorWhite  = "\033[37m"
)

type RequestLog struct {
	Timestamp    time.Time
	Method       string
	Path         string
	StatusCode   int
	Latency      time.Duration
	IP           string
	UserAgent    string
	ResponseSize int
}

// Logger middleware logs requests and responses asynchronously
func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Capture request details
		reqLog := RequestLog{
			Timestamp: start,
			Method:    c.Method(),
			Path:      c.Path(),
			IP:        c.IP(),
			UserAgent: c.Get("User-Agent"),
		}

		// Continue processing request
		err := c.Next()

		// Calculate response details
		reqLog.StatusCode = c.Response().StatusCode()
		reqLog.Latency = time.Since(start)
		reqLog.ResponseSize = len(c.Response().Body())

		// Log asynchronously to avoid blocking
		go logRequest(reqLog)

		return err
	}
}

func logRequest(r RequestLog) {
	// Color based on status code
	statusColor := getStatusColor(r.StatusCode)
	methodColor := getMethodColor(r.Method)
	latencyColor := getLatencyColor(r.Latency)

	fmt.Printf("%s[%s]%s %s%s%s %s%s%s | Status: %s%d%s | Duration: %s%v%s | IP: %s%s%s | Size: %s%d bytes%s\n",
		colorCyan, r.Timestamp.Format("2006-01-02 15:04:05"), colorReset,
		methodColor, r.Method, colorReset,
		colorWhite, r.Path, colorReset,
		statusColor, r.StatusCode, colorReset,
		latencyColor, r.Latency, colorReset,
		colorBlue, r.IP, colorReset,
		colorPurple, r.ResponseSize, colorReset,
	)
}

func getStatusColor(status int) string {
	switch {
	case status >= 200 && status < 300:
		return colorGreen
	case status >= 300 && status < 400:
		return colorCyan
	case status >= 400 && status < 500:
		return colorYellow
	case status >= 500:
		return colorRed
	default:
		return colorWhite
	}
}

func getMethodColor(method string) string {
	switch method {
	case "GET":
		return colorBlue
	case "POST":
		return colorGreen
	case "PUT":
		return colorYellow
	case "DELETE":
		return colorRed
	case "PATCH":
		return colorPurple
	default:
		return colorWhite
	}
}

func getLatencyColor(latency time.Duration) string {
	switch {
	case latency < 100*time.Millisecond:
		return colorGreen
	case latency < 500*time.Millisecond:
		return colorYellow
	default:
		return colorRed
	}
}
