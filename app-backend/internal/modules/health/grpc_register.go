package health

import (
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func RegisterGRPCService(server *grpc.Server, db *gorm.DB) {
	service := NewService(db)
	grpcServer := NewGRPCServer(service)

	RegisterHealthServiceServer(server, grpcServer)
}
