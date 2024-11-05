package in

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/application"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/domain/entities"
)

type ProductControllerType struct {
	productService application.IProductService
}

func ProductController(productService application.IProductService) *ProductControllerType {
	return &ProductControllerType{productService}
}

// GetAllProducts obtiene todos los productos
func (h *ProductControllerType) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct obtiene un producto por su ID
func (h *ProductControllerType) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "invalid product id"})
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct crea un nuevo producto
func (h *ProductControllerType) CreateProduct(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	err := h.productService.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}
