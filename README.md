gtunnel
======

A stunnel like util write by golang

There are two work mode.

###Client mode

In this mode, listen for incoming plain connections and forward to server to use SSL/TLS connections

###Server mode

In this mode, listen for incoming SSL/TLS connections and forward to backend use plain connections

###Build

    go get github.com/fangdingjun/gtunnel
    cd $GOPATH/src/github.com/fangdingjun/gtunnel
    go build

###Usage

server mode

    ./gtunnel --server --cert server.crt --key server.key --port 8001 --remote 127.0.0.1:80

listen for SSL/TLS connections on port 8001 and forward to 127.0.0.1:80

client mode

    ./gtunnel --client --port 8002 --remote www.example.com:8081

listen for plain connections on port 8002 and forward to www.example.com:8081 to use SSL/TLS connections

use `./gtunnel -h` see more options

