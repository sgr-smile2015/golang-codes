package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr        string
	Port        int
	Protocol    string
	Timeout     time.Duration
	MaxConnects int
	TLS         *tls.Config
}

type Option func(*Server)

func Protocol(p string) Option {
	return func(server *Server) {
		server.Protocol = p
	}
}

func Timeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}

func MaxConnects(maxConnects int) Option {
	return func(server *Server) {
		server.MaxConnects = maxConnects
	}
}

func TLS(tls *tls.Config) Option {
	return func(server *Server) {
		server.TLS = tls
	}
}

func NewServer(addr string, port int, options ...func(server *Server)) (*Server, error) {
	svr := Server{
		Addr:        addr,
		Port:        port,
		Protocol:    "tcp",
		Timeout:     30 * time.Second,
		MaxConnects: 1024,
		TLS:         nil,
	}

	for _, option := range options {
		option(&svr)
	}

	return &svr, nil

}

func main() {
	fmt.Println("function options mod.")
	s1, _ := NewServer("localhost", 8010)
	s2, _ := NewServer("localhost", 8011, Protocol("udp"))
	server, _ := NewServer("0.0.0.0", 443, Timeout(300*time.Second), MaxConnects(2048))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(server)
}
