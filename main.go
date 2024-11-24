package main

import (
	"api-restfull-crud-in-memory/controllers"
	"api-restfull-crud-in-memory/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewItemRepository()

	r := gin.Default()

	r.GET("/items", controllers.GetItems(repo))
	r.GET("/items/:id", controllers.GetItemByID(repo))
	r.POST("/items", controllers.CreateItem(repo))
	r.PUT("/items/:id", controllers.UpdateItem(repo))
	r.DELETE("/items/:id", controllers.DeleteItem(repo))

	r.Run(":8080")
}
