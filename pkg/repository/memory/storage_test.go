package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jared-paxton/foo-api/pkg/models"
)

func TestGetFoo(t *testing.T) {
	storage := NewStorage()

	foos := addTestData(storage)

	for _, foo := range foos {
		fetchedFoo, err := storage.GetFoo(foo.ID)
		if err != nil {
			t.Error("expected no error, but got:", err)
		}
		if !fetchedFoo.Equals(&foo) {
			t.Errorf("foo name (%s) and/or ID (%s) was not equal to fetchedFoo name (%s) and/or ID (%s)\n",
				foo.Name, foo.ID, fetchedFoo.Name, fetchedFoo.ID)
		}
	}
}

func addTestData(memStorage *storage) []models.Foo {
	foos := []models.Foo{
		{
			ID:   uuid.NewString(),
			Name: "name0",
		},
		{
			ID:   uuid.NewString(),
			Name: "name1",
		},
		{
			ID:   uuid.NewString(),
			Name: "name2",
		},
	}

	for _, foo := range foos {
		memStorage.foos[foo.ID] = foo
	}

	return foos
}
