package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userctx             = "userid"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewMessageErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewMessageErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	//Парсим токен
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewMessageErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userctx, userId)
}

func GetuserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userctx)
	if !ok {
		NewMessageErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewMessageErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}
