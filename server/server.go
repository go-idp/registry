package server

import (
	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/zoox"
)

// Server is the interface for the server
type Server interface {
	Run(cfg *Config) error
	//
	Info() zoox.HandlerFunc
}

type server struct {
	requests safe.Int64
}

// New creates a new server
func New() Server {
	return &server{}
}
