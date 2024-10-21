package mediator

// Mediator defines the interface for communication between components.
type Mediator interface {
	Notify(sender Component, event string)
}
