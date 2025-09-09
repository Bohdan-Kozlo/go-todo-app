package handler

import (
	"net/http"
	"time"

	"github.com/bohdan-kozlo/todo-app/pkg/apperror"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type apiError struct {
	Error     string `json:"error"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type apiData[T any] struct {
	Data T `json:"data"`
}

func writeError(c *gin.Context, ae *apperror.AppError) {
	if ae == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, apiError{Error: "internal_error", Message: "internal server error", Timestamp: time.Now().UTC().Format(time.RFC3339)})
		return
	}
	entry := logrus.WithField("code", ae.Code)
	if ae.Err != nil {
		entry = entry.WithError(ae.Err)
	}
	if ae.HTTPStatus >= 500 {
		entry.Error(ae.Message)
	} else {
		entry.Warn(ae.Message)
	}
	c.AbortWithStatusJSON(ae.HTTPStatus, apiError{Error: ae.Code, Message: ae.Message, Timestamp: time.Now().UTC().Format(time.RFC3339)})
}

func newDataResponse[T any](c *gin.Context, statusCode int, data T) {
	c.JSON(statusCode, apiData[T]{Data: data})
}
