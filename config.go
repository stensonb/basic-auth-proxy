package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/creack/goproxy/registry"
)

type config struct {
	User     string
	Password string
	Registry registry.DefaultRegistry
}

func NewConfig(configFilePath string) config {
	b, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var m config
	err = json.Unmarshal(b, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}
