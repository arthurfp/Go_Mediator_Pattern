package mediator

import "fmt"

// ComponentB interacts with the Mediator.
type ComponentB struct {
	BaseComponent
}

func (b *ComponentB) DoSomethingElse() {
	fmt.Println("ComponentB is doing something else...")
	b.Notify("event_from_B")
}

func (b *ComponentB) React(event string) {
	fmt.Printf("ComponentB reacted to '%s'.\n", event)
}
