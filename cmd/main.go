package main

import (
	"github.com/yasniel1408/hexa-ddd-golang-gin/infrastructure/router"
)

// @title Hexagonal DDD Golang Gin API
// @version 1.0
// @description API para la gestión de autenticación, usuarios y productos
// @termsOfService http://swagger.io/terms/

// @contact.name Soporte
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
