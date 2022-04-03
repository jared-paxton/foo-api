// Package foo implements a the business logic in the
// application for handling foo objects.
package foo

import (
	"github.com/jared-paxton/foo-api/pkg/models"
)

// Repository defines an interface that must be implemented
// by any data storage system that interacts with Foo objects.
type Repository interface {
	GetFoo(id string) (models.Foo, error)
}

// Service defines an interface to handle the business logic
// for Foo objects.
type Service interface {
	Get(id string) (models.Foo, error)
}

type service struct {
	repo Repository
}

// NewFooService initializes a Foo service with the given
// repository (data storage system).
func NewFooService(r Repository) Service {
	return &service{
		repo: r,
	}
}

// Get fetches the Foo object with the given unique id.
func (fs *service) Get(id string) (models.Foo, error) {
	foo, err := fs.repo.GetFoo(id)
	if err != nil {
		return foo, err
	}

	return foo, nil
}
