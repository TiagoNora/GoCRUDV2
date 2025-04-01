package service

import (
	"github.com/TiagoNora/GoCRUDV2/config/db"
	"github.com/TiagoNora/GoCRUDV2/config/kafkaConfig"
	"github.com/TiagoNora/GoCRUDV2/schemas"
	"github.com/rs/zerolog/log"
	"strconv"
)

type ProductService interface {
	Create(schemas.Product) (schemas.Product, error)
	Update(schemas.Product, schemas.Product) (schemas.Product, error)
	Delete(id string) (schemas.Product, error)
	FindAll() ([]schemas.Product, error)
	FindById(id string) (schemas.Product, error)
}

type productService struct {
	kafkaProducer *kafkaConfig.Producer
}

func (p productService) Update(productOld schemas.Product, productNew schemas.Product) (schemas.Product, error) {
	if err := db.GetDB().Model(&productOld).Updates(&productNew).Error; err != nil {
		return schemas.Product{}, err
	}
	return productOld, nil
}

func (p productService) Delete(id string) (schemas.Product, error) {
	var product schemas.Product
	if err := db.GetDB().Delete(&product, "id = ?", id).Error; err != nil {
		return schemas.Product{}, err
	}
	return product, nil
}

func (p productService) FindById(id string) (schemas.Product, error) {
	var product schemas.Product
	if err := db.GetDB().First(&product, "id = ?", id).Error; err != nil {
		return schemas.Product{}, err
	}
	return product, nil
}

func (p productService) Create(product schemas.Product) (schemas.Product, error) {
	if err := db.GetDB().Create(&product).Error; err != nil {
		return schemas.Product{}, err
	}

	if p.kafkaProducer != nil {
		err := p.kafkaProducer.SendMessage("products", strconv.Itoa(int(product.ID)), map[string]interface{}{
			"action":  "create",
			"product": product,
		})

		log.Info().Msgf("Kafka Producer id: %s", strconv.Itoa(int(product.ID)))

		if err != nil {
			log.Error().Msgf("Error Kafka Producer id: %s", strconv.Itoa(int(product.ID)))
		}
	}

	return product, nil
}

func (p productService) FindAll() ([]schemas.Product, error) {
	var products []schemas.Product
	if err := db.GetDB().Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func NewProductService() ProductService {
	producer, err := kafkaConfig.NewProducer()
	if err != nil {
		log.Error().Msgf("Aviso: Não foi possível conectar ao Kafka: %v", err)
		return &productService{kafkaProducer: nil}
	}

	return &productService{kafkaProducer: producer}
}
