package internal

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/mauriciomd/go-gateway/config"
	"github.com/mauriciomd/go-gateway/plugins"
)

type Server struct {
	httpClient *http.Client
	config     *config.ServiceConfig
}

func NewServer(c *config.ServiceConfig) *Server {
	return &Server{
		config:     c,
		httpClient: &http.Client{},
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	service, route := s.config.GetServiceRoute(path)
	if service == nil || route == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ok := route.IsMethodAllowed(r.Method)
	if !ok {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	f := func(w http.ResponseWriter, r *http.Request) {
		res, err := s.doRequest(r, service)
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		s.writeResponse(w, res)
	}

	for _, plugin := range service.Plugins {
		m, err := plugins.GetMiddleware(plugin.Name)
		if err != nil {
			msg := fmt.Sprintf("Plugin %s not found", plugin.Name)
			slog.Error(msg)
			continue
		}

		f = m(plugin, f)
	}

	f(w, r)
}

func (s *Server) doRequest(r *http.Request, service *config.Service) (*http.Response, error) {
	url := service.Url + r.URL.Path
	req, _ := http.NewRequest(r.Method, url, r.Body)

	req.Header = r.Header

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Server) writeResponse(w http.ResponseWriter, res *http.Response) error {
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	for k := range res.Header {
		w.Header().Set(k, res.Header.Get(k))
	}

	w.WriteHeader(res.StatusCode)
	w.Write(data)

	return nil
}
