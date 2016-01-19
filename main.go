package main

import (
	"flag"
	"github.com/fangdingjun/iniflags"
	"log"
)

var remote string

func main() {
	var server, client bool
	flag.StringVar(&remote, "remote", "", "remote server")
	flag.IntVar(&port, "port", 8080, "the listen port")
	flag.BoolVar(&server, "server", false, "tls server mode")
	flag.BoolVar(&client, "client", false, "tls client mode")
	flag.StringVar(&cert, "cert", "", "the certificate file")
	flag.StringVar(&key, "key", "", "the private key")
	iniflags.Parse()

	if remote == "" {
		log.Fatal("please use --remote to special the server")
	}

	if server {
		if cert == "" || key == "" {
			log.Fatal("in server mode, you must special the certificate and private key")
		}
		server_main()
		return
	}

	if client {
		local_main()
		return
	}

	log.Fatal("please use --server or --client to special a work mode")
}
