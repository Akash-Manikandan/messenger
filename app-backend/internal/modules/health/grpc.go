package health

import "context"

type GRPCServer struct {
	UnimplementedHealthServiceServer
	service Service
}

func NewGRPCServer(s Service) *GRPCServer {
	return &GRPCServer{service: s}
}

func (g *GRPCServer) Check(ctx context.Context, _ *HealthRequest) (*HealthResponse, error) {
	return &HealthResponse{
		Status:           g.service.Status(),
		PostgresDbStatus: g.service.PostgresStatus(),
	}, nil
}
