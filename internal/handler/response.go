package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorRes struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorRes{Message: message})
}

type dataRes[T any] struct {
	Data T `json:"data"`
}

func newDataResponse[T any](c *gin.Context, statusCode int, data T) {
	c.JSON(statusCode, dataRes[T]{Data: data})
}
