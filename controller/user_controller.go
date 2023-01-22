package controller

import (
	"github.com/RizkiMufrizal/gin-clean-architecture/common"
	"github.com/RizkiMufrizal/gin-clean-architecture/configuration"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
	"github.com/RizkiMufrizal/gin-clean-architecture/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

type UserController struct {
	service.UserService
	configuration.Config
}

func (controller UserController) Route(app *gin.Engine) {
	app.POST("/v1/api/authentication", controller.Authentication)
}

// Authentication func Authenticate user.
// @Description authenticate user.
// @Summary authenticate user
// @Tags Authenticate user
// @Accept json
// @Produce json
// @Param request body model.UserModel true "Request Body"
// @Success 200 {object} model.GeneralResponse
// @Router /v1/api/authentication [post]
func (controller UserController) Authentication(c *gin.Context) {
	var request model.UserModel
	err := c.BindJSON(&request)
	exception.PanicLogging(err)

	result := controller.UserService.Authentication(c.Copy(), request)
	var userRoles []map[string]interface{}
	for _, userRole := range result.UserRoles {
		userRoles = append(userRoles, map[string]interface{}{
			"role": userRole.Role,
		})
	}

	jwtSecret := controller.Config.Get("JWT_SECRET_KEY")
	jwtExpired, err := strconv.Atoi(controller.Config.Get("JWT_EXPIRE_MINUTES_COUNT"))
	exception.PanicLogging(err)

	tokenJwtResult := common.GenerateToken(result.Username, userRoles, jwtSecret, jwtExpired)
	resultWithToken := map[string]interface{}{
		"token":    tokenJwtResult,
		"username": result.Username,
		"role":     userRoles,
	}
	c.JSON(200, model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    resultWithToken,
	})
}
