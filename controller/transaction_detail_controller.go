package controller

import (
	"github.com/RizkiMufrizal/gin-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gin-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
	"github.com/RizkiMufrizal/gin-clean-architecture/service"
	"github.com/gin-gonic/gin"
)

type TransactionDetailController struct {
	service.TransactionDetailService
	configuration.Config
}

func NewTransactionDetailController(transactionDetailService *service.TransactionDetailService, config configuration.Config) *TransactionDetailController {
	return &TransactionDetailController{TransactionDetailService: *transactionDetailService, Config: config}
}

func (controller TransactionDetailController) Route(app *gin.Engine) {
	app.GET("/v1/api/transaction-detail/:id", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.FindById)
}

// FindById func gets one exists transaction detail.
// @Description Get one exists transaction detail.
// @Summary get one exists transaction detail
// @Tags Transaction Detail
// @Accept json
// @Produce json
// @Param id path string true "Transaction Detail Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/transaction-detail/{id} [get]
func (controller TransactionDetailController) FindById(c *gin.Context) {
	id := c.Param("id")

	result := controller.TransactionDetailService.FindById(c.Copy(), id)
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
