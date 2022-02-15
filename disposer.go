package light

type Disposer struct {
	handlers map[string][]EventHandler
}

func (m *Disposer) AddHandler(name string, handler EventHandler) {
	if _, ok := m.handlers[name]; !ok {
		m.handlers[name] = make([]func(Event) error, 0)
	}
	m.handlers[name] = append(m.handlers[name], handler)
}

func (m *Disposer) Emit(event Event) error {
	for _, handler := range m.handlers[event.Name] {
		if err := handler(event); err != nil {
			return err
		}
	}
	return nil
}

func (m *Disposer) AsyncEmit(event Event) {
	for _, handler := range m.handlers[event.Name] {
		if err := handler(event); err != nil {
			panic(err)
		}
	}
}

func NewDisposer() Disposer {
	return Disposer{
		handlers: make(map[string][]EventHandler),
	}
}
