package handler

import (
	"github.com/bohdan-kozlo/todo-app/internal/service"
	"github.com/bohdan-kozlo/todo-app/pkg/apperror"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.wrap(h.signUp))
		auth.POST("sign-in", h.wrap(h.signIn))
	}

	api := router.Group("/api", h.userIdentify)
	{
		lists := api.Group("/lists")
		{
			lists.POST("", h.wrap(h.createList))
			lists.GET("", h.wrap(h.getAllLists))
			lists.GET("/:id", h.wrap(h.getListById))
			lists.PUT("/:id", h.wrap(h.updateList))
			lists.DELETE("/:id", h.wrap(h.deleteList))

			items := lists.Group(":id/items")
			{
				items.POST("", h.wrap(h.createItem))
				items.GET("", h.wrap(h.getAllItems))
				items.GET("/:item_id", h.wrap(h.getItemById))
				items.PUT("/:item_id", h.wrap(h.updateItem))
				items.DELETE("/:item_id", h.wrap(h.deleteItem))
			}
		}
	}

	return router
}

type appHandler func(*gin.Context) *apperror.AppError

func (h *Handler) wrap(fn appHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(c); err != nil {
			writeError(c, err)
		}
	}
}
