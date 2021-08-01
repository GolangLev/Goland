package handler

import (
	todo "github.com/GolangLev/Goland"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		NewMessageErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewMessageErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		NewMessageErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewMessageErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
