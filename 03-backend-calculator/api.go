package main

import (
	"calc-api/service/calc"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	calcHandler := calc.NewHandler()
	calcHandler.RegisterRoutes(router)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started in PORT %s", s.addr)
	return server.ListenAndServe()
}
