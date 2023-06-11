package main

import (
	"strconv"

	"golang/rabbitmq"
)

func main() {
	one := rabbitmq.NewRabbitMQRouting("texchange", "texchange_one")
	two := rabbitmq.NewRabbitMQRouting("texchange", "texchange_two")
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			two.PublishRouting([]byte("测试 routing 模式 two " + strconv.Itoa(i)))
		} else {
			one.PublishRouting([]byte("测试 routing 模式 one " + strconv.Itoa(i)))
		}
	}
}
