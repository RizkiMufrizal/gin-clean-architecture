package controller

import (
	"github.com/RizkiMufrizal/gin-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/RizkiMufrizal/gin-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
	"github.com/RizkiMufrizal/gin-clean-architecture/service"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service.TransactionService
	configuration.Config
}

func NewTransactionController(transactionService *service.TransactionService, config configuration.Config) *TransactionController {
	return &TransactionController{TransactionService: *transactionService, Config: config}
}

func (controller TransactionController) Route(app *gin.Engine) {
	app.POST("/v1/api/transaction", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.Create)
	app.DELETE("/v1/api/transaction/:id", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.Delete)
	app.GET("/v1/api/transaction/:id", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.FindById)
	app.GET("/v1/api/transaction", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.FindAll)
}

// Create func create transaction.
// @Description create transaction.
// @Summary create transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body model.TransactionCreateUpdateModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/transaction [post]
func (controller TransactionController) Create(c *gin.Context) {
	var request model.TransactionCreateUpdateModel
	err := c.BindJSON(&request)
	exception.PanicLogging(err)

	response := controller.TransactionService.Create(c.Copy(), request)
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// Delete func delete one exists transaction.
// @Description delete one exists transaction.
// @Summary delete one exists transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/transaction/{id} [delete]
func (controller TransactionController) Delete(c *gin.Context) {
	id := c.Param("id")

	controller.TransactionService.Delete(c.Copy(), id)
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

// FindById func gets one exists transaction.
// @Description Get one exists transaction.
// @Summary get one exists transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/transaction/{id} [get]
func (controller TransactionController) FindById(c *gin.Context) {
	id := c.Param("id")

	result := controller.TransactionService.FindById(c.Copy(), id)
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// FindAll func gets all exists transaction.
// @Description Get all exists transaction.
// @Summary get all exists transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/transaction [get]
func (controller TransactionController) FindAll(c *gin.Context) {
	result := controller.TransactionService.FindAll(c.Copy())
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
