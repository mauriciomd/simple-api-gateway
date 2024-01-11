package plugins

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mauriciomd/go-gateway/config"
)

func AuthorizationHeader(p config.Plugin, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := getToken(p, r)
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, err = jwt.Parse(token, func(_ *jwt.Token) (interface{}, error) {
			secret := p.Input["secret"].(string)
			return []byte(secret), nil
		})

		if err != nil {
			slog.Error("Could not parse the token.", slog.String("error", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		f(w, r)
	}
}

func getToken(p config.Plugin, r *http.Request) (string, error) {
	key := p.Input["key_name"].(string)

	if p.Input["key_in_header"] != nil && p.Input["key_in_header"].(bool) {
		header := r.Header.Get(key)
		parts := strings.Split(header, " ")
		if len(parts) < 2 {
			return "", errors.New("Invalid header format. Expected 'Bearer <jwt-token>'.")
		}

		return parts[1], nil
	}

	if p.Input["key_in_query"] != nil && p.Input["key_in_query"].(bool) {
		return r.URL.Query().Get(key), nil
	}

	return "", errors.New("")
}
