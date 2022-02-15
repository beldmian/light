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
	handlers map[string][]EventHandler
}

func (m *Manager) AddHandler(name string, handler EventHandler) {
	if _, ok := m.handlers[name]; !ok {
		m.handlers[name] = make([]func(Event) error, 0)
	}
	m.handlers[name] = append(m.handlers[name], handler)
}

func (m *Manager) Emit(event Event) error {
	for _, handler := range m.handlers[event.Name] {
		if err := handler(event); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) AsyncEmit(event Event) {
	for _, handler := range m.handlers[event.Name] {
		if err := handler(event); err != nil {
			panic(err)
		}
	}
}

func NewManager() Manager {
	return Manager{
		handlers: make(map[string][]EventHandler),
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
