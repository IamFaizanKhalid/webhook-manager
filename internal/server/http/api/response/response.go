package response

import (
	"encoding/json"
	"errors"
	errors2 "github.com/IamFaizanKhalid/webhook-api/internal/errors"
	"net/http"
)

func Encode(w http.ResponseWriter, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(payload)
}

func EncodeErr(w http.ResponseWriter, err error) error {
	e, ok := err.(*errors2.Response)
	if !ok {
		return errors.New("unable to parse error")
	}

	w.WriteHeader(e.HTTPStatusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(e)
}
