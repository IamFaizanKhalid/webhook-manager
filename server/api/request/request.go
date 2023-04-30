package request

import (
	"encoding/json"
	"fmt"
	"github.com/IamFaizanKhalid/webhook-api/server/logic/output"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Payload interface {
	Validate() error
}

func Decode(r *http.Request, p Payload) error {
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		return err
	}

	return p.Validate()
}

func GetParam(r *http.Request, param string) (string, error) {
	value := chi.URLParam(r, param)
	if value == "" {
		return "", output.ErrInvalidRequest(fmt.Errorf("`%s` is required in URL", param))
	}

	return value, nil
}

func GetRequiredQuery(r *http.Request, key string) (string, error) {
	value := GetQuery(r, key)
	if value == "" {
		return "", output.ErrInvalidRequest(fmt.Errorf("`%s` is required in query", key))
	}

	return value, nil
}

func GetQuery(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
