package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
)

//var backend = "127.0.0.1:8080"
var cert = "server.crt"
var key = "server.key"
var port = 9000

func server_main() {
	certificate, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}
	l, err := tls.Listen("tcp", fmt.Sprintf(":%d", port), &tls.Config{
		Certificates: []tls.Certificate{certificate},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle_server(c)
	}
}

func handle_server(c net.Conn) {
	log.Printf("accept connection from %s", c.RemoteAddr())
	s, err := net.Dial("tcp", remote)
	if err != nil {
		log.Println(err)
		c.Close()
		return
	}

	ch := make(chan int)
	go func() {
		count, _ := io.Copy(s, c)
		s.Close()
		log.Printf("write %d bytes to %s", count, s.RemoteAddr())
		ch <- 1
	}()

	co, _ := io.Copy(c, s)
	c.Close()
	log.Printf("write %d bytes to %s", co, c.RemoteAddr())
	<-ch
}
