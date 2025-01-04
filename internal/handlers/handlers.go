package handlers

import (
	"valeraninja/noteapp/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.GET("/", h.getAllItems)
			items.POST("/", h.createItem)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)

		}
	}
	return router
}
