package plugins

import (
	"errors"
	"net/http"

	"github.com/mauriciomd/go-gateway/config"
)

type Middleware func(p config.Plugin, f http.HandlerFunc) http.HandlerFunc

var plugins map[string]Middleware

func RegisterMiddleware(name string, m Middleware) {
	if plugins == nil {
		plugins = make(map[string]Middleware)
	}

	plugins[name] = m
}

func GetMiddleware(name string) (Middleware, error) {
	m, ok := plugins[name]
	if !ok {
		return nil, errors.New("Plugin not found.")
	}

	return m, nil
}
