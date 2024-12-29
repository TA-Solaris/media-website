package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"github.com/TA-Solaris/common"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.RegisterRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start the HTTP server")
	}
}