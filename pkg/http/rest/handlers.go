package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jared-paxton/foo-api/pkg/repository"
)

func (app *application) getFoo(w http.ResponseWriter, r *http.Request) {
	fooID := chi.URLParam(r, "id")

	foo, err := app.fooService.Get(fooID)
	if err != nil {
		switch err {
		case repository.ErrFooNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}

		return
	}

	err = sendJSON(w, 200, foo)
	if err != nil {
		app.logger.Println("could not send response to getFoo request")
	}
}
