package health

import "github.com/Akash-Manikandan/app-backend/internal/registry"

func init() {
	registry.Register(RegisterRoutes)
	registry.RegisterGRPC(RegisterGRPCService)
}
