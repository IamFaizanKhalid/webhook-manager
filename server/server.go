package server

import (
	"context"
	"github.com/IamFaizanKhalid/webhook-api/internal/server/http"
	"github.com/IamFaizanKhalid/webhook-api/internal/services/webhook"
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
	HooksFile   string `yaml:"hooksFile"`
	WebhookPort int    `yaml:"webhookPort"`
}

func (s *Server) Start(ctx context.Context) error {
	// webhook server
	wh, err := webhook.NewServer(s.cfg.WebhookPort, s.cfg.HooksFile, os.Stdout)
	if err != nil {
		return err
	}
	err = wh.Start()
	if err != nil {
		return err
	}
	defer wh.Stop()

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

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals
		log.Println("received signal: ", sig)
		err := wh.Stop()
		if err != nil {
			log.Println(err)
		}
		ctx.Done()
	}()

	return httpSrv.Start()
}
