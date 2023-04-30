package response

import (
	"encoding/json"
	"errors"
	"github.com/IamFaizanKhalid/webhook-api/server/logic/output"
	"net/http"
)

func Encode(w http.ResponseWriter, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(payload)
}

func EncodeErr(w http.ResponseWriter, err error) error {
	e, ok := err.(output.ErrResponse)
	if !ok {
		return errors.New("parseRequest: target is expected to be of type *request.Request")
	}

	w.WriteHeader(e.HTTPStatusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(e)
}
