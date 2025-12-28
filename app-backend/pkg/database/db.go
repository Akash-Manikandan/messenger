package database

import (
	"context"
	"log"
	"time"

	"github.com/Akash-Manikandan/app-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// RedactLogger wraps GORM logger to redact sensitive data
type RedactLogger struct {
	logger.Interface
}

func (l RedactLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	sql = RedactSQL(sql)
	l.Interface.Trace(ctx, begin, func() (string, int64) {
		return sql, rows
	}, err)
}

// InitDB initializes the database connection
func InitDB(dsn string) (*gorm.DB, error) {
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: RedactLogger{
			Interface: logger.Default.LogMode(logger.Info),
		},
	})

	if err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return DB, nil
}

// AutoMigrate runs auto migration for all models
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		// Add other models here as you create them
	)
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
