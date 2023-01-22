package controller

import (
	"bytes"
	"encoding/json"
	"github.com/RizkiMufrizal/gin-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/RizkiMufrizal/gin-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
	repository "github.com/RizkiMufrizal/gin-clean-architecture/repository/impl"
	service "github.com/RizkiMufrizal/gin-clean-architecture/service/impl"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http/httptest"
)

func createTestApp() *gin.Engine {
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

	return app
}

// setup configuration
var config = configuration.New("../.env.test")
var database = configuration.NewDatabase(config)
var redis = configuration.NewRedis(config)

// repository
var productRepository = repository.NewProductRepositoryImpl(database)
var transactionRepository = repository.NewTransactionRepositoryImpl(database)
var transactionDetailRepository = repository.NewTransactionDetailRepositoryImpl(database)
var userRepository = repository.NewUserRepositoryImpl(database)

// service
var productService = service.NewProductServiceImpl(&productRepository, redis)
var transactionService = service.NewTransactionServiceImpl(&transactionRepository)
var transactionDetailService = service.NewTransactionDetailServiceImpl(&transactionDetailRepository)
var userService = service.NewUserServiceImpl(&userRepository)

// controller
var productController = NewProductController(&productService, config)
var transactionController = NewTransactionController(&transactionService, config)
var transactionDetailController = NewTransactionDetailController(&transactionDetailService, config)
var userController = NewUserController(&userService, config)

var appTest = createTestApp()

func authenticationCreate() map[string]interface{} {
	userRepository.DeleteAll()

	password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	exception.PanicLogging(err)
	roles := []string{"ROLE_ADMIN", "ROLE_USER"}
	userRepository.Create("admin", string(password), roles)

	userModel := model.UserModel{
		Username: "admin",
		Password: "admin",
	}

	userRequestBody, _ := json.Marshal(userModel)

	userRequest := httptest.NewRequest("POST", "/v1/api/authentication", bytes.NewBuffer(userRequestBody))
	userRequest.Header.Set("Content-Type", "application/json")
	userRequest.Header.Set("Accept", "application/json")

	userResponse := httptest.NewRecorder()

	appTest.ServeHTTP(userResponse, userRequest)

	userResponseBody, _ := io.ReadAll(userResponse.Body)
	userWebResponse := model.GeneralResponse{}
	_ = json.Unmarshal(userResponseBody, &userWebResponse)

	userJsonData, _ := json.Marshal(userWebResponse.Data)

	tokenResponse := map[string]interface{}{}
	_ = json.Unmarshal(userJsonData, &tokenResponse)

	return tokenResponse
}
