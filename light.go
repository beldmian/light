package light

type EventHandler = func(Event) error

type Event struct {
	Name    string
	Payload map[string]interface{}
}

func NewEvent(name string) Event {
	return Event{
		Name: name,
	}
}

type Manager struct {
	handlers map[string]EventHandler
}

func (m *Manager) AddHandler(name string, handler EventHandler) {
	m.handlers[name] = handler
}

func (m *Manager) Emit(event Event) error {
	return m.handlers[event.Name](event)
}

func (m *Manager) AsyncEmit(event Event) {
	if err := m.handlers[event.Name](event); err != nil {
		panic(err)
	}
}

func NewManager() Manager {
	return Manager{
		handlers: make(map[string]func(Event) error),
	}
}

var globalManager Manager = NewManager()

func Handle(eventName string, handler EventHandler) {
	globalManager.AddHandler(eventName, handler)
}

func Emit(event Event) error {
	return globalManager.Emit(event)
}

func AsyncEmit(event Event) {
	globalManager.AsyncEmit(event)
}
