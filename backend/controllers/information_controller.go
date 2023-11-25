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

func GetInformations(c *gin.Context) {
	informations, err := services.GetInformations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, informations)
}

func GetInformation(c *gin.Context) {
	id := c.Param("id")
	information, err := services.GetInformation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, information)
}

func CreateInformation(c *gin.Context) {
	var newInformation models.Information
	if err := c.BindJSON(&newInformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateInformation(newInformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newInformation)
}

func UpdateInformation(c *gin.Context) {
	id := c.Param("id")
	var newInformation models.Information
	if err := c.BindJSON(&newInformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateInformation(id, newInformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, newInformation)
}

func DeleteInformation(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteInformation(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Information deleted"})
}

func InsertInformations(c *gin.Context) {
	informations, err := services.InsertInformations()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, informations)
}

func GetInformationsByTerm(c *gin.Context) {
	id := c.Param("id")
	informations, err := services.GetInformationsByTerm(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, informations)
}

func ApproveInformationWP(c *gin.Context) {
	id := c.Param("id")
	var newInformation models.Information
	if err := c.BindJSON(&newInformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.ApproveInformationWP(id, newInformation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, newInformation)
}
