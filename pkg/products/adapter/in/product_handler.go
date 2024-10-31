package in

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/application"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/domain/entities"
)

type ProductHandler struct {
	productService application.ProductService
}

func NewProductHandler(productService application.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

// GetAllProducts obtiene todos los productos
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct obtiene un producto por su ID
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct crea un nuevo producto
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entities.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.productService.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}
