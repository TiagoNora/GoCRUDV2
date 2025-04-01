package kafkaConfig

import (
	"encoding/json"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"
)

type Consumer struct {
	consumer *kafka.Consumer
}

func NewConsumer() (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("KAFKA_SERVERS"),
		"auto.offset.reset":  os.Getenv("KAFKA_AUTO_OFFSET_RESET"),
		"enable.auto.commit": os.Getenv("KAFKA_ENABLE_AUTO_COMMIT"),
		"group.id":           os.Getenv("KAFKA_GROUP_ID"),
	})

	if err != nil {
		return nil, err
	}

	return &Consumer{consumer: c}, nil
}

func (c *Consumer) Subscribe(topics []string) error {
	return c.consumer.SubscribeTopics(topics, nil)
}

func (c *Consumer) ConsumeMessages(timeout time.Duration, messageHandler func(topic string, key []byte, value []byte) error) error {
	for {
		msg, err := c.consumer.ReadMessage(timeout)
		if err != nil {
			if err.(kafka.Error).Code() == kafka.ErrTimedOut {
				continue
			}
			log.Error().Err(err).Msg("Erro ao ler mensagem")
			return err
		}

		if err := messageHandler(*msg.TopicPartition.Topic, msg.Key, msg.Value); err != nil {
			log.Error().Err(err).
				Str("topic", *msg.TopicPartition.Topic).
				Msg("Erro ao processar mensagem")
		}
	}
}

func ConsumeMessagesWith[T any](c *Consumer, timeout time.Duration, messageHandler func(topic string, key []byte, value T) error) error {
	return c.ConsumeMessages(timeout, func(topic string, key []byte, value []byte) error {
		var decodedValue T
		if err := json.Unmarshal(value, &decodedValue); err != nil {
			return err
		}
		return messageHandler(topic, key, decodedValue)
	})
}

func (c *Consumer) Close() {
	c.consumer.Close()
}
