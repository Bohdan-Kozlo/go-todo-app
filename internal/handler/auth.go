package handler

import (
	"net/http"

	"github.com/bohdan-kozlo/todo-app/internal/models"
	"github.com/bohdan-kozlo/todo-app/pkg/apperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) *apperror.AppError {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	userId, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		return apperror.Internal("failed to create user", err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": userId})
	return nil
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) *apperror.AppError {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		return apperror.Unauthorized("invalid credentials", err)
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
	return nil
}
