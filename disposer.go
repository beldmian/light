package light

type Disposer struct {
	handlers    map[string][]EventHandler
	subscribers []SubscribeHandler
}

func (m *Disposer) AddHandler(name string, handler EventHandler) {
	if _, ok := m.handlers[name]; !ok {
		m.handlers[name] = make([]func(Event) error, 0)
	}
	m.handlers[name] = append(m.handlers[name], handler)
}

func (m *Disposer) AddSubscriber(subscriber SubscribeHandler) {
	m.subscribers = append(m.subscribers, subscriber)
}

func (m *Disposer) Emit(event Event) error {
	for _, subscriber := range m.subscribers {
		if err := subscriber(event); err != nil {
			return err
		}
	}
	if _, ok := m.handlers[event.Name]; !ok {
		return nil
	}
	for _, handler := range m.handlers[event.Name] {
		if err := handler(event); err != nil {
			return err
		}
	}
	return nil
}

func (m *Disposer) AsyncEmit(event Event) {
	for _, subscriber := range m.subscribers {
		go func(subscriber SubscribeHandler) {
			if err := subscriber(event); err != nil {
				panic(err)
			}
		}(subscriber)
	}
	if _, ok := m.handlers[event.Name]; !ok {
		return
	}
	for _, handler := range m.handlers[event.Name] {
		go func(handler EventHandler) {
			if err := handler(event); err != nil {
				panic(err)
			}
		}(handler)
	}
}

func NewDisposer() Disposer {
	return Disposer{
		handlers: make(map[string][]EventHandler),
	}
}
