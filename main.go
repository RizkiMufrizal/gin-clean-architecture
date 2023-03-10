package main

import (
	"github.com/RizkiMufrizal/gin-clean-architecture/client/restclient"
	"github.com/RizkiMufrizal/gin-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gin-clean-architecture/controller"
	_ "github.com/RizkiMufrizal/gin-clean-architecture/docs"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/RizkiMufrizal/gin-clean-architecture/middleware"
	repository "github.com/RizkiMufrizal/gin-clean-architecture/repository/impl"
	service "github.com/RizkiMufrizal/gin-clean-architecture/service/impl"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Clean Architecture
// @version 1.0.0
// @description Baseline project using Gin
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email gin@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9999
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description Authorization For JWT
func main() {
	//setup configuration
	config := configuration.New()
	database := configuration.NewDatabase(config)
	redis := configuration.NewRedis(config)

	//repository
	productRepository := repository.NewProductRepositoryImpl(database)
	transactionRepository := repository.NewTransactionRepositoryImpl(database)
	transactionDetailRepository := repository.NewTransactionDetailRepositoryImpl(database)
	userRepository := repository.NewUserRepositoryImpl(database)

	//rest client
	httpBinRestClient := restclient.NewHttpBinRestClient()

	//service
	productService := service.NewProductServiceImpl(&productRepository, redis)
	transactionService := service.NewTransactionServiceImpl(&transactionRepository)
	transactionDetailService := service.NewTransactionDetailServiceImpl(&transactionDetailRepository)
	userService := service.NewUserServiceImpl(&userRepository)
	httpBinService := service.NewHttpBinServiceImpl(&httpBinRestClient)

	//controller
	productController := controller.NewProductController(&productService, config)
	transactionController := controller.NewTransactionController(&transactionService, config)
	transactionDetailController := controller.NewTransactionDetailController(&transactionDetailService, config)
	userController := controller.NewUserController(&userService, config)
	httpBinController := controller.NewHttpBinController(&httpBinService)

	//setup fiber
	gin.SetMode(gin.DebugMode)
	app := gin.Default()
	app.Use(gin.CustomRecovery(exception.ErrorHandler))
	app.Use(middleware.CORSMiddleware())

	//routing
	productController.Route(app)
	transactionController.Route(app)
	transactionDetailController.Route(app)
	userController.Route(app)
	httpBinController.Route(app)

	//swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//start app
	err := app.Run(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
