// Package ezcors handles CORS configuration file rs.cors Options.
package ezcors

import (
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// CORS defines the supported properties.
type CORS struct {
	AllowedOrigins     []string `yaml:"allowedOrigins"`
	AllowCredentials   bool     `yaml:"allowCredentials"`
	AllowedMethods     []string `yaml:"allowedMethods"`
	AllowedHeaders     []string `yaml:"allowedHeaders"`
	ExposedHeaders     []string `yaml:"exposedHeaders"`
	MaxAge             int      `yaml:"maxAge"`
	OptionsPassthrough bool     `yaml:"optionsPassthrough"`
	Debug              bool     `yaml:"debug"`
}

// Config defines CORS configuration for every environment.
type Config map[string]CORS

// Option defines CORS options.
type Option struct {
	// Path is the location of `cors.yml` file.
	Path string
}

// NewConfig decodes config file and returns CORS Config.
// The function will look for cors.yml file from the current directory. If
// nothing can found it will try to look into the config directory for possible
// cors.yml file.
func NewConfig(opts ...Option) (Config, error) {
	path := "cors.yml"

	for _, opt := range opts {
		if opt.Path != "" {
			path = opt.Path
		}
	}

	file, err := os.Open(path)
	if err != nil {
		path = "config/cors.yml"
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}
	defer file.Close()

	config := Config{}
	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

// AllowedOrigin sets the allowed CORS origin for the request.
// It checks for the `Origin` key from request headers otherwise it will use
// the first item from `cors` slice of allowed origins.
func AllowedOrigin(h http.Header, cors []string) string {
	found := false
	origin := h.Get("Origin")

	for _, v := range cors {
		if v == origin {
			found = true
		}
	}

	if !found {
		return cors[0]
	}

	return origin
}
