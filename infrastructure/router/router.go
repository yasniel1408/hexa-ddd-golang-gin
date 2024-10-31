package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/yasniel1408/hexa-ddd-golang-gin/docs"
	"github.com/yasniel1408/hexa-ddd-golang-gin/infrastructure/db"
	authIn "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/auth/adapter/in"
	authApp "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/auth/application"
	middlewares "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/common/middlewares"
	productIn "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/adapter/in"
	productOut "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/adapter/out"
	productApp "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/application"
	userIn "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/adapter/in"
	userOut "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/adapter/out"
	userApp "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/application"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	jwtKey := []byte("my_secret_key")

	// Inicializar la base de datos
	database := db.InitDB()

	// Middleware de autenticación
	authMiddleware := middlewares.AuthMiddleware(jwtKey)

	// Repositorios
	userRepo := userOut.NewUserRepository(database)
	productRepo := productOut.NewProductRepository(database)

	// Servicios
	authService := authApp.NewAuthService(userRepo, jwtKey)
	userService := userApp.NewUserService(userRepo)
	productService := productApp.NewProductService(productRepo)

	// Handlers
	authHandler := authIn.NewAuthHandler(authService)
	userHandler := userIn.NewUserHandler(userService)
	productHandler := productIn.NewProductHandler(productService)

	// Rutas
	api := r.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		api.GET("/users/:id", userHandler.GetUser)

		api.GET("/products", productHandler.GetAllProducts)
		api.GET("/products/:id", productHandler.GetProduct)
		api.POST("/products", authMiddleware, productHandler.CreateProduct)
	}

	// Documentación Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
