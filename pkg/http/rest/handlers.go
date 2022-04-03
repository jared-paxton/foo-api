package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) fetchFoo(w http.ResponseWriter, r *http.Request) {
	fooID := chi.URLParam(r, "id")

	foo, err := app.fooService.Get(fooID)
	if err != nil {
		sendError(w, err)
		return
	}

	err = sendJSON(w, 200, foo)
	if err != nil {
		app.logger.Println("could not send response to getFoo request")
	}
}

func (app *application) addFoo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var fooReq FooRequest
	err := decoder.Decode(&fooReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newFoo, err := app.fooService.New(fooReq.Name)
	if err != nil {
		sendError(w, err)
		return
	}

	err = sendJSON(w, 200, newFoo)
	if err != nil {
		app.logger.Println("could not send response to addFoo request")
	}
}
