package main

import (
	"mediator-pattern-go/pkg/mediator"
)

func main() {
	concreteMediator := &mediator.ConcreteMediator{}
	componentA := &mediator.ComponentA{}
	componentB := &mediator.ComponentB{}

	concreteMediator.SetComponentA(componentA)
	concreteMediator.SetComponentB(componentB)

	componentA.DoSomething()
	componentB.DoSomethingElse()
}
