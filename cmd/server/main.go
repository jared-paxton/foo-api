package main

import (
	"github.com/jared-paxton/foo-api/pkg/http/rest"
	"github.com/jared-paxton/foo-api/pkg/repository/memory"
)

func main() {
	port := 8080
	storage := memory.NewStorage()

	rest.RunServer(storage, port)
}
