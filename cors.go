package ezcors

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// CORS data
type CORS struct {
	AllowedOrigins []string `yaml:"allowedOrigins"`
}

// Config defines CORS configuration for every environment.
type Config map[string]CORS

// NewConfig decodes config file and returns Cors pointer.
func NewConfig() (Config, error) {
	file, err := os.Open("cors.yml")
	if err != nil {
		file, err := os.Open("config/cors.yml")
		if err != nil {
			return nil, err
		}
		defer file.Close()
	}
	defer file.Close()

	config := Config{}
	err = yaml.NewDecoder(file).Decode(&config)

	return config, err
}
