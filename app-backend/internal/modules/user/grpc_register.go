package user

import (
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func RegisterGRPCService(server *grpc.Server, db *gorm.DB, redis *redis.Client) {
	service := NewService(db, redis)
	grpcServer := NewGRPCServer(service)

	RegisterUserServiceServer(server, grpcServer)
}
