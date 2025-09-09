package handler

import (
	"strconv"

	"github.com/bohdan-kozlo/todo-app/pkg/apperror"
	"github.com/gin-gonic/gin"
)

func getUserIdFromContext(c *gin.Context) (int, *apperror.AppError) {
	userId, err := getUserId(c)
	if err != nil {
		return 0, apperror.Internal("failed to get user id", err)
	}
	return userId, nil
}

func getListIdFromParam(c *gin.Context) (int, *apperror.AppError) {
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, apperror.BadRequest("invalid list id param", err)
	}
	return listId, nil
}

func getItemIdFromParam(c *gin.Context) (int, *apperror.AppError) {
	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		return 0, apperror.BadRequest("invalid item id param", err)
	}
	return itemId, nil
}
