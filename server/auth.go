package server

import (
	"github.com/IamFaizanKhalid/webhook-manager/internal/errors"
	"net/http"
)

type authenticator struct {
	apiKey string
}

func (a *authenticator) Authenticate(r *http.Request) error {
	if r.Header.Get("X-API-KEY") != a.apiKey {
		return errors.Unauthorized
	}

	return nil
}
