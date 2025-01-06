package handlers

import (
	"net/http"
	"valeraninja/noteapp/internal/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary Create a note
// @Description Create a note by note struct
// @Tags Notes
// @Accept       json
// @Produce      json
// @Param        note  body  models.Note  true  "Note struct"
// @Success      200   {object}  ErrorResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Router       /api/items/    [POST]
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

// @Summary Get all notes
// @Description Get all notes
// @Tags Notes
// @Accept       json
// @Produce      json
// @Success      200   {object}  ErrorResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Router       /api/items/    [GET]
func (h *Handler) getAllItems(c *gin.Context) {
	notes, err := h.noteService.GetAllItems()
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, notes)
}

// @Summary Get note by ID
// @Description Get note by ID
// @Tags Notes
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Note ID"
// @Success      200   {object}  ErrorResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Router       /api/items/{id}   [GET]
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

// @Summary      Update a note
// @Description  Updates a note by ID. Can changes only Title and Description.
// @Tags         Notes
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Note ID"
// @Param        note  body      models.Note true  "Note struct"
// @Success      200   {object}  ErrorResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Router       /api/items/{id} [PUT]
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

// @Summary Delete note by ID
// @Description Delete note by ID
// @Tags Notes
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Note ID"
// @Success      200   {object}  ErrorResponse
// @Failure      400   {object}  ErrorResponse
// @Failure      404   {object}  ErrorResponse
// @Router       /api/items/{id}   [DELETE]
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