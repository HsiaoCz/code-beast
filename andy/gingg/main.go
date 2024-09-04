package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	userHandler := &UserHandlers{}
	app.GET("/user", userHandler.HandleCreateUser)
	app.Run(":3001")
}
