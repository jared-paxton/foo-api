// Package memory implements storing the application data in memory.
package memory

import (
	"github.com/jared-paxton/foo-api/pkg/models"
	"github.com/jared-paxton/foo-api/pkg/repository"
)

// storage contains the data structures capable of holding
// application data in memory.
type storage struct {
	foos map[string]models.Foo
}

// NewStorage allocates an initializes the in-memory store for
// the application data.
func NewStorage() *storage {
	return &storage{
		foos: make(map[string]models.Foo),
	}
}

// GetFoo returns the Foo object from memory based on the
// given id. It returns a specfici error if it is not found.
func (mem *storage) GetFoo(id string) (models.Foo, error) {
	var foo models.Foo

	if foo, ok := mem.foos[id]; ok {
		return foo, nil
	}

	return foo, repository.ErrFooNotFound
}


