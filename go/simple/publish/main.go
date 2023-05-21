package main

import (
	"fmt"
	"learn/go/rabbitmq"
	"strconv"
)

func main() {
	mq := rabbitmq.NewRabbitMQSimple("simple")
	defer mq.Destructor()

	for i := 1; i <= 100; i++ {
		mq.PublishSimple([]byte("测试 simple 模式" + strconv.Itoa(i)))
	}
	fmt.Println("Done")
}
