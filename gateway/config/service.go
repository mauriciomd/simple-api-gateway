package config

import "regexp"

type Service struct {
	Name    string   `yaml:"name"`
	Url     string   `yaml:"url"`
	Plugins []Plugin `yaml:"plugins"`
	Routes  []Route  `yaml:"routes"`
}

type Plugin struct {
	Name  string         `yaml:"name"`
	Input map[string]any `yaml:"input,omitempty"`
}

type Route struct {
	Name    string   `yaml:"name"`
	Paths   []string `yaml:"paths"`
	Methods []string `yaml:"methods"`
	regExp  *regexp.Regexp
}

func (r *Route) IsMethodAllowed(method string) bool {
	for _, m := range r.Methods {
		if m == method {
			return true
		}
	}

	return false
}
