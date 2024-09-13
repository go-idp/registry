package server

import (
	"github.com/go-zoox/core-utils/safe"
	"github.com/go-zoox/zoox"
)

type Server interface {
	Run(cfg *Config) error
	//
	Info() zoox.HandlerFunc
}

type server struct {
	requests safe.Int64
}

func New() Server {
	return &server{}
}
