package rest

import (
	"encoding/json"
	"net/http"

	"github.com/jared-paxton/foo-api/pkg/repository"
)

func sendJSON(w http.ResponseWriter, status int, data interface{}) error {
	// Format according to provided sample output
	json, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
	// Added to more closely match provided sample output
	w.Write([]byte("\n"))

	return nil
}

func sendError(w http.ResponseWriter, err error) {
	switch err {
	case repository.ErrFooNotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
