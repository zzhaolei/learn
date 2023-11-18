package kafka

import (
	"context"
	"encoding/json"
	"time"

	"pipeline-manager/config"
	"pipeline-manager/services/http/model"

	"github.com/rs/zerolog/log"
	kafkaGo "github.com/segmentio/kafka-go"
)

type consumer struct {
	reader *kafkaGo.Reader
}

func NewConsumer() *consumer {
	return &consumer{
		reader: kafkaGo.NewReader(
			kafkaGo.ReaderConfig{
				Brokers:          config.Env.KafkaServers,
				GroupID:          config.Env.KafkaGroupId,
				Topic:            config.Env.KafkaTopic,
				ReadBatchTimeout: time.Millisecond * 10,
			},
		),
	}
}

func (c *consumer) Consumer(ctx context.Context) {
	defer c.reader.Close()

	var taskMsg model.TaskMessage

	log.Debug().Msg("start consumer message...")
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Error().Str("err", err.Error()).Msg("fetch kafka message fail")
			continue
		}

		if err = json.Unmarshal(msg.Value, &taskMsg); err != nil {
			log.Error().Str("err", err.Error()).Msg("unmarshal msg to struct fail")
			continue
		}
		log.Debug().Str("msg.key", string(msg.Key)).Interface("msg.value", taskMsg).Send()
	}
}
