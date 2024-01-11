package plugins

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/mauriciomd/go-gateway/config"
)

func APILog(p config.Plugin, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("Request received on %s with method %s.", r.URL, r.Method)
		slog.Info(msg)

		f(w, r)
	}
}

func AddHeader(p config.Plugin, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for header, value := range p.Input {
			r.Header.Add(header, value.(string))
		}

		f(w, r)
	}
}

func RequestSizeLimiter(p config.Plugin, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("could not read request body", slog.String("error", err.Error()))
		}

		if p.Input["allowed_payload_size"] == nil || len(body) >= p.Input["allowed_payload_size"].(int) {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			return
		}

		f(w, r)
	}
}
