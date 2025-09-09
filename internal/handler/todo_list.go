package handler

import (
	"net/http"
	"strconv"

	"github.com/bohdan-kozlo/todo-app/internal/models"
	"github.com/bohdan-kozlo/todo-app/pkg/apperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) *apperror.AppError {
	userId, err := getUserId(c)
	if err != nil {
		return apperror.Internal("failed to get user id", err)
	}
	var input models.TodoList
	if err := c.BindJSON(&input); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}
	listId, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		return apperror.Internal("failed to create list", err)
	}
	c.JSON(http.StatusOK, map[string]any{"id": listId})
	return nil
}

func (h *Handler) getAllLists(c *gin.Context) *apperror.AppError {
	userId, err := getUserId(c)
	if err != nil {
		return apperror.Internal("failed to get user id", err)
	}
	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		return apperror.Internal("failed to fetch lists", err)
	}
	newDataResponse(c, http.StatusOK, lists)
	return nil
}

func (h *Handler) getListById(c *gin.Context) *apperror.AppError {
	userId, err := getUserId(c)
	if err != nil {
		return apperror.Internal("failed to get user id", err)
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apperror.BadRequest("invalid id param", err)
	}
	list, err := h.services.TodoList.GetById(userId, listId)
	if err != nil {
		if isNotFound(err) {
			return apperror.NotFound("list not found", err)
		}
		return apperror.Internal("failed to fetch list", err)
	}
	newDataResponse(c, http.StatusOK, list)
	return nil
}

func (h *Handler) updateList(c *gin.Context) *apperror.AppError {
	userId, err := getUserId(c)
	if err != nil {
		return apperror.Internal("failed to get user id", err)
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apperror.BadRequest("invalid id param", err)
	}
	var input models.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}
	if input.Title == nil && input.Description == nil {
		return apperror.BadRequest("no values to update", nil)
	}
	if err := h.services.TodoList.Update(userId, listId, input); err != nil {
		if isNotFound(err) {
			return apperror.NotFound("list not found", err)
		}
		return apperror.Internal("failed to update list", err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	return nil
}

func (h *Handler) deleteList(c *gin.Context) *apperror.AppError {
	userId, err := getUserId(c)
	if err != nil {
		return apperror.Internal("failed to get user id", err)
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apperror.BadRequest("invalid id param", err)
	}
	if err := h.services.TodoList.Delete(userId, listId); err != nil {
		if isNotFound(err) {
			return apperror.NotFound("list not found", err)
		}
		return apperror.Internal("failed to delete list", err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	return nil
}
