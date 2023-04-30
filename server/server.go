package server

import (
	"context"
	"github.com/IamFaizanKhalid/webhook-api/internal/server/http"
	"github.com/IamFaizanKhalid/webhook-api/server/api"
	"github.com/IamFaizanKhalid/webhook-api/server/logic"
	"github.com/IamFaizanKhalid/webhook-api/server/repo"
)

type Server struct {
	cfg *Config
}

func New(cfg Config) *Server {
	return &Server{&cfg}
}

type Config struct {
	HttpPort  int    `yaml:"httpPort"`
	HooksFile string `yaml:"hooksFile"`
}

func (s *Server) Start(ctx context.Context) error {
	// data store
	ds, err := repo.NewRepo(s.cfg.HooksFile)
	if err != nil {
		return err
	}

	// service
	svc := logic.New(ds)

	// server
	httpSrv := http.New(s.cfg.HttpPort,
		api.NewPing(),
		api.NewHook(svc),
		api.NewStatic("server/static"),
	)

	return httpSrv.Start()
}
