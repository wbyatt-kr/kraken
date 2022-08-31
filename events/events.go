package events

type Event struct {
	handlers []func()
}

func (e *Event) On(f func()) {
	e.handlers = append(e.handlers, f)
}

func (e *Event) Emit() {
	for _, handler := range e.handlers {
		handler()
	}
}