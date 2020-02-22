package main

import (
	"flag"
	"log"

	"github.com/armon/go-socks5"
)

func main() {
	flgConf := flag.String("c", "", "config file path")
	flag.Parse()

	log.SetFlags(log.Lshortfile)

	cfg, err := ParseConfig(*flgConf)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%v\n", cfg)

	auth := socks5.StaticCredentials(cfg.Users)

	s := NewServer(&ServerConfig{
		Addr: cfg.Addr,
	}, auth)

	log.Println(s.ListenAndServe())
}
