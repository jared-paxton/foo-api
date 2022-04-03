// Package models defines the entities used in the application.
package models

// Foo represents the Foo object and its properties.
type Foo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Equals returns true if all the members of the Foo objects
// are equal, false otherwise.
func (foo *Foo) Equals(f *Foo) bool {
	if foo.ID != f.ID {
		return false
	}
	if foo.Name != f.Name {
		return false
	}

	return true
}
