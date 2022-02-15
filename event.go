package light

type Event struct {
	Name    string
	Payload map[string]interface{}
}

func NewEvent(name string) Event {
	return Event{
		Name: name,
	}
}

type EventHandler = func(Event) error
