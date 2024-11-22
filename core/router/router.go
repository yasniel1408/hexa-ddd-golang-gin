package router

import (
	"github.com/gin-gonic/gin"
	fscache "github.com/iqquee/fs-cache"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yasniel1408/hexa-ddd-golang-gin/core/db"
	_ "github.com/yasniel1408/hexa-ddd-golang-gin/docs"
	middlewares "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/common/middlewares"
	userApp "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/application"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/domain"
	"github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/input_adapters/http"
	cache "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/cache"
	userOut "github.com/yasniel1408/hexa-ddd-golang-gin/pkg/identity/infrastructure/output_adapters/sql"
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

	// ADAPTERS
	// Out
	userRepo := userOut.UserSqliteAdapter(database)

	// Servicios
	userFactory := domain.UserFactory{}
	authService := userApp.AuthService(userRepo, jwtKey, userFactory)
	userService := userApp.UserService(userRepo, mycache)

	// Http
	authController := http.AuthController(authService)
	userController := http.UserController(userService)

	// Rutas
	api := router.Group("/api/identity")
	{
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)
		api.GET("/users/:id", authMiddleware, userController.GetUser)

	}

	// Documentación Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
