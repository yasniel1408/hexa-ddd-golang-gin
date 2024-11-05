package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yasniel1408/hexa-ddd-golang-gin/core/db"
	_ "github.com/yasniel1408/hexa-ddd-golang-gin/docs"
	middlewares "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/common/middlewares"
	productApp "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/application"
	productIn "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/infrastructure/in"
	productOut "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/infrastructure/out"
	userApp "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/application"
	userIn "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/in"
	userOut "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/out"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	jwtKey := []byte("my_secret_key")

	// Inicializar la base de datos
	database := db.InitDB()

	// Middleware de autenticación
	authMiddleware := middlewares.AuthMiddleware(jwtKey)

	// Repositorios
	userRepo := userOut.UserRepository(database)
	productRepo := productOut.ProductRepository(database)

	// Servicios
	authService := userApp.NewAuthService(userRepo, jwtKey)
	userService := userApp.NewUserService(userRepo)
	productService := productApp.NewProductService(productRepo)

	// Handlers
	authController := userIn.AuthController(authService)
	userController := userIn.UserController(userService)
	productController := productIn.ProductController(productService)

	// Rutas
	api := router.Group("/api")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
		api.GET("/users/:id", userController.GetUser)

		api.GET("/products", productController.GetAllProducts)
		api.GET("/products/:id", productController.GetProduct)
		api.POST("/products", authMiddleware, productController.CreateProduct)
	}

	// Documentación Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
