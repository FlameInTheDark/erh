[![Go Report Card](https://goreportcard.com/badge/github.com/FlameInTheDark/erh)](https://goreportcard.com/report/github.com/FlameInTheDark/erh) [![Go Reference](https://pkg.go.dev/badge/github.com/FlameInTheDark/erh.svg)](https://pkg.go.dev/github.com/FlameInTheDark/erh)

# Exchange Rate API

API client for https://exchangerate.host

## Usage example

`go get -u github.com/FlameInTheDark/erh`

```go
package main
import (
    "log"
    "github.com/FlameInTheDark/erh"
)

func main() {
	c := erh.NewClient()
    conv, err := c.Convert("USD", "EUR", 30.25, erh.ArgPlaces(2))
    if err != nil {
        log.Fatal(err)
    }
}
```

## Methods

### [Convert](https://pkg.go.dev/github.com/FlameInTheDark/erh#Convert)
### [Historical](https://pkg.go.dev/github.com/FlameInTheDark/erh#Historical)
### [Latest](https://pkg.go.dev/github.com/FlameInTheDark/erh#Latest)
### [Symbols](https://pkg.go.dev/github.com/FlameInTheDark/erh#Symbols)
### [Timeseries](https://pkg.go.dev/github.com/FlameInTheDark/erh#TimeSeries)
