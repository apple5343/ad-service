package middleware

import (
	"fmt"
	"net/http"
	"server/pkg/errors"
	"server/pkg/errors/validate"
	"server/pkg/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func ErrorHandlerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if nil == err {
			return nil
		}
		if c.Response().Committed {
			return err
		}
		if httpErr, ok := err.(*echo.HTTPError); ok {
			return c.JSON(httpErr.Code, map[string]string{
				"status":  "error",
				"message": fmt.Sprintf("%v", httpErr.Message),
			})
		}
		switch {
		case validate.IsValidationError(err):
			logger.Error(err.Error(), zap.String("error_type", "validation_error"))
			return c.JSON(http.StatusBadRequest, map[string]string{"status": "error", "message": err.Error()})
		case errors.IsCustomError(err):
			customErr := errors.GetCommonError(err)
			if customErr.Code() == errors.Internal {
				logger.Error(err.Error(), zap.String("error_type", "internal_error"))
				return c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "Internal server error"})
			}
			logger.Error(customErr.Error(), zap.String("error_type", "custom_error"))
			return c.JSON(parseCommonCode(customErr.Code()), map[string]string{"status": "error", "message": customErr.Error()})
		default:
			logger.Error(err.Error(), zap.String("error_type", "unknown_error"))
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": "error", "message": "Internal server error"})
		}
	}
}

func parseCommonCode(code errors.Code) int {
	switch code {
	case 0:
		return http.StatusOK
	case 1:
		return http.StatusBadRequest
	case 2:
		return http.StatusNotFound
	case 3:
		return http.StatusInternalServerError
	case 4:
		return http.StatusConflict
	case 5:
		return http.StatusUnauthorized
	case 6:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}
