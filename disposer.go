package light

import "sync"

// Disposer provide accest to emiting and handling events
type Disposer struct {
	handlers         map[string][]EventHandler
	subscribers      []SubscribeHandler
	subscribersMutex sync.RWMutex
}

// Handle adds handler to event by its name
func (m *Disposer) Handle(name string, handler EventHandler) {
	if _, ok := m.handlers[name]; !ok {
		m.handlers[name] = make([]func(Event) error, 0)
	}
	m.handlers[name] = append(m.handlers[name], handler)
}

// Subscribe adds handler to all emitted events
func (m *Disposer) Subscribe(subscriber SubscribeHandler) {
	m.subscribersMutex.Lock()
	m.subscribers = append(m.subscribers, subscriber)
	m.subscribersMutex.Unlock()
}

// Emit emits event
func (m *Disposer) Emit(event Event) error {
	m.subscribersMutex.RLock()
	defer m.subscribersMutex.RUnlock()
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

// AsyncEmit emits event asynchronously
func (m *Disposer) AsyncEmit(event Event) {
	m.subscribersMutex.RLock()
	defer m.subscribersMutex.RUnlock()
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

// NewDisposer creates disposer
func NewDisposer() Disposer {
	return Disposer{
		handlers: make(map[string][]EventHandler),
	}
}
