package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "valeraninja/noteapp/docs"
	"valeraninja/noteapp/internal/models"
)

type NoteItem interface {
	CreateItem(note models.Note) (int, error)
	GetAllItems() ([]models.Note, error)
	GetItemById(id int) (models.Note, error)
	DeleteItem(id int) error
	UpdateItem(id int, note models.Note) error
}

type Handler struct {
	noteService NoteItem
}

func NewHandler(noteService NoteItem) *Handler {
	return &Handler{noteService: noteService}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
