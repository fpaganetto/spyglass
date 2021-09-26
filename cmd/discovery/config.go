package main

import (
	"flag"
)

type config struct {
	port             int
	localEnvironment bool
}

func NewConfig() config {
	cnf := new(config)

	flag.BoolVar(&cnf.localEnvironment, "local", false, "(optional) run from local environment")
	flag.IntVar(&cnf.port, "port", 8090, "choose port runtime")

	flag.Parse()
	return *cnf
}
