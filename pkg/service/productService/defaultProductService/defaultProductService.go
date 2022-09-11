package defaultProductService

import (
	"Ecommerce-Product-Manager/pkg/config"
	"Ecommerce-Product-Manager/pkg/domain/repositories/productRepository"
)

type DefaultProductService struct {
	Repo productRepository.ProductRepositoryDb
}

func NewDefaultProductService(repo productRepository.ProductRepositoryDb) DefaultProductService {
	config.Logger.Info("creating new instance from default product service")
	return DefaultProductService{Repo: repo}
}
