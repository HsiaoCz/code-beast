package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlers struct{}

func (u *UserHandlers) HandleCreateUser(c *gin.Context) {
	username := c.Query("username")
	email := c.Query("email")
	user := &User{
		Username: username,
		Email:    email,
	}
	c.JSON(http.StatusOK, user)
}
