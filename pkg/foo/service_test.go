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

func (ms *mockStorage) DeleteFoo(id string) error {
	return nil
}

func TestGet(t *testing.T) {
	mockStorage := initMockStorage()
	mockFooService := NewFooService(mockStorage)

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
			_, err := mockFooService.Get(test.fooID)

			if err != test.want {
				t.Errorf("test failed: %v - errors not the same. got: %v, wanted: %v\n", test.name, err, test.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	mockStorage := initMockStorage()
	mockFooService := NewFooService(mockStorage)

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
			newFoo, err := mockFooService.New(test.fooName)

			if err != test.wantErr {
				t.Errorf("test failed: %v - should not have gotten this error: %s\n", test.name, err)
			}
			if newFoo.Name != test.wantFooName {
				t.Errorf("test failed: %v - new Foo name (%s) was not the same as wanted Foo name (%s)\n", test.name, newFoo.Name, test.wantFooName)
			}

			generatedIDs = append(generatedIDs, newFoo.ID)
		})
	}

	if !areIDsUnique(generatedIDs) {
		t.Error("test failed: generated IDs should be unique, but they are not.")
	}
}

func TestRemove(t *testing.T) {
	mockStorage := initMockStorage()
	mockFooService := NewFooService(mockStorage)

	tests := []struct {
		name  string
		fooID string
		want  error
	}{
		{
			name:  "should delete foo successfully",
			fooID: testFoos[0].ID,
			want:  nil,
		},
		{
			name:  "should not delete foo since it doesn't exist",
			fooID: "blahblahblah",
			want:  repository.ErrFooNotFound,
		},
		{
			name:  "should delete the last foo successfully",
			fooID: testFooIDs[1],
			want:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mockFooService.Remove(test.fooID)

			if err != test.want {
				t.Errorf("test failed: %v - should not have gotten this error: %s - for id: %s\n", test.name, err, test.fooID)
			}
		})
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
