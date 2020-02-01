package server

import "sync"

type Server struct {

}

var (
	SgtServer *Server
	once sync.Once
)
