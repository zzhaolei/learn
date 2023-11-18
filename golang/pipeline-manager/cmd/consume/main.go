package main

import (
	"context"

	"pipeline-manager/utils/kafka"
)

func main() {
	consumer := kafka.NewConsumer()
	consumer.Consumer(context.Background())
}
