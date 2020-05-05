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
  allowCredentials: false
  allowedMethods:
    - GET
    - POST
    - PUT
    - PATCH
    - DELETE
  debug: true
test:
  allowedOrigins:
    - http://127.0.0.2
    - http://testhost
  allowCredentials: true
  allowedMethods:
    - POST
  debug: true
stage:
  allowedOrigins:
    - http://127.0.0.3
    - http://stagehost
  allowCredentials: true
  allowedMethods:
    - POST
  debug: false
prod:
  allowedOrigins:
    - http://127.0.0.4
    - http://prodhost
  allowCredentials: true
  allowedMethods:
    - POST
  debug: false
`

func checkCORSValue(t *testing.T, want, got interface{}) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want allowed origins `%v`, got `%v`", want, got)
	}
}

func TestConfig_FromString(t *testing.T) {
	config := ezcors.Config{}

	if err := yaml.Unmarshal([]byte(corsData), &config); err != nil {
		t.Fatalf("Error unmarshalling yaml: %v", err)
	}

	// Check allowed origins
	checkCORSValue(t, []string{"http://127.0.0.1", "http://devhost"}, config["dev"].AllowedOrigins)
	checkCORSValue(t, []string{"http://127.0.0.2", "http://testhost"}, config["test"].AllowedOrigins)
	checkCORSValue(t, []string{"http://127.0.0.3", "http://stagehost"}, config["stage"].AllowedOrigins)
	checkCORSValue(t, []string{"http://127.0.0.4", "http://prodhost"}, config["prod"].AllowedOrigins)

	// Check allowed methods
	checkCORSValue(t, []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, config["dev"].AllowedMethods)
	checkCORSValue(t, []string{"POST"}, config["test"].AllowedMethods)
	checkCORSValue(t, []string{"POST"}, config["stage"].AllowedMethods)
	checkCORSValue(t, []string{"POST"}, config["prod"].AllowedMethods)

	// Check allow credentials
	checkCORSValue(t, false, config["dev"].AllowCredentials)
	checkCORSValue(t, true, config["test"].AllowCredentials)
	checkCORSValue(t, true, config["stage"].AllowCredentials)
	checkCORSValue(t, true, config["prod"].AllowCredentials)

	// Check debug
	checkCORSValue(t, true, config["dev"].Debug)
	checkCORSValue(t, true, config["test"].Debug)
	checkCORSValue(t, false, config["stage"].Debug)
	checkCORSValue(t, false, config["prod"].Debug)
}

func TestConfig_FromReader(t *testing.T) {
	config := ezcors.Config{}

	r := strings.NewReader(corsData)

	if err := yaml.NewDecoder(r).Decode(&config); err != nil {
		t.Fatalf("Error unmarshalling yaml: %v", err)
	}

	// Check allowed origins
	checkCORSValue(t, []string{"http://127.0.0.1", "http://devhost"}, config["dev"].AllowedOrigins)
	checkCORSValue(t, []string{"http://127.0.0.2", "http://testhost"}, config["test"].AllowedOrigins)
	checkCORSValue(t, []string{"http://127.0.0.3", "http://stagehost"}, config["stage"].AllowedOrigins)
	checkCORSValue(t, []string{"http://127.0.0.4", "http://prodhost"}, config["prod"].AllowedOrigins)

	// Check allowed methods
	checkCORSValue(t, []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, config["dev"].AllowedMethods)
	checkCORSValue(t, []string{"POST"}, config["test"].AllowedMethods)
	checkCORSValue(t, []string{"POST"}, config["stage"].AllowedMethods)
	checkCORSValue(t, []string{"POST"}, config["prod"].AllowedMethods)

	// Check allow credentials
	checkCORSValue(t, false, config["dev"].AllowCredentials)
	checkCORSValue(t, true, config["test"].AllowCredentials)
	checkCORSValue(t, true, config["stage"].AllowCredentials)
	checkCORSValue(t, true, config["prod"].AllowCredentials)

	// Check debug
	checkCORSValue(t, true, config["dev"].Debug)
	checkCORSValue(t, true, config["test"].Debug)
	checkCORSValue(t, false, config["stage"].Debug)
	checkCORSValue(t, false, config["prod"].Debug)
}

func TestNewConfig(t *testing.T) {
	config, err := ezcors.NewConfig()
	if err != nil {
		t.Fatalf("Error getting config: %v", err)
	}

	// Check allowed origins
	checkCORSValue(t, []string{"http://127.0.0.1", "http://devhost"}, config["dev"].AllowedOrigins)
	checkCORSValue(t, []string{"http://127.0.0.2", "http://testhost"}, config["test"].AllowedOrigins)
	checkCORSValue(t, []string{"http://127.0.0.3", "http://stagehost"}, config["stage"].AllowedOrigins)
	checkCORSValue(t, []string{"http://127.0.0.4", "http://prodhost"}, config["prod"].AllowedOrigins)

	// Check allowed methods
	checkCORSValue(t, []string{"GET", "POST", "PUT", "PATCH", "DELETE"}, config["dev"].AllowedMethods)
	checkCORSValue(t, []string{"POST"}, config["test"].AllowedMethods)
	checkCORSValue(t, []string{"POST"}, config["stage"].AllowedMethods)
	checkCORSValue(t, []string{"POST"}, config["prod"].AllowedMethods)

	// Check allow credentials
	checkCORSValue(t, false, config["dev"].AllowCredentials)
	checkCORSValue(t, true, config["test"].AllowCredentials)
	checkCORSValue(t, true, config["stage"].AllowCredentials)
	checkCORSValue(t, true, config["prod"].AllowCredentials)

	// Check debug
	checkCORSValue(t, true, config["dev"].Debug)
	checkCORSValue(t, true, config["test"].Debug)
	checkCORSValue(t, false, config["stage"].Debug)
	checkCORSValue(t, false, config["prod"].Debug)
}
