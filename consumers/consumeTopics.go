package consumers

import "github.com/rs/zerolog/log"

func ConsumeTopics() {
	err := ConsumeProducts("products")
	if err != nil {
		log.Error().Err(err).Msg("Falha ao consumir produtos")
	}
}
