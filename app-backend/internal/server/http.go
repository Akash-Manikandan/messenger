package server

import (
	"log"

	"github.com/Akash-Manikandan/app-backend/internal/config"
	"github.com/Akash-Manikandan/app-backend/internal/middleware"
	"github.com/Akash-Manikandan/app-backend/internal/modules/user"
	"github.com/Akash-Manikandan/app-backend/internal/registry"
	"github.com/Akash-Manikandan/app-backend/pkg/database"
	redisClient "github.com/Akash-Manikandan/app-backend/pkg/redis"
	"github.com/gofiber/fiber/v2"
)

func StartHTTP() {
	cfg := config.Load()

	// Initialize database
	db, err := database.InitDB(cfg.POSTGRES_URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize Redis
	redis, err := redisClient.InitRedis(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Run migrations if enabled
	if cfg.MIGRATE_DB {
		log.Println("Running database migrations...")
		if err := database.AutoMigrate(db); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		log.Println("Migrations completed successfully")
	} else {
		log.Println("Database migrations disabled (MIGRATE_DB=false)")
	}

	app := fiber.New()

	// Apply logger middleware globally
	app.Use(middleware.Logger())

	// Register routes
	registry.Load(app, db)

	// Register user module routes (needs DB and Redis)
	user.RegisterRoutes(app, db, redis)

	addr := ":" + cfg.HTTPPort
	log.Printf("HTTP server running on %s\n", addr)
	log.Fatal(app.Listen(addr))
}
