package controllers

import (
	// gin
	"net/http"

	"github.com/gin-gonic/gin"
	// models
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
	// services
	"github.com/rez-dev/backup_pingeso_2023/backend/services"
)

func GetCategories(c *gin.Context) {
	categories, err := services.GetCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, categories)
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := services.GetCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, category)
}

func CreateCategory(c *gin.Context) {
	var newCategory models.Category
	if err := c.BindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateCategory(newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newCategory)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var newCategory models.Category
	if err := c.BindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateCategory(id, newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, newCategory)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteCategory(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func GetParentsWordPress(c *gin.Context, id int) {
	parents, err := services.GetParentsWordPress(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, parents)
}

func InsertCategories(c *gin.Context) {
	categories, err := services.InsertCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, categories)
}

func GetCategoriesByTerm(c *gin.Context) {
	id := c.Param("id")
	categories, err := services.GetCategoriesByTerm(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, categories)
}
