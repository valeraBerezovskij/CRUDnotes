package handlers

import (
	"net/http"
	"valeraninja/noteapp/internal/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
	var note models.Note

	if err := c.BindJSON(&note); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.noteService.CreateItem(note)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	notes, err := h.noteService.GetAllItems()
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, notes)
}

func (h *Handler) getItemById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	note, err := h.noteService.GetItemById(idInt)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, note)
}

func (h *Handler) updateItem(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var note models.Note

	if err := c.BindJSON(&note); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.noteService.UpdateItem(idInt, note)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Note updated successfully",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	err = h.noteService.DeleteItem(idInt)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted successfully",
	})
}