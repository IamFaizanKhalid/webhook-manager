package server

import (
	"context"
	"github.com/IamFaizanKhalid/webhook-api/internal/auth"
	"github.com/IamFaizanKhalid/webhook-api/internal/server/http"
	"github.com/IamFaizanKhalid/webhook-api/server/api"
	"github.com/IamFaizanKhalid/webhook-api/server/logic"
	"github.com/IamFaizanKhalid/webhook-api/server/repo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg *Config
}

func New(cfg Config) *Server {
	return &Server{&cfg}
}

type Config struct {
	HttpPort    int    `yaml:"httpPort"`
	ApiKey      string `yaml:"apiKey"`
	HooksFile   string `yaml:"hooksFile"`
	WebhookPort int    `yaml:"webhookPort"`
}

func (s *Server) Start(ctx context.Context) error {
	// data store
	ds, err := repo.NewRepo(s.cfg.HooksFile)
	if err != nil {
		return err
	}

	// auth
	authHandler := auth.New(&authenticator{apiKey: s.cfg.ApiKey})

	// service
	svc := logic.New(ds)

	// server
	httpSrv := http.New(s.cfg.HttpPort, authHandler,
		api.NewPing(),
		api.NewHook(svc),
		api.NewStatic("server/static"),
	)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals
		log.Println("received signal from the os: ", sig)
		ctx.Done()
		os.Exit(0)
	}()

	return httpSrv.Start()
}
