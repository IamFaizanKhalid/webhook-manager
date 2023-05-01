package http

import (
	"fmt"
	"github.com/IamFaizanKhalid/webhook-manager/internal/auth"
	"github.com/IamFaizanKhalid/webhook-manager/internal/server/http/api"
	"log"
	"net/http"
)

type Server struct {
	port int
	auth *auth.Auth
	apis []api.API
}

func New(port int, auth *auth.Auth, apis ...api.API) *Server {
	return &Server{port: port, auth: auth, apis: apis}
}

// Start starts the api server
func (srv *Server) Start() error {
	r := srv.buildRouter(srv.auth)

	httpPort := fmt.Sprintf(":%d", srv.port)
	log.Printf("Starting server on %v\n", httpPort)

	return http.ListenAndServe(httpPort, r)
}
