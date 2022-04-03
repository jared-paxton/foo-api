// Package foo implements a the business logic in the
// application for handling foo objects.
package foo

import (
	"github.com/google/uuid"
	"github.com/jared-paxton/foo-api/pkg/models"
	"github.com/jared-paxton/foo-api/pkg/repository"
)

// Repository defines an interface that must be implemented
// by any data storage system that interacts with Foo objects.
type Repository interface {
	GetFoo(id string) (models.Foo, error)
	CreateFoo(models.Foo) error
	DeleteFoo(id string) error
}

// Service defines an interface to handle the business logic
// for Foo objects.
type Service interface {
	Get(id string) (models.Foo, error)
	New(name string) (models.Foo, error)
	Remove(id string) error
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

// Get fetches the Foo object with the given unique id
// from storage, and returns an error if it's not found.
func (fs *service) Get(id string) (models.Foo, error) {
	foo, err := fs.repo.GetFoo(id)
	if err != nil {
		return foo, err
	}

	return foo, nil
}

// New create a new Foo object with the given name and generates
// a UUID for the id. It then stores it, or returns an error.
func (fs *service) New(name string) (models.Foo, error) {
	newFoo := models.Foo{
		ID:   uuid.NewString(),
		Name: name,
	}

	// Check if there is another Foo in storage with the same UUID
	// (highly unlikely). An error is expected here.
	_, err := fs.repo.GetFoo(newFoo.ID)
	if err != repository.ErrFooNotFound {
		return newFoo, err
	}

	fs.repo.CreateFoo(newFoo)
	return newFoo, nil
}

// Remove checks storage for the Foo object given the id. It deletes
// it if found, or returns an error.
func (fs *service) Remove(id string) error {
	_, err := fs.repo.GetFoo(id)
	if err != nil {
		return err
	}

	fs.repo.DeleteFoo(id)
	return nil
}
