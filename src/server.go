package main

import (
	"log"
	"sync"

	"github.com/armon/go-socks5"
)

type Authentication struct {
	mu    sync.Mutex
	users map[string]string
}

func NewAuthentication(users map[string]string) *Authentication {
	return &Authentication{
		users: users,
	}
}

func (a *Authentication) Valid(username, password string) bool {
	pwd, ok := a.users[username]
	if !ok {
		log.Printf("user %s not exist\n", username)
		return false
	}

	if pwd != password {
		log.Printf("user %s invalid password\n", username)
		return false
	}

	return true
}

type ServerConfig struct {
	Addr string
}

type Server struct {
	cfg            *ServerConfig
	Authentication *Authentication
}

func NewServer(cfg *ServerConfig, auth *Authentication) *Server {
	return &Server{
		cfg:            cfg,
		Authentication: auth,
	}
}

func (s *Server) ListenAndServe() error {
	socks5Config := &socks5.Config{
		Credentials: s.Authentication,
	}

	srv, err := socks5.New(socks5Config)
	if err != nil {
		return err
	}

	return srv.ListenAndServe("tcp", s.cfg.Addr)
}
