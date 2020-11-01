package server

import (
	"net"
	"net/http"

	"github.com/Jose-Huerta-WebBeds/GoTraining1/api/handlers"
)

//StartServer creates an initialize a Server object
func StartServer(protocol string, port string) error {
	server := new(http.Server)
	server = new(http.Server)
	server.Handler = handlers.Handler()
	netPort, _ := net.Listen(protocol, port)
	server.Serve(netPort)
	return nil
}

//StartDefaultServer creates a new server with all the default values
func StartDefaultServer() error {
	return StartServer("tcp", ":8090")
}
