package main

import (
	"fmt"
	"learn/go/rabbitmq"
)

func main() {
	mq := rabbitmq.NewRabbitMQSimple("simple")
	defer mq.Destructor()
	mq.ConsumeSimple()
	fmt.Println("done")
}
