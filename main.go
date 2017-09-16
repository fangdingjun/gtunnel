package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var dialer = &net.Dialer{Timeout: 10 * time.Second}

func main() {
	var configfile string

	flag.StringVar(&configfile, "c", "config.yaml", "config file")
	flag.Parse()

	cfg, err := loadConfig(configfile)
	if err != nil {
		log.Fatal(err)
	}

	initServer(cfg)

	select {}
}

func initServer(cfg *conf) {
	for _, srv := range *cfg {
		go initListener(srv)
	}
}

func initListener(srv server) {
	var l net.Listener
	var err error

	host := net.JoinHostPort(srv.Listen.Host, fmt.Sprintf("%d", srv.Listen.Port))

	if srv.Listen.Cert != "" && srv.Listen.Key != "" {
		cert, err := tls.LoadX509KeyPair(srv.Listen.Cert, srv.Listen.Key)
		if err != nil {
			log.Fatal(err)
		}
		config := &tls.Config{
			Certificates: []tls.Certificate{cert},
		}

		l, err = tls.Listen("tcp", host, config)
	} else {
		l, err = net.Listen("tcp", host)
	}

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			break
		}
		go handleConn(conn, srv.Backend)
	}
}

func handleConn(conn net.Conn, b backend) {
	var c net.Conn
	var err error

	host := net.JoinHostPort(b.Host, fmt.Sprintf("%d", b.Port))
	if b.TLS {
		hostname := b.Host
		if b.Hostname != "" {
			hostname = b.Hostname
		}
		config := &tls.Config{
			ServerName:         hostname,
			InsecureSkipVerify: b.Insecure,
		}
		c, err = tls.DialWithDialer(dialer, "tcp", host, config)
	} else {
		c, err = dialer.Dial("tcp", host)
	}

	if err != nil {
		log.Println(err)
		return
	}

	pipeAndClose(conn, c)
}

func pipeAndClose(c1, c2 net.Conn) {
	defer c1.Close()
	defer c2.Close()

	ch := make(chan struct{}, 2)
	go func() {
		io.Copy(c1, c2)
		ch <- struct{}{}
	}()

	go func() {
		io.Copy(c2, c1)
		ch <- struct{}{}
	}()

	<-ch
}
