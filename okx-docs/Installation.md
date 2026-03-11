# Installation

## Requirements

- Go 1.21 or higher
- Git (for installation via `go get`)

## Install via go get

```bash
go get github.com/tigusigalpa/okx-go
```

## Install specific version

```bash
go get github.com/tigusigalpa/okx-go@v0.1.0
```

## Verify installation

Create a simple test file:

```go
package main

import (
    "fmt"
    "github.com/tigusigalpa/okx-go"
)

func main() {
    client := okx.NewRestClient("", "", "")
    fmt.Printf("Client created: %T\n", client)
}
```

Run it:

```bash
go run main.go
```

You should see: `Client created: *okx.RestClient`

## Dependencies

The library has minimal dependencies:

- `github.com/gorilla/websocket` v1.5.1 - WebSocket support
- `github.com/stretchr/testify` v1.8.4 - Testing only

All dependencies are automatically installed via `go get`.

## Update to latest version

```bash
go get -u github.com/tigusigalpa/okx-go
```

## Next Steps

- [Getting Started](Getting-Started)
- [Authentication](Authentication)
