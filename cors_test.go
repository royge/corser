package ezcors_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/royge/ezcors"
	"syreclabs.com/go/faker"
)

func checkCORSValue(t *testing.T, want, got interface{}) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want allowed origins `%v`, got `%v`", want, got)
	}
}

func TestNewConfig(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		config, err := ezcors.NewConfig()
		if err != nil {
			t.Fatalf("Error getting config: %v", err)
		}

		// Check allowed origins
		checkCORSValue(t, []string{"*"}, config["dev"].AllowedOrigins)
		checkCORSValue(t, []string{"http://127.0.0.2", "http://testhost.com"}, config["test"].AllowedOrigins)
		checkCORSValue(t, []string{"http://127.0.0.3", "http://stagehost.com"}, config["stage"].AllowedOrigins)
		checkCORSValue(t, []string{"http://127.0.0.4", "http://prodhost.com"}, config["prod"].AllowedOrigins)

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
	})

	t.Run("custom path", func(t *testing.T) {
		path := "testdata/cors.yml"
		config, err := ezcors.NewConfig(ezcors.Option{
			Path: path,
		})
		if err != nil {
			t.Fatalf("Error getting config: %v", err)
		}

		// Check allowed origins
		checkCORSValue(t, []string{"*"}, config["dev"].AllowedOrigins)
		checkCORSValue(t, []string{"http://127.0.0.2", "http://testhostcustom.com"}, config["test"].AllowedOrigins)
		checkCORSValue(t, []string{"http://127.0.0.3", "http://stagehostcustom.com"}, config["stage"].AllowedOrigins)
		checkCORSValue(t, []string{"http://127.0.0.4", "http://prodhostcustom.com"}, config["prod"].AllowedOrigins)

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
	})
}

func TestAllowedOrigin(t *testing.T) {
	emptyHeader := http.Header{}

	withDiffOriginHeader := http.Header{}
	withDiffOriginHeader.Add("Origin", faker.Internet().Url())

	allowedOrigins := []string{
		faker.Internet().Url(),
		faker.Internet().Url(),
		faker.Internet().Url(),
	}

	withOriginHeader := http.Header{}
	withOriginHeader.Add("Origin", allowedOrigins[1])

	tests := []struct {
		name   string
		header http.Header
		input  []string
		want   string
	}{
		{
			"no origin",
			emptyHeader,
			allowedOrigins,
			allowedOrigins[0],
		},
		{
			"with different origin",
			withDiffOriginHeader,
			allowedOrigins,
			allowedOrigins[0],
		},
		{
			"with origin",
			withOriginHeader,
			allowedOrigins,
			withOriginHeader.Get("Origin"),
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := ezcors.AllowedOrigin(tc.header, tc.input)

			if got != tc.want {
				t.Errorf("AllowedOrigin() want (%v), got (%v)", tc.want, got)
			}
		})
	}
}
