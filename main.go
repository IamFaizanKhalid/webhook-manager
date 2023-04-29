package main

import (
	"github.com/IamFaizanKhalid/webhook-api/file"
	"log"
	"net/http"
)

const port = ":8000"

func main() {
	// Start server
	log.Printf("Starting server on the port %s", port)
	log.Fatal(http.ListenAndServe(port, router()))
}

var hooks *file.File

func init() {
	var err error
	hooks, err = file.Parse("hooks.yml")
	if err != nil {
		panic(err)
	}
}
