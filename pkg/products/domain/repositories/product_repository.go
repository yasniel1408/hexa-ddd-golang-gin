package repositories

import "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/domain/entities"

type ProductRepository interface {
    GetAll() ([]entities.Product, error)
    GetByID(id uint) (entities.Product, error)
    Create(product entities.Product) error
}