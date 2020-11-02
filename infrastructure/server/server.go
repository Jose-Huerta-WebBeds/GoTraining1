package server

import (
	"net"
	"net/http"

	"github.com/Jose-Huerta-WebBeds/GoTraining1/api/handlers"
)

//Server an HTTP server with protections
type Server struct {
	httpServer *http.Server
	port       string
	protocol   string
}

//NewServer Creates a new server and configures it
func NewServer(protocol string, port string, handler http.Handler) (*Server, error) {
	s := new(Server)
	s.httpServer = new(http.Server)
	s.httpServer.Handler = handlers.Handler()
	s.port = port
	s.protocol = protocol
	return s, nil
}

//Start starts the server
func (s *Server) Start() error {
	netPort, err := net.Listen(s.protocol, s.port)
	if err != nil {
		return err
	}
	s.httpServer.Serve(netPort)
	return nil
}
