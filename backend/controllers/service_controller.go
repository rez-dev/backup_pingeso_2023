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

func GetServices(c *gin.Context) {
	services, err := services.GetServices()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func GetService(c *gin.Context) {
	id := c.Params.ByName("id")
	service, err := services.GetService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, service)
}

func CreateService(c *gin.Context) {
	var newService models.Service
	if err := c.BindJSON(&newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.CreateService(newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newService)
}

func UpdateService(c *gin.Context) {
	id := c.Param("id")
	var newService models.Service
	if err := c.BindJSON(&newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.UpdateService(id, newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, newService)
}

func DeleteService(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteService(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Service deleted"})
}

func GetServicesWordPress(c *gin.Context) {
	services, err := services.GetServicesWordPress()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func InsertServices(c *gin.Context) {
	services, err := services.InsertServices()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func ApproveServiceWP(c *gin.Context) {
	id := c.Param("id")
	var newService models.Service
	if err := c.BindJSON(&newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.ApproveServiceWP(id, newService); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, newService)
}

func ApproveService(c *gin.Context) {
	id := c.Param("id")
	type requestService struct {
		Id_user     int    `json:"id_user"`
		Description string `json:"description"`
	}
	var request requestService
	c.BindJSON(&request)

	if err := services.ApproveService(id, request.Id_user, request.Description); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Service approved"})
}

func GetServicesByUser(c *gin.Context) {
	id := c.Param("id")
	services, err := services.GetServicesByUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func AssignServices(c *gin.Context) {
	type requestAssign struct {
		Id_user   int   `json:"id_user"`
		Servicios []int `json:"servicios"`
	}
	var request requestAssign
	c.BindJSON(&request)
	if err := services.AssignServices(request.Id_user, request.Servicios); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Services assigned"})
}

func DesassignService(c *gin.Context) {
	id := c.Param("id")
	if err := services.DesassignService(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Service desassigned"})
}

func CountServicesByUser(c *gin.Context) {
	id := c.Param("id")
	count, err := services.CountServicesByUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, count)
}

func GetServicesByState(c *gin.Context) {
	state := c.Param("state")
	services, err := services.GetServicesByState(state)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func GetUserByService(c *gin.Context) {
	id := c.Param("id")
	name, err := services.GetUserByService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"name": name})
}
