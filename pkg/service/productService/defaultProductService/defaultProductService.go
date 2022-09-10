package defaultProductService

import (
	"Ecommerce-Product-Manager/pkg/domain/repositories/productRepository"
)

type DefaultProductService struct {
	Repo productRepository.ProductRepositoryDb
}

func NewDefaultProductService(repo productRepository.ProductRepositoryDb) DefaultProductService {
	return DefaultProductService{Repo: repo}
}
