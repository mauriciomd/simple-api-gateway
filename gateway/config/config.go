package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"
)

type ServiceConfig struct {
	Services     []Service `yaml:"services"`
	lastModified time.Time
}

func New(filename string, timer time.Duration) (*ServiceConfig, error) {
	var config ServiceConfig

	err := config.loadConfiguration(filename)
	if err != nil {
		return nil, err
	}

	config.startObserver(filename, timer)
	return &config, nil
}

func (c *ServiceConfig) startObserver(filename string, timer time.Duration) {
	ticker := time.NewTicker(timer)
	go func() {
		for {
			select {
			case <-ticker.C:
				state, _ := os.Stat(filename)
				if c.lastModified != state.ModTime() {
					err := c.loadConfiguration(filename)
					if err != nil {
						message := fmt.Sprintf("Could not refresh the config - %s.\n", err.Error())
						slog.Error(message)
					} else {
						slog.Info("Configuration updated.")
					}
				}
			}
		}
	}()
}

func (c *ServiceConfig) GetServiceRoute(p string) (*Service, *Route) {
	for _, service := range c.Services {
		for _, route := range service.Routes {
			if route.regExp.Match([]byte(p)) {
				return &service, &route
			}
		}
	}

	return nil, nil
}
