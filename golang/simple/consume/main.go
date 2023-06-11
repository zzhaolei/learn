package main

import (
	"fmt"

	"golang/rabbitmq"
)

func main() {
	mq := rabbitmq.NewRabbitMQSimple("simple")
	defer mq.Destructor()
	mq.ConsumeSimple()
	fmt.Println("done")
}
