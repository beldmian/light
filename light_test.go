package light_test

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/beldmian/light"
)

func TestEmit(t *testing.T) {
	var output bytes.Buffer

	light.Handle("hello", func(e light.Event) error {
		fmt.Fprintf(&output, "Hello, %s", e.Payload["name"])
		return nil
	})

	light.Emit(light.Event{
		Name:    "hello",
		Payload: map[string]interface{}{"name": "light"},
	})

	if output.String() != "Hello, light" {
		t.Errorf("Error in testing manager\nexpexted: \"Hello, light\"\ngot: \"%s\"", output.String())
	}
}

func TestEmit_multipleHandlers(t *testing.T) {
	var out = make([]int, 0)

	light.Handle("a", func(e light.Event) error {
		out = append(out, 1)
		return nil
	})

	light.Handle("a", func(e light.Event) error {
		out = append(out, 2)
		return nil
	})

	light.Emit(light.NewEvent("a"))

	if len(out) != 2 {
		t.Errorf("Error in multiple handlers test")
	}
}

func TestEmitAsync(t *testing.T) {
	var output bytes.Buffer

	light.Handle("hello", func(e light.Event) error {
		fmt.Fprintf(&output, "Hello, %s", e.Payload["name"])
		return nil
	})

	light.AsyncEmit(light.Event{
		Name:    "hello",
		Payload: map[string]interface{}{"name": "light"},
	})

	time.Sleep(time.Second)

	if output.String() != "Hello, light" {
		t.Errorf("Error in testing manager\nexpexted: \"Hello, light\"\ngot: \"%s\"", output.String())
	}
}

func TestNewEvent(t *testing.T) {
	event := light.NewEvent("name")
	if event.Name != "name" {
		t.Errorf("Error in testing manager\nexpexted: \"name\"\ngot: \"%s\"", event.Name)
	}
}

func TestAddSubscriber(t *testing.T) {
	var out = make([]int, 0)
	light.Subscribe(func(e light.Event) error {
		out = append(out, 1)
		return nil
	})

	light.Emit(light.NewEvent("a"))
	light.Emit(light.NewEvent("b"))

	if len(out) != 2 {
		t.Error("Wrong answer in TestAddSubscriber")
	}
}

func BenchmarkEmit(b *testing.B) {
	event := light.NewEvent("bench")
	light.Handle("bench", func(e light.Event) error {
		return nil
	})
	for i := 0; i < b.N; i++ {
		light.Emit(event)
	}
}
