package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
)

//var server string = "www.ratafee.nl:443"

func local_main() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listen on %s...", l.Addr())
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle_local(c)
	}
}

func handle_local(c net.Conn) {
	//defer c.Close()
	log.Printf("accept connection from %s", c.RemoteAddr())
	log.Printf("connect to %s...", remote)
	s, err := tls.Dial("tcp", remote, nil)
	if err != nil {
		log.Printf("connect to %s failed: %s", remote, err.Error())
		c.Close()
		return
	}
	ch := make(chan int)
	//defer s.Close()
	go func() {
		count, _ := io.Copy(s, c)
		log.Printf("write %d bytes to %s", count, s.RemoteAddr())
		s.Close()
		ch <- 1
	}()
	co, _ := io.Copy(c, s)
	log.Printf("write %d bytes to %s", co, c.RemoteAddr())
	c.Close()
	<-ch
}
