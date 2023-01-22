package exception

import (
	"encoding/json"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	errorResult, validationError := err.(ValidationError)
	if validationError {
		data := errorResult.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)
		c.AbortWithStatusJSON(400, model.GeneralResponse{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
		return
	}

	notFoundErrorResult, notFoundError := err.(NotFoundError)
	if notFoundError {
		c.AbortWithStatusJSON(404, model.GeneralResponse{
			Code:    404,
			Message: "Not Found",
			Data:    notFoundErrorResult.Message,
		})
		return
	}

	UnauthorizedErrorResult, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		c.AbortWithStatusJSON(401, model.GeneralResponse{
			Code:    401,
			Message: "Unauthorized",
			Data:    UnauthorizedErrorResult.Message,
		})
		return
	}

	c.AbortWithStatusJSON(500, model.GeneralResponse{
		Code:    500,
		Message: "General Error",
		Data:    err,
	})
	return
}
