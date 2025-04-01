package consumers

import (
	"github.com/TiagoNora/GoCRUDV2/config/kafkaConfig"
	"github.com/TiagoNora/GoCRUDV2/schemas"
	"github.com/rs/zerolog/log"
	"time"
)

type MessagePayload struct {
	Action  string          `json:"action"`
	Product schemas.Product `json:"product"`
}

func StartProductConsumer(topic string) {
	go func() {
		consumer, err := kafkaConfig.NewConsumer()
		if err != nil {
			log.Error().Err(err).Msg("Falha ao criar consumidor")
			return
		}
		defer consumer.Close()

		if err := consumer.Subscribe([]string{topic}); err != nil {
			log.Error().Err(err).Str("topic", topic).Msg("Falha ao inscrever no tópico")
			return
		}

		timeout := 100 * time.Millisecond

		if err := kafkaConfig.ConsumeMessagesWith[MessagePayload](consumer, timeout, handleProduct); err != nil {
			log.Error().Err(err).Str("topic", topic).Msg("Erro no consumo de mensagens")
		}
	}()

	log.Info().Str("topic", topic).Msg("Consumidor de produtos iniciado em background")
}

func handleProduct(topic string, key []byte, payload MessagePayload) error {
	product := payload.Product

	log.Info().
		Str("topic", topic).
		Str("action", payload.Action).
		Int("productId", int(product.ID)).
		Str("productName", product.Name).
		Float64("price", product.Price).
		Msg("Produto recebido")

	return nil
}

func ConsumeProducts(topic string) error {
	StartProductConsumer(topic)
	return nil
}
