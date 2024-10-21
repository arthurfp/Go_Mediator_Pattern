package mediator

import "fmt"

// ConcreteMediator implements the Mediator interface.
type ConcreteMediator struct {
	componentA *ComponentA
	componentB *ComponentB
}

func (m *ConcreteMediator) SetComponentA(a *ComponentA) {
	m.componentA = a
	a.SetMediator(m)
}

func (m *ConcreteMediator) SetComponentB(b *ComponentB) {
	m.componentB = b
	b.SetMediator(m)
}

// Notify handles communication between components.
func (m *ConcreteMediator) Notify(sender Component, event string) {
	switch sender {
	case m.componentA:
		fmt.Printf("Mediator received event '%s' from ComponentA. Notifying ComponentB.\n", event)
		m.componentB.React(event)
	case m.componentB:
		fmt.Printf("Mediator received event '%s' from ComponentB. Notifying ComponentA.\n", event)
		m.componentA.React(event)
	}
}
