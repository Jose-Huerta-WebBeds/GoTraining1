package server

import (
	"net"
	"net/http"
	"sync/atomic"
)

//Server an HTTP server with protections
type Server struct {
	httpServer         *http.Server
	port               string
	protocol           string
	concurrentSessions int64
	limit              int
}

//NewServer Creates a new server and configures it
func NewServer(protocol string, port string, handler http.Handler, limit int) (*Server, error) {
	s := new(Server)
	s.limit = limit
	s.httpServer = new(http.Server)
	s.httpServer.Handler = s.preHandler(handler)
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
		allow, status := s.preHandlerLimiter() //logic injected before
		defer s.postHandlerLimiter()           //logic injected after

		if allow != true {
			http.Error(w, http.StatusText(status), status)
			return
		}

		trueHandler.ServeHTTP(w, r)

	})

	return mux
}

//Here we can implement any logic for limiting, like concurrent threads,
//or limit req/sec, per IP (well it would need some arameters...)
func (s *Server) preHandlerLimiter() (bool, int) {
	a := int(atomic.AddInt64(&s.concurrentSessions, 1))
	if a > s.limit {
		return false, http.StatusTooManyRequests
	}
	return true, 0
}

func (s *Server) postHandlerLimiter() {
	atomic.AddInt64(&s.concurrentSessions, -1)
}
