package light

var globalManager Disposer = NewDisposer()

func Handle(eventName string, handler EventHandler) {
	globalManager.AddHandler(eventName, handler)
}

func Emit(event Event) error {
	return globalManager.Emit(event)
}

func AsyncEmit(event Event) {
	globalManager.AsyncEmit(event)
}
