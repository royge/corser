package ezcors_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/royge/ezcors"
	"gopkg.in/yaml.v2"
)

var corsData = `
dev:
  allowedOrigins:
    - http://127.0.0.1
    - http://devhost
test:
  allowedOrigins:
    - http://127.0.0.2
    - http://testhost
stage:
  allowedOrigins:
    - http://127.0.0.3
    - http://stagehost
prod:
    allowedOrigins:
      - http://127.0.0.4
      - http://prodhost
`

func checkCors(t *testing.T, want, got []string) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want allowed origins `%v`, got `%v`", want, got)
	}
}

func TestConfig_FromString(t *testing.T) {
	config := ezcors.Config{}

	if err := yaml.Unmarshal([]byte(corsData), &config); err != nil {
		t.Fatalf("Error unmarshalling yaml: %v", err)
	}

	checkCors(t, []string{"http://127.0.0.1", "http://devhost"}, config["dev"].AllowedOrigins)
	checkCors(t, []string{"http://127.0.0.2", "http://testhost"}, config["test"].AllowedOrigins)
	checkCors(t, []string{"http://127.0.0.3", "http://stagehost"}, config["stage"].AllowedOrigins)
	checkCors(t, []string{"http://127.0.0.4", "http://prodhost"}, config["prod"].AllowedOrigins)
}

func TestConfig_FromReader(t *testing.T) {
	config := ezcors.Config{}

	r := strings.NewReader(corsData)

	if err := yaml.NewDecoder(r).Decode(&config); err != nil {
		t.Fatalf("Error unmarshalling yaml: %v", err)
	}

	checkCors(t, []string{"http://127.0.0.1", "http://devhost"}, config["dev"].AllowedOrigins)
	checkCors(t, []string{"http://127.0.0.2", "http://testhost"}, config["test"].AllowedOrigins)
	checkCors(t, []string{"http://127.0.0.3", "http://stagehost"}, config["stage"].AllowedOrigins)
	checkCors(t, []string{"http://127.0.0.4", "http://prodhost"}, config["prod"].AllowedOrigins)
}

func TestNewConfig(t *testing.T) {
	config, err := ezcors.NewConfig()
	if err != nil {
		t.Fatalf("Error getting config: %v", err)
	}

	checkCors(t, []string{"http://127.0.0.1", "http://devhost"}, config["dev"].AllowedOrigins)
	checkCors(t, []string{"http://127.0.0.2", "http://testhost"}, config["test"].AllowedOrigins)
	checkCors(t, []string{"http://127.0.0.3", "http://stagehost"}, config["stage"].AllowedOrigins)
	checkCors(t, []string{"http://127.0.0.4", "http://prodhost"}, config["prod"].AllowedOrigins)
}
