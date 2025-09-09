package handler

import (
	"net/http"
	"strconv"

	"github.com/bohdan-kozlo/todo-app/internal/models"
	"github.com/bohdan-kozlo/todo-app/pkg/apperror"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) *apperror.AppError {
	userId, appErr := getUserIdFromContext(c)
	if appErr != nil {
		return appErr
	}

	listId, appErr := getListIdFromParam(c)
	if appErr != nil {
		return appErr
	}

	var input models.TodoItem
	if err := c.BindJSON(&input); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	itemId, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		if isNotFound(err) {
			return apperror.NotFound("list not found", err)
		}
		return apperror.Internal("failed to create item", err)
	}

	c.JSON(http.StatusOK, map[string]any{"id": itemId})
	return nil
}

func (h *Handler) getAllItems(c *gin.Context) *apperror.AppError {
	userId, err := getUserId(c)
	if err != nil {
		return apperror.Internal("failed to get user id", err)
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apperror.BadRequest("invalid list id param", err)
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		if isNotFound(err) {
			return apperror.NotFound("list not found", err)
		}
		return apperror.Internal("failed to fetch items", err)
	}
	newDataResponse(c, http.StatusOK, items)
	return nil
}

func (h *Handler) getItemById(c *gin.Context) *apperror.AppError {
	userId, appErr := getUserIdFromContext(c)
	if appErr != nil {
		return appErr
	}

	listId, appErr := getListIdFromParam(c)
	if appErr != nil {
		return appErr
	}

	itemId, appErr := getItemIdFromParam(c)
	if appErr != nil {
		return appErr
	}

	item, err := h.services.TodoItem.GetById(userId, listId, itemId)
	if err != nil {
		if isNotFound(err) {
			return apperror.NotFound("item not found", err)
		}
		return apperror.Internal("failed to fetch item", err)
	}
	newDataResponse(c, http.StatusOK, item)
	return nil
}

func (h *Handler) updateItem(c *gin.Context) *apperror.AppError {
	userId, appErr := getUserIdFromContext(c)
	if appErr != nil {
		return appErr
	}

	listId, appErr := getListIdFromParam(c)
	if appErr != nil {
		return appErr
	}

	itemId, appErr := getItemIdFromParam(c)
	if appErr != nil {
		return appErr
	}

	var input models.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}
	if input.Title == nil && input.Description == nil && input.Completed == nil {
		return apperror.BadRequest("no values to update", nil)
	}
	if err := h.services.TodoItem.Update(userId, listId, itemId, input); err != nil {
		if isNotFound(err) {
			return apperror.NotFound("item or list not found", err)
		}
		return apperror.Internal("failed to update item", err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	return nil
}

func (h *Handler) deleteItem(c *gin.Context) *apperror.AppError {
	userId, err := getUserId(c)
	if err != nil {
		return apperror.Internal("failed to get user id", err)
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return apperror.BadRequest("invalid list id param", err)
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		return apperror.BadRequest("invalid item id param", err)
	}
	if err := h.services.TodoItem.Delete(userId, listId, itemId); err != nil {
		if isNotFound(err) {
			return apperror.NotFound("item or list not found", err)
		}
		return apperror.Internal("failed to delete item", err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	return nil
}
