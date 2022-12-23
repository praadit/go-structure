package exception

import (
	"encoding/json"
	"fmt"
	"go-best-practice/internal/constants"
	"go-best-practice/internal/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	if err != nil {
		if maintenanceErr, ok := err.(*MaintenanceError); ok {
			c.Logger().Error(maintenanceErr.Message)
			maintenanceErrorHandler(maintenanceErr, c)
			return
		}
		if unauthorized, ok := err.(*UnauthorizedError); ok {
			c.Logger().Error(unauthorized.Message)
			unauthorizedErrorHandler(unauthorized, c)
			return
		}
		if forbidden, ok := err.(*ForbiddendError); ok {
			c.Logger().Error(forbidden.Message)
			forbiddenErrorHandler(forbidden, c)
			return
		}
		if notfound, ok := err.(*NotFoundError); ok {
			c.Logger().Error(notfound.Message)
			notFoundErrorHandler(notfound, c)
			return
		}
		if validationError, ok := err.(*ValidationError); ok {
			c.Logger().Error(validationError.Message)
			validationErrorHandler(validationError, c)
			return
		}
		if badRequest, ok := err.(*BadRequestError); ok {
			c.Logger().Error(badRequest.Message)
			badRequestErrorHandler(badRequest, c)
			return
		}

		endpoint := c.Request().URL.String()
		var json map[string]interface{} = map[string]interface{}{}
		c.Bind(&json)

		c.Logger().Error(fmt.Sprintf("internal error : %s, with request %v, return error %s", endpoint, json, err.Error()))
		internalServerErrorHandler(err, c)
	}
}

func internalServerErrorHandler(err error, c echo.Context) {
	c.JSON(http.StatusInternalServerError,
		response.BaseResponse{
			ErrorCode:       constants.ERR_SYSTEM_UnknownError,
			Message:         "Internal Error!",
			InternalMessage: err.Error(),
		},
	)
}

func notFoundErrorHandler(err *NotFoundError, c echo.Context) {
	c.JSON(http.StatusNotFound,
		response.BaseResponse{
			ErrorCode: err.Code,
			Message:   err.Message,
		},
	)
}
func maintenanceErrorHandler(err *MaintenanceError, c echo.Context) {
	c.JSON(http.StatusServiceUnavailable,
		response.BaseResponse{
			ErrorCode: err.Code,
			Message:   err.Message,
		},
	)
}

func unauthorizedErrorHandler(err *UnauthorizedError, c echo.Context) {
	c.JSON(http.StatusUnauthorized,
		response.BaseResponse{
			ErrorCode: err.Code,
			Message:   err.Message,
		},
	)
}

func forbiddenErrorHandler(err *ForbiddendError, c echo.Context) {
	c.JSON(http.StatusForbidden,
		response.BaseResponse{
			ErrorCode: err.Code,
			Message:   err.Message,
		},
	)
}

func badRequestErrorHandler(err *BadRequestError, c echo.Context) {
	c.JSON(http.StatusBadRequest,
		response.BaseResponse{
			ErrorCode: err.Code,
			Message:   err.Message,
		},
	)
}

func validationErrorHandler(err *ValidationError, c echo.Context) {
	var stringArray []string
	hasError := json.Unmarshal([]byte(err.Message), &stringArray)
	if hasError != nil {
		panic(hasError)
	}

	c.JSON(http.StatusBadRequest,
		response.BaseResponse{
			ErrorCode: err.Code,
			Message:   stringArray,
		},
	)
}
