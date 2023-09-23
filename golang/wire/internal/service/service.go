package service

import "fmt"

type TInterface interface {
	String() string
	Error() string
}

type Message string

func NewMessage(phrase string) (Message, func(), error) {
	return Message(phrase), func() {
		fmt.Println("-----------")
	}, nil
}

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) (Greeter, func()) {
	return Greeter{Message: m}, func() {
		fmt.Println("-----------")
	}
}

func (g Greeter) Greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func NewEvent(g Greeter) (Event, error) {
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
