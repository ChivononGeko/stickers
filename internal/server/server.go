package server

import (
	"net/http"
	"stickers/internal/processor"
	"stickers/internal/server/handlers"
	"stickers/internal/server/routes"
)

type Server struct {
	mux       *http.ServeMux
	processor *processor.Processor
}

func NewServer(p *processor.Processor) *Server {
	s := &Server{
		mux:       http.NewServeMux(),
		processor: p,
	}

	orderHandler := handlers.NewOrderHandler(p)
	routes.RegisterRoutes(s.mux, orderHandler)

	return s
}

func (s *Server) Router() *http.ServeMux {
	return s.mux
}
