package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app/pkg/todoItems"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todoItems.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllList(c *gin.Context) {

}

func (h *Handler) getByIdList(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
