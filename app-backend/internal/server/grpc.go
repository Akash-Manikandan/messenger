package server

import (
	"log"
	"net"

	"github.com/Akash-Manikandan/app-backend/internal/config"
	"github.com/Akash-Manikandan/app-backend/internal/middleware"
	"github.com/Akash-Manikandan/app-backend/internal/modules/user"
	"github.com/Akash-Manikandan/app-backend/internal/registry"
	"github.com/Akash-Manikandan/app-backend/pkg/database"
	redisClient "github.com/Akash-Manikandan/app-backend/pkg/redis"
	"google.golang.org/grpc"
)

func StartGRPC() {
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

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatal(err)
	}

	// Create gRPC server with logging interceptors
	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryServerInterceptor()),
		grpc.StreamInterceptor(middleware.StreamServerInterceptor()),
	)

	// Register services
	registry.LoadGRPC(server, db)
	user.RegisterGRPCService(server, db, redis)

	log.Printf("gRPC server running on :%s\n", cfg.GRPCPort)
	log.Fatal(server.Serve(lis))
}
