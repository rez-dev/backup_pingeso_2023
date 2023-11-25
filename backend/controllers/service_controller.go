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
		c.JSON(500, gin.H{
			"message": "Error al obtener los servicios",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func GetService(c *gin.Context) {
	id := c.Param("id")
	services, err := services.GetService(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error al obtener el servicio",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func CreateService(c *gin.Context) {
	var newService models.Service
	c.BindJSON(&newService)
	err := services.CreateService(newService)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error al crear el servicio",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, newService)
}

func UpdateService(c *gin.Context) {
	id := c.Param("id")
	var newService models.Service
	c.BindJSON(&newService)
	err := services.UpdateService(id, newService)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error al actualizar el servicio",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, newService)
}

func DeleteService(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteService(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error al eliminar el servicio",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Servicio eliminado",
	})
}

func InsertServices(c *gin.Context) {
	services, err := services.InsertServices()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error al insertar los servicios",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

// func GetServicesWordPress(c *gin.Context) {
// 	services, err := services.GetServicesWordPress()
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"message": "Error al obtener los servicios",
// 		})
// 		return
// 	}
// 	c.IndentedJSON(http.StatusOK, services)
// }

func ApproveServiceWP(c *gin.Context) {
	id := c.Param("id")
	var newService models.Service
	c.BindJSON(&newService)
	err := services.ApproveServiceWP(id, newService)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error al aprobar el servicio",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, newService)
}

func GetServicesByUser(c *gin.Context) {
	id := c.Param("id")
	services, err := services.GetServicesByUser(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error al obtener los servicios",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func AssignServices(c *gin.Context) {
	type requestService struct {
		IDUser    int   `json:"id_user"`
		Servicios []int `json:"servicios"`
	}
	var request requestService
	c.BindJSON(&request)

	err := services.AssignServices(request.IDUser, request.Servicios)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error al asignar los servicios",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Servicios asignados",
	})
}

func CountServicesByUser(c *gin.Context) {
	id := c.Param("id")
	count, err := services.CountServicesByUser(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error al contar los servicios",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"count": count})
}
