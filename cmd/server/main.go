package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/IamFaizanKhalid/webhook-api/internal/utils"
	"github.com/IamFaizanKhalid/webhook-api/server"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// config
	_ = utils.LoadConfig(".env")

	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = server.New(cfg).Start(ctx)
	if err != nil {
		log.Fatalf("start server err: %v", err)
	}
}

func readConfig() (server.Config, error) {
	var c server.Config

	// parsing from arguments
	flag.IntVar(&c.HttpPort, "port", 6543, "Port for http server")
	flag.StringVar(&c.HooksFile, "config", "", "Hooks config file")
	flag.Parse()

	remainingArgs := flag.Args()
	if len(remainingArgs) > 0 {
		fmt.Println("Unknown arguments:", strings.Join(remainingArgs, " "))
		os.Exit(-1)
	}

	// parsing from environment variables
	if x, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 64); err != nil {
		c.HttpPort = int(x)
	}
	if x := os.Getenv("HOOKS_FILE"); x != "" {
		c.HooksFile = x
	}

	// setting default values
	if c.HttpPort == 0 {
		c.HttpPort = 8000
	}

	if c.HooksFile == "" {
		c.HooksFile = "hooks.yml"
	}

	return c, nil
}