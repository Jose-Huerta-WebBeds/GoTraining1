package serverfactory

import (
	"github.com/Jose-Huerta-WebBeds/GoTraining1/api/handlers"
	"github.com/Jose-Huerta-WebBeds/GoTraining1/infrastructure/server"
)

const (
	protocol = "tcp"
	port     = ":8090"
)

//StartMainServer creates and Starts the httpServer
func StartMainServer() {
	s, _ := server.NewServer(protocol, port, handlers.Handler())
	s.Start()
}
