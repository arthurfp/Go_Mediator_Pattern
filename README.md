# Mediator Pattern in Go

## Overview
This repository demonstrates the application of the Mediator design pattern in Go. The project showcases how to reduce tight coupling between components by introducing a mediator that manages communication, emphasizing best practices in design patterns.

## Pattern Description
The Mediator pattern centralizes communication between objects to simplify dependencies. Instead of components communicating directly, they interact through a mediator, which coordinates their interactions.

### Key Components
- **Mediator Interface**: Defines a method to notify the mediator of events.
- **ConcreteMediator**: Implements the Mediator interface, managing communication between components.
- **BaseComponent**: Provides a base implementation for components to interact with the mediator.
- **ComponentA and ComponentB**: Concrete components that communicate through the mediator.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the pattern in action.
- **pkg/mediator/**: Houses the Mediator pattern implementation.
    - **mediator.go**: Defines the `Mediator` interface.
    - **base_component.go**: Implements the base component for interaction.
    - **concrete_mediator.go**: Implements the mediator managing communication.
    - **component_a.go**: Implements `ComponentA` for interaction via the mediator.
    - **component_b.go**: Implements `ComponentB` for interaction via the mediator.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone https://github.com/your-username/mediator-pattern-go.git
cd mediator-pattern-go
```

### Example Output
When you run the application, you should see:
```yaml
ComponentA is doing something...
Mediator received event 'event_from_A' from ComponentA. Notifying ComponentB.
ComponentB reacted to 'event_from_A'.
ComponentB is doing something else...
Mediator received event 'event_from_B' from ComponentB. Notifying ComponentA.
ComponentA reacted to 'event_from_B'.

```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp