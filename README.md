# ezcors

[![Build Status](https://travis-ci.org/royge/ezcors.svg?branch=master)](https://travis-ci.org/royge/ezcors)
[![Go Report Card](https://goreportcard.com/badge/github.com/royge/ezcors)](https://goreportcard.com/report/github.com/royge/ezcors)

Easy handling of CORS configuration file for Go http handlers.

## How To Use

```
go get github.com/royge/ezcors
```

### Example

The `cors.yml` file.

```yaml
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
```

Usage example:

```go
import (
	"fmt"

	"github.com/royge/ezcors"
)

// ExampleNewConfig shows an example on how to read configuration from a yaml
// file.
// You can check the content of cors.yml file for reference.
func ExampleNewConfig() {
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

	// Output:
	// dev allowed origins: [http://127.0.0.1 http://devhost]
	// dev allowed methods: [GET POST PUT PATCH DELETE]
	// dev allow credentials: false
	// dev debug: true
	// test allowed origins: [http://127.0.0.2 http://testhost]
	// test allowed methods: [POST]
	// test allow credentials: true
	// test debug: true
	// stage allowed origins: [http://127.0.0.3 http://stagehost]
	// prod allowed origins: [http://127.0.0.4 http://prodhost]
}
```
