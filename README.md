# Exchange Rate API

API client for https://exchangerates.host

## Usage example

`go get -u github.com/FlameInTheDark/erh`

```go
import (
    "log"
    "github.com/FlameInTheDark/erh"
)

func main() {
    conv, err := erh.Convert("USD", "EUR", 30.25, erh.ArgPlaces(2))
    if err != nil {
        log.Fatal(err)
    }
}
```