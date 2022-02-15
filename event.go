package light

// Event struct provide access to event data
type Event struct {
	// Name is name of event
	Name string
	// Payload is data to supply into event handler
	Payload map[string]interface{}
}

// NewEvent creates event with supplied name
func NewEvent(name string) Event {
	return Event{
		Name: name,
	}
}

// EventHandler provide type for event handleing
type EventHandler = func(Event) error
