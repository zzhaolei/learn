package main

import (
	"fmt"
	"strconv"

	"golang/rabbitmq"
)

func main() {
	mq := rabbitmq.NewRabbitMQPubSub("testexchange")
	for i := 1; i <= 1; i++ {
		mq.PublishPub([]byte("测试 exchange 功能" + strconv.Itoa(i)))
	}
	fmt.Println("Done")
}
