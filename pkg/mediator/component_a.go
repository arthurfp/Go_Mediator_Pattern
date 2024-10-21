package mediator

import "fmt"

// ComponentA interacts with the Mediator.
type ComponentA struct {
	BaseComponent
}

func (a *ComponentA) DoSomething() {
	fmt.Println("ComponentA is doing something...")
	a.Notify("event_from_A")
}

func (a *ComponentA) React(event string) {
	fmt.Printf("ComponentA reacted to '%s'.\n", event)
}
