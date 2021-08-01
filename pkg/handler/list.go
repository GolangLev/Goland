package handler

import (
	todo "github.com/GolangLev/Goland"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateList(c *gin.Context) {
	userId, err := GetuserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		NewMessageErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//Вызов метода сервиса
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		NewMessageErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type GetAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) GetAllLists(c *gin.Context) {
	userId, err := GetuserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		NewMessageErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) GetListsById(c *gin.Context) {
	userId, err := GetuserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewMessageErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		NewMessageErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateList(c *gin.Context) {
	userId, err := GetuserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewMessageErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		NewMessageErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, id, input); err != nil {
		NewMessageErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "Record update",
	})
}

func (h *Handler) DeleteList(c *gin.Context) {
	userId, err := GetuserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewMessageErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)
	if err != nil {
		NewMessageErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "Good",
	})
}
