package controller

import (
	"github.com/RizkiMufrizal/gin-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/RizkiMufrizal/gin-clean-architecture/middleware"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
	"github.com/RizkiMufrizal/gin-clean-architecture/service"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service.ProductService
	configuration.Config
}

func NewProductController(productService *service.ProductService, config configuration.Config) *ProductController {
	return &ProductController{ProductService: *productService, Config: config}
}

func (controller ProductController) Route(app *gin.Engine) {
	app.POST("/v1/api/product", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.Create)
	app.PUT("/v1/api/product/:id", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.Update)
	app.DELETE("/v1/api/product/:id", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.Delete)
	app.GET("/v1/api/product/:id", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.FindById)
	app.GET("/v1/api/product", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.FindAll)
}

// Create func create product.
// @Description create product.
// @Summary create product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product [post]
func (controller ProductController) Create(c *gin.Context) {
	var request model.ProductCreateOrUpdateModel
	err := c.BindJSON(&request)
	exception.PanicLogging(err)

	response := controller.ProductService.Create(c.Copy(), request)
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// Update func update one exists product.
// @Description update one exists product.
// @Summary update one exists product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body model.ProductCreateOrUpdateModel true "Request Body"
// @Param id path string true "Product Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product/{id} [put]
func (controller ProductController) Update(c *gin.Context) {
	var request model.ProductCreateOrUpdateModel
	id := c.Param("id")
	err := c.BindJSON(&request)
	exception.PanicLogging(err)

	response := controller.ProductService.Update(c.Copy(), request, id)
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// Delete func delete one exists product.
// @Description delete one exists product.
// @Summary delete one exists product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product/{id} [delete]
func (controller ProductController) Delete(c *gin.Context) {
	id := c.Param("id")

	controller.ProductService.Delete(c.Copy(), id)
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

// FindById func gets one exists product.
// @Description Get one exists product.
// @Summary get one exists product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product/{id} [get]
func (controller ProductController) FindById(c *gin.Context) {
	id := c.Param("id")

	result := controller.ProductService.FindById(c.Copy(), id)
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// FindAll func gets all exists products.
// @Description Get all exists products.
// @Summary get all exists products
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/product [get]
func (controller ProductController) FindAll(c *gin.Context) {
	result := controller.ProductService.FindAll(c.Copy())
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
