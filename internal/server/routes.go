package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}
