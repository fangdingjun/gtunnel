package main

import (
	//"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type server struct {
	Listen  listen
	Backend backend
}

type conf []server

type listen struct {
	Host string
	Port int
	Cert string
	Key  string
}

type backend struct {
	Host     string
	Port     int
	Hostname string
	TLS      bool
	Insecure bool
}

func loadConfig(fn string) (*conf, error) {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}

	c := new(conf)
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
