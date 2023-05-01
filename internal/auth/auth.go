package auth

import (
	"github.com/IamFaizanKhalid/webhook-api/internal/server/http/api/response"
	"net/http"
)

type Authenticator interface {
	Authenticate(*http.Request) error
}

type Auth struct {
	client Authenticator
}

func New(client Authenticator) *Auth {
	return &Auth{client}
}

func (a *Auth) HttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := a.client.Authenticate(r); err != nil {
			response.EncodeErr(w, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
