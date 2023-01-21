package controller

import (
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
	"github.com/RizkiMufrizal/gin-clean-architecture/service"
	"github.com/gin-gonic/gin"
)

type HttpBinController struct {
	service.HttpBinService
}

func NewHttpBinController(httpBinService *service.HttpBinService) *HttpBinController {
	return &HttpBinController{HttpBinService: *httpBinService}
}

func (controller HttpBinController) Route(app *gin.Engine) {
	app.GET("/v1/api/httpbin", controller.PostHttpBin)
}

func (controller HttpBinController) PostHttpBin(c *gin.Context) {
	controller.HttpBinService.PostMethod(c.Copy())
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    nil,
	})
}
