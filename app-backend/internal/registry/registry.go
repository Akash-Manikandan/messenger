package registry

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type RouteRegistrar func(app *fiber.App, db *gorm.DB)
type GRPCRegistrar func(server *grpc.Server, db *gorm.DB)

var registrars []RouteRegistrar
var grpcRegistrars []GRPCRegistrar

func Register(r RouteRegistrar) {
	registrars = append(registrars, r)
}

func RegisterGRPC(r GRPCRegistrar) {
	grpcRegistrars = append(grpcRegistrars, r)
}

func Load(app *fiber.App, db *gorm.DB) {
	for _, r := range registrars {
		r(app, db)
	}
}

func LoadGRPC(server *grpc.Server, db *gorm.DB) {
	for _, r := range grpcRegistrars {
		r(server, db)
	}
}
