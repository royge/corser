package ezcors

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

// CORS defines the supported properties.
type CORS struct {
	AllowedOrigins   []string `yaml:"allowedOrigins"`
	AllowCredentials bool     `yaml:"allowCredentials"`
	AllowedMethods   []string `yaml:"allowedMethods"`
	Debug            bool     `yaml:"debug"`
}

// Config defines CORS configuration for every environment.
type Config map[string]CORS

// NewConfig decodes config file and returns CORS Config.
// The function will look for cors.yml file from the current directory. If
// nothing can found it will try to look into the config directory for possible
// cors.yml file.
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
