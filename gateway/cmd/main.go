package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/mauriciomd/go-gateway/config"
	"github.com/mauriciomd/go-gateway/internal"
	"github.com/mauriciomd/go-gateway/plugins"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Configuration file not found.")
	}

	filename := os.Args[1]
	timer := time.Duration(10 * time.Second)
	serviceConfig, err := config.New(filename, timer)
	if err != nil {
		panic(err)
	}

	plugins.RegisterMiddleware("http_log", plugins.APILog)
	plugins.RegisterMiddleware("add_header", plugins.AddHeader)
	plugins.RegisterMiddleware("request_size_limiting", plugins.RequestSizeLimiter)
	plugins.RegisterMiddleware("jwt_auth", plugins.AuthorizationHeader)

	server := internal.NewServer(serviceConfig)

	slog.Info("Server listening on port 8080")
	http.Handle("/", server)
	http.ListenAndServe(":8080", nil)
}
