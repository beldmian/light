package light

var globalManager Disposer = NewDisposer()

// Handle adds handler to event by its name
func Handle(eventName string, handler EventHandler) {
	globalManager.Handle(eventName, handler)
}

// Emit emits event
func Emit(event Event) error {
	return globalManager.Emit(event)
}

// AsyncEmit emits event asynchronously
func AsyncEmit(event Event) {
	globalManager.AsyncEmit(event)
}

// Subscribe adds handler to all emitted events
func Subscribe(subscriber SubscribeHandler) {
	globalManager.Subscribe(subscriber)
}
