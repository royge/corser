package ezcors_test

import (
	"fmt"

	"github.com/royge/ezcors"
)

// ExampleNewConfig shows an example on how to read configuration from a yaml
// file.
//
// You can check the content of cors.yml file for reference.
func ExampleNewConfig() {
	// Check for CORS config from cors.yml or config/cors.yml.
	config, err := ezcors.NewConfig()
	if err != nil {
		panic("don't panic")
	}

	fmt.Println("dev allowed origins:", config["dev"].AllowedOrigins)
	fmt.Println("dev allowed methods:", config["dev"].AllowedMethods)
	fmt.Println("dev allow credentials:", config["dev"].AllowCredentials)
	fmt.Println("dev debug:", config["dev"].Debug)

	fmt.Println("test allowed origins:", config["test"].AllowedOrigins)
	fmt.Println("test allowed methods:", config["test"].AllowedMethods)
	fmt.Println("test allow credentials:", config["test"].AllowCredentials)
	fmt.Println("test debug:", config["test"].Debug)

	fmt.Println("stage allowed origins:", config["stage"].AllowedOrigins)
	fmt.Println("prod allowed origins:", config["prod"].AllowedOrigins)

	// Check for CORS config from testdata/cors.yml.
	config, err = ezcors.NewConfig(ezcors.Option{
		Path: "testdata/cors.yml",
	})
	if err != nil {
		panic("don't panic")
	}

	fmt.Println("-----------")
	fmt.Println("Custom Path")
	fmt.Println("test allowed origins:", config["test"].AllowedOrigins)
	fmt.Println("stage allowed origins:", config["stage"].AllowedOrigins)
	fmt.Println("prod allowed origins:", config["prod"].AllowedOrigins)

	// Output:
	// dev allowed origins: [*]
	// dev allowed methods: [GET POST PUT PATCH DELETE]
	// dev allow credentials: false
	// dev debug: true
	// test allowed origins: [http://127.0.0.2 http://testhost.com]
	// test allowed methods: [POST]
	// test allow credentials: true
	// test debug: true
	// stage allowed origins: [http://127.0.0.3 http://stagehost.com]
	// prod allowed origins: [http://127.0.0.4 http://prodhost.com]
	// -----------
	// Custom Path
	// test allowed origins: [http://127.0.0.2 http://testhostcustom.com]
	// stage allowed origins: [http://127.0.0.3 http://stagehostcustom.com]
	// prod allowed origins: [http://127.0.0.4 http://prodhostcustom.com]
}
