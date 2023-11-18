package kafka

import (
	"context"
	"fmt"
	"time"

	"pipeline-manager/config"

	kafkaGo "github.com/segmentio/kafka-go"
)

type producer struct {
	writer *kafkaGo.Writer
}

func NewProducer() *producer {
	return &producer{
		writer: kafkaGo.NewWriter(kafkaGo.WriterConfig{
			Brokers:      config.Env.KafkaServers,
			Topic:        config.Env.KafkaTopic,
			Balancer:     &kafkaGo.LeastBytes{},
			BatchTimeout: time.Millisecond * 10,
		}),
	}
}

func (p *producer) SendMessage(ctx context.Context, key []byte, value []byte) error {
	msg := kafkaGo.Message{
		Key:   key,
		Value: value,
	}

	now := time.Now()
	err := p.writer.WriteMessages(ctx, msg)
	fmt.Println(time.Since(now))
	return err
}
