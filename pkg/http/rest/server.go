package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jared-paxton/foo-api/pkg/foo"
	"github.com/jared-paxton/foo-api/pkg/repository"
)

type application struct {
	fooService foo.Service
	logger     *log.Logger
}

// RunServer starts the REST API server on the given port with
// the specified data storage system.
func RunServer(storage repository.Storage, port int) {
	app := application{
		fooService: foo.NewFooService(storage),
		logger:     log.New(os.Stderr, "", log.Ldate|log.Ltime),
	}

	fmt.Printf("REST server starting at port: %d\n", port)

	address := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(address, app.routes())
	if err != nil {
		app.logger.Fatalln(err)
	}
}
