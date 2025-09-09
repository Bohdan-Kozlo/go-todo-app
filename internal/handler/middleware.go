package handler

import (
	"errors"
	"strings"

	"github.com/bohdan-kozlo/todo-app/pkg/apperror"
	"github.com/gin-gonic/gin"
)

const (
	userCtx = "userId"
)

func (h *Handler) userIdentify(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		writeError(c, apperror.Unauthorized("no Authorization header", nil))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		writeError(c, apperror.Unauthorized("invalid Authorization header", nil))
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		writeError(c, apperror.Unauthorized("invalid token", err))
		return
	}

	c.Set(userCtx, userId)
	c.Next()
}

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}
	userIdInt, ok := userId.(int)
	if !ok {
		return 0, errors.New("user id has invalid type")
	}
	return userIdInt, nil
}
