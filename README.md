# light
[![godoc](https://pkg.go.dev/badge/github.com/beldmian/light?status.svg)](https://pkg.go.dev/github.com/beldmian/light)
[![codecov](https://codecov.io/gh/beldmian/light/graph/badge.svg?token=TQKZJ8ZLZO)](https://codecov.io/gh/beldmian/light)
[![sourcegrpah](https://sourcegraph.com/github.com/beldmian/light/-/badge.svg)](https://sourcegraph.com/github.com/beldmian/light)

Simple to use, "zero-allocation" event system

## Installation

To install package use command:
```bash
go get github.com/beldmian/light
```

## Example

```go
light.Handle("hello", func(e light.Event) error {
    fmt.Printf("Hello, %s\n", e.Payload["name"])
    return nil
})

light.Emit(light.Event{
    Name:    "hello",
    Payload: map[string]interface{}{"name": "light"},
})
```

## Contribute

All contributions are welcome. Please open issue if you want to provide some new features.