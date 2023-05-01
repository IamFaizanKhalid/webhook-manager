package http

import (
	"fmt"
	"github.com/IamFaizanKhalid/webhook-api/internal/errors"
	"github.com/IamFaizanKhalid/webhook-api/internal/server/http/api"
	"github.com/IamFaizanKhalid/webhook-api/internal/server/http/api/response"
	"log"
	"net/http"
)

type Server struct {
	port   int
	apiKey string
	apis   []api.API
}

func New(port int, apiKey string, apis ...api.API) *Server {
	return &Server{port: port, apiKey: apiKey, apis: apis}
}

// Start starts the api server
func (srv *Server) Start() error {
	r := srv.buildRouter(srv.authMiddleware)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	httpPort := fmt.Sprintf(":%d", srv.port)
	log.Printf("Starting server on %v\n", httpPort)

	return http.ListenAndServe(httpPort, r)
}

func (srv *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-KEY") != srv.apiKey {
			response.EncodeErr(w, errors.Unauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
