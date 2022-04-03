// Package repository creates errors that the rest of the application
// can referene, and an interface for data storage systems to implement.
package repository

import (
	"errors"

	"github.com/jared-paxton/foo-api/pkg/models"
)

// ErrFooNotFound is returned if the Foo object is not found
// within the data store.
var ErrFooNotFound error = errors.New("foo does not exist")

// Storage is an interface which, if implemented, allows the
// implementor to be the data storage system for the application.
// It contains all functions from relevant Repository interfaces.
type Storage interface {
	GetFoo(id string) (models.Foo, error)
}
