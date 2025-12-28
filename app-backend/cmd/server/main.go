package main

import (
	_ "github.com/Akash-Manikandan/app-backend/internal/modules"
	"github.com/Akash-Manikandan/app-backend/internal/server"
)

func main() {
	go server.StartGRPC()
	server.StartHTTP()
}
