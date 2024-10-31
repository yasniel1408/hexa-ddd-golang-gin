package out

import (
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/domain/entities"
    "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/domain/repositories"
    "gorm.io/gorm"
)

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repositories.ProductRepository {
    return &productRepository{db}
}

func (r *productRepository) GetAll() ([]entities.Product, error) {
    var products []entities.Product
    result := r.db.Find(&products)
    return products, result.Error
}

func (r *productRepository) GetByID(id uint) (entities.Product, error) {
    var product entities.Product
    result := r.db.First(&product, id)
    return product, result.Error
}

func (r *productRepository) Create(product entities.Product) error {
    return r.db.Create(&product).Error
}