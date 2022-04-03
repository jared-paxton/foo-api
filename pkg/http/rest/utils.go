package rest

import (
	"encoding/json"
	"net/http"
)

func sendJSON(w http.ResponseWriter, status int, data interface{}) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)

	return nil
}
