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
	fmt.Println("test allowed origins:", config["test"].AllowedOrigins)
	fmt.Println("stage allowed origins:", config["stage"].AllowedOrigins)
	fmt.Println("prod allowed origins:", config["prod"].AllowedOrigins)

	// Output:
	// dev allowed origins: [http://127.0.0.1 http://devhost]
	// test allowed origins: [http://127.0.0.2 http://testhost]
	// stage allowed origins: [http://127.0.0.3 http://stagehost]
	// prod allowed origins: [http://127.0.0.4 http://prodhost]
}
```
