package handler

import (
	"github.com/gin-gonic/gin"
	restgo "github.com/go-rest-api"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input restgo.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
