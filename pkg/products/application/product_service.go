package application

import (
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/domain/entities"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/domain/repositories"
)

type IProductService interface {
	GetAllProducts() ([]entities.Product, error)
	GetProductByID(id uint) (entities.Product, error)
	CreateProduct(product entities.Product) error
}

type productService struct {
	productRepo repositories.ProductRepository
}

func ProductService(productRepo repositories.ProductRepository) IProductService {
	return &productService{productRepo}
}

func (s *productService) GetAllProducts() ([]entities.Product, error) {
	return s.productRepo.GetAll()
}

func (s *productService) GetProductByID(id uint) (entities.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *productService) CreateProduct(product entities.Product) error {
	return s.productRepo.Create(product)
}
