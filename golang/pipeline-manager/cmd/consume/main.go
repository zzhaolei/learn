package main

import (
	"context"

	"golang/pipeline-manager/utils/kafka"
)

func main() {
	consumer := kafka.NewConsumer()
	consumer.Consumer(context.Background())
}
