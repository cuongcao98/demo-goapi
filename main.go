package main

import (
	"github.com/gin-gonic/gin"
	"project/controller"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/public", "./public")

	client := r.Group("/api")
	{
		client.GET("/user/:id", controller.Read)
		client.GET("/alluser", controller.ReadAll)
		client.POST("/user/create", controller.Create)
		client.PATCH("/user/update/:id", controller.Update)
		client.DELETE("/user/:id", controller.Delete)
		client.GET("/national/:id", controller.Read_nationalid)
		client.GET("/birth/:id", controller.Read_birth)
	}
	
	return r
}


func main() {
	r := setupRouter()
    r.Run(":8080")
}