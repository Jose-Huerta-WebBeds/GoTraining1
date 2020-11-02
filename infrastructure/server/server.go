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
	s.httpServer.Handler = s.preHandler(handlers.Handler())
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

//this pre Handler can inject any logic before or after the real handler
func (s *Server) preHandler(trueHandler http.Handler) *http.ServeMux {
	mux := new(http.ServeMux)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		allow := s.preHandlerLimiter()
		defer s.postHandlerLimiter()

		if allow != true {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		trueHandler.ServeHTTP(w, r)

	})

	return mux
}

func (s *Server) preHandlerLimiter() bool {
	return false
}

func (s *Server) postHandlerLimiter() {

}
