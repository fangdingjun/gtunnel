gtunnel
======

A stunnel-like util write by golang

work mode
========

- listen plain, forward through plain
- listen plain, forward through TLS
- listen TLS, forward through plain
- listen TLS, forward through TLS


usage
====

    go get github.com/fangdingjun/gtunnel
    cp $GOPATH/src/github.com/fangdingjun/config_example.yaml config.yaml
    vim config.yaml
    $GOPATH/bin/gtunnel -c config.yaml

