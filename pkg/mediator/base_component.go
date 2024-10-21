package mediator

// Component defines the interface for components that interact with the Mediator.
type Component interface {
	SetMediator(mediator Mediator)
	Notify(event string)
}

// BaseComponent provides a default implementation for components.
type BaseComponent struct {
	mediator Mediator
}

// SetMediator assigns a Mediator to the component.
func (b *BaseComponent) SetMediator(mediator Mediator) {
	b.mediator = mediator
}

// Notify sends an event to the Mediator.
func (b *BaseComponent) Notify(event string) {
	if b.mediator != nil {
		b.mediator.Notify(b, event)
	}
}
