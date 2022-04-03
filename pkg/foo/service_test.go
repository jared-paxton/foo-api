package foo

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jared-paxton/foo-api/pkg/models"
	"github.com/jared-paxton/foo-api/pkg/repository"
)

type mockStorage struct {
	foos []models.Foo
}

var testFooIDs = []string{
	uuid.NewString(),
	uuid.NewString(),
}

var testFoos []models.Foo = []models.Foo{
	{
		ID:   testFooIDs[0],
		Name: "foo1",
	},
	{
		ID:   testFooIDs[1],
		Name: "foo2",
	},
}

func initMockStorage() *mockStorage {
	storage := &mockStorage{
		foos: testFoos,
	}

	return storage
}

func (ms *mockStorage) GetFoo(id string) (models.Foo, error) {
	var dummyFoo models.Foo

	for _, foo := range ms.foos {
		if foo.ID == id {
			return dummyFoo, nil
		}
	}

	return dummyFoo, repository.ErrFooNotFound
}

func (ms *mockStorage) CreateFoo(models.Foo) error {
	return nil
}

func TestGet(t *testing.T) {
	mockStorage := initMockStorage()
	mockService := NewFooService(mockStorage)

	tests := []struct {
		name  string
		fooID string
		want  error
	}{
		{
			name:  "should get foo successfully",
			fooID: testFoos[0].ID,
			want:  nil,
		},
		{
			name:  "should get foo successfully",
			fooID: testFoos[0].ID,
			want:  nil,
		},
		{
			name:  "should not get Foo and should return an error",
			fooID: "blahblahblah",
			want:  repository.ErrFooNotFound,
		},
		{
			name:  "should not get Foo and should return an error",
			fooID: "",
			want:  repository.ErrFooNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := mockService.Get(test.fooID)

			if err != test.want {
				t.Errorf("test: %v failed - errors not the same. got: %v, wanted: %v\n", test.name, err, test.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	mockStorage := initMockStorage()
	mockService := NewFooService(mockStorage)

	fooName := "testName"

	tests := []struct {
		name        string
		fooName     string
		wantFooName string
		wantErr     error
	}{
		{
			name:        "should add foo successfully",
			fooName:     fooName,
			wantFooName: fooName,
			wantErr:     nil,
		},
		{
			name:        "should add foo with same name successfully",
			fooName:     fooName,
			wantFooName: fooName,
			wantErr:     nil,
		},
		{
			name:        "should add foo with different name successfully",
			fooName:     "diffName",
			wantFooName: "diffName",
			wantErr:     nil,
		},
	}

	var generatedIDs []string

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newFoo, err := mockService.New(test.fooName)

			if err != test.wantErr {
				t.Errorf("test: %v failed - should not have gotten an error (%s)\n", test.name, err)
			}
			if newFoo.Name != test.wantFooName {
				t.Errorf("test: %v failed - should not have gotten an error (%s)\n", test.name, err)
			}

			generatedIDs = append(generatedIDs, newFoo.ID)
		})
	}

	if !areIDsUnique(generatedIDs) {
		t.Error("test: generated IDs should be unique, but they are not.")
	}
}

func areIDsUnique(ids []string) bool {
	for i, idToCheck := range ids {
		for j, otherID := range ids {
			if i == j {
				continue
			}
			if idToCheck == otherID {
				return false
			}
		}
	}
	return true
}
