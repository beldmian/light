# light - event system for go

## Installation

To install package use command:
```bash
go get github.com/beldmian/light
```

## Usage

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