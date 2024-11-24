package controllers

import (
	"net/http"

	"api-restfull-crud-in-memory/models"

	"api-restfull-crud-in-memory/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func GetItems(repo *repository.ItemRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		items := repo.GetAll()
		c.JSON(http.StatusOK, items)
	}
}

func CreateItem(repo *repository.ItemRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if err := validate.Struct(item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		item = repo.Create(item)
		c.JSON(http.StatusCreated, item)
	}
}

func GetItemByID(repo *repository.ItemRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		item, found := repo.GetByID(id)
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
			return
		}
		c.JSON(http.StatusOK, item)
	}
}

func UpdateItem(repo *repository.ItemRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var item models.Item
		if err := c.ShouldBindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if updatedItem, found := repo.Update(id, item); found {
			c.JSON(http.StatusOK, updatedItem)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		}
	}
}

func DeleteItem(repo *repository.ItemRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if success := repo.Delete(id); success {
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		}
	}
}
