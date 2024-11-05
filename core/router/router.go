package router

import (
	"github.com/gin-gonic/gin"
	fscache "github.com/iqquee/fs-cache"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yasniel1408/hexa-ddd-golang-gin/core/db"
	_ "github.com/yasniel1408/hexa-ddd-golang-gin/docs"
	middlewares "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/common/middlewares"
	productApp "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/application"
	productIn "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/infrastructure/in"
	productOut "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/products/infrastructure/out"
	userApp "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/application"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/in_adapters/http"
	cache "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/out_adapters/cache"
	userOut "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/users/infrastructure/out_adapters/sql"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	jwtKey := []byte("my_secret_key")

	// Inicializar la base de datos
	database := db.InitDB()

	// cache
	var myfscache fscache.Operations = fscache.New()
	mycache := cache.CacheUsersAdapter(&myfscache)

	// Middleware de autenticación
	authMiddleware := middlewares.AuthMiddleware(jwtKey)

	// Repositorios
	userRepo := userOut.UserRepository(database)
	productRepo := productOut.ProductRepository(database)

	// Servicios
	authService := userApp.AuthService(userRepo, jwtKey)
	userService := userApp.UserService(userRepo, mycache)
	productService := productApp.ProductService(productRepo)

	// Controller
	authController := http.AuthController(authService)
	userController := http.UserController(userService)
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
