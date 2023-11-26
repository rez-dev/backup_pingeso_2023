package controllers

import (
	// gin
	"net/http"

	"github.com/gin-gonic/gin"
	// models

	// services
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
	"github.com/rez-dev/backup_pingeso_2023/backend/services"
)

func GetUsers(c *gin.Context) {
	services, err := services.GetServices()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, services)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	service, err := services.GetService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, service)
}

func CreateUser(c *gin.Context) {
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

func UpdateUser(c *gin.Context) {
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

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteService(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Service deleted"})
}

func GetUsersByUnity(c *gin.Context) {
	unity := c.Param("unity")
	role := c.Param("role")
	users, err := services.GetUsersByUnity(unity, role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func GetUserByRole(c *gin.Context) {
	role := c.Param("role")
	users, err := services.GetUserByRole(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func Signup(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := services.Signup(newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, newUser)
}

func Login(c *gin.Context) {
	var loginReq struct {
		Email    string
		Password string
	}
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := services.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	//user.(models.User).
	c.IndentedJSON(http.StatusOK, gin.H{"message": user})
}
