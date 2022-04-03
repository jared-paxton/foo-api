package rest

// FooRequest represents the necessary data in a request to create
// a new Foo object.
type FooRequest struct {
	Name string `json:"name"`
}
