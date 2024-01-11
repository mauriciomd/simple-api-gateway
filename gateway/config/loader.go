package config

import (
	"io"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

func (c *ServiceConfig) loadConfiguration(filename string) error {
	filePointer, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer filePointer.Close()

	err = c.refreshConfig(filePointer)
	if err != nil {
		return err
	}

	stat, err := filePointer.Stat()
	if err != nil {
		return err
	}

	c.lastModified = stat.ModTime()
	return nil
}

func (c *ServiceConfig) refreshConfig(f *os.File) error {
	fileContent, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileContent, c)
	if err != nil {
		return err
	}

	c.createRouteRegex()
	return nil
}
func (c *ServiceConfig) createRouteRegex() {
	r, _ := regexp.Compile("{[a-zA-Z0-9]+}")

	for i := range c.Services {
		for j := range c.Services[i].Routes {
			replaced := r.ReplaceAll([]byte(c.Services[i].Routes[j].Paths[0]), []byte("([a-zA-Z0-9]+)"))
			regExp, err := regexp.Compile("^" + string(replaced) + "$")
			if err != nil {
				continue
			}
			c.Services[i].Routes[j].regExp = regExp
		}
	}
}
