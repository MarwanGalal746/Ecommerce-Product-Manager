package defaultProductService

import "Ecommerce-Product-Manager/pkg/domain/repositories/productRepository"

func NewDefaultProductService(repo productRepository.ProductRepositoryDb) DefaultProductService {
	return DefaultProductService{Repo: repo}
}
