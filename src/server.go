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
		log.Printf("user %s invalid password %s\n", username, password)
		return false
	}

	return true
}

type ServerConfig struct {
	Addr string
}

type Server struct {
	cfg   *ServerConfig
	cator socks5.Authenticator
}

func NewServer(cfg *ServerConfig, auth socks5.StaticCredentials) *Server {
	cator := socks5.UserPassAuthenticator{Credentials: auth}
	return &Server{
		cfg:   cfg,
		cator: cator,
	}
}

func (s *Server) ListenAndServe() error {
	socks5Config := &socks5.Config{
		AuthMethods: []socks5.Authenticator{s.cator},
	}

	srv, err := socks5.New(socks5Config)
	if err != nil {
		return err
	}

	return srv.ListenAndServe("tcp", s.cfg.Addr)
}
