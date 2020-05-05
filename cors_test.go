package ezcors_test

import (
	"reflect"
	"testing"

	"github.com/royge/ezcors"
)

func checkCORSValue(t *testing.T, want, got interface{}) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want allowed origins `%v`, got `%v`", want, got)
	}
}

func TestNewConfig(t *testing.T) {
	config, err := ezcors.NewConfig()
	if err != nil {
		t.Fatalf("Error getting config: %v", err)
	}

	// Check allowed origins
	checkCORSValue(t, []string{"*"}, config["dev"].AllowedOrigins)
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
