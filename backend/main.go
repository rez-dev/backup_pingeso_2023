package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rez-dev/backup_pingeso_2023/backend/controllers"
)

func main() {
	router := gin.Default()

	// MANEJO DE CORS
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:3000/", "http://localhost:8080", "http://localhost:8080/"}
	config.AllowOrigins = []string{"http://localhost:3000"}
	//config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	// router.POST("/signup", controllers.Signup)
	// router.POST("/login", controllers.Login)
	// router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// el usuario logeado pueda
	//cors
	//router.Use(middleware.CorsMiddleware())

	router.GET("/services", controllers.GetServices)
	router.GET("/services/:id", controllers.GetService)
	router.POST("/service", controllers.CreateService)
	router.PUT("/services/:id", controllers.UpdateService)
	router.DELETE("/services/:id", controllers.DeleteService)
	// Ruta a la que los Revisores, Admin y SuperAdmin pueden acceder
	router.GET("/servicesByUser/:id", controllers.GetServicesByUser)
	// Ruta específica con middlewares en cadena
	router.PUT("/assignServices", controllers.AssignServices)
	router.GET("/countServicesByUser/:id", controllers.CountServicesByUser)
	router.GET("/servicesByState/:state", controllers.GetServicesByState) //En la ruta se agrega el estado del servicio
	router.PUT("/desassignService/:id", controllers.DesassignService)     // En el body se debe enviar el id del servicio
	// router.GET("/getUserByService/:id", controllers.GetUserByService)     // En la ruta se agrega el id del servicio
	router.PUT("/approveService/:id", controllers.ApproveService) // En la ruta se agrega el id_wp_term del servicio

	// router.GET("/categories", controllers.GetCategories)
	// router.GET("/categories/:id", controllers.GetCategory)
	// router.POST("/category", controllers.CreateCategory)
	// router.PUT("/categories/:id", controllers.UpdateCategory)
	// router.DELETE("/categories/:id", controllers.DeleteCategory)
	// router.GET("/categoriesByTerm/:id", controllers.GetCategoriesByTerm)

	router.GET("/informations", controllers.GetInformations)
	router.GET("/informations/:id", controllers.GetInformation)
	router.POST("/information", controllers.CreateInformation)
	router.PUT("/informations/:id", controllers.UpdateInformation)
	router.DELETE("/informations/:id", controllers.DeleteInformation)
	// router.GET("/informationwp", controllers.GetInformationWordPress)
	router.GET("/informationByTerm/:id", controllers.GetInformationsByTerm)

	// router.GET("/users", controllers.GetUsers)
	// router.GET("/users/:id", controllers.GetUser)
	// router.POST("/user", controllers.CreateUser)
	// router.PUT("/users/:id", controllers.UpdateUser)
	// router.DELETE("/users/:id", controllers.DeleteUser)
	// router.GET("/usersByUnity/:unity", controllers.GetUsersByUnity) // parametro role queda opcional y llamas a la ruta como localhost:8080/usersByUnity/VRAE?role=revisor sino localhost:8080/usersByUnity/VRAE
	// router.GET("/usersByRole/:role", controllers.GetUsersByRole)
	// router.GET("/parent/:id", controllers.GetParentWordPress)

	// INTERACCIONES CON LA BASE DE DATOS DE WORDPRESS
	router.GET("/insertServices", controllers.InsertServices)
	router.GET("/insertInformations", controllers.InsertInformations)
	// router.GET("/insertCategories", controllers.InsertCategories)
	router.PUT("/approveServiceWP/:id", controllers.ApproveServiceWP)         //Se debe ingresar el id del wp_term (id en bd wordpress) y tambien el json con los datos del servicio de nuestra bd
	router.PUT("/approveInformationWP/:id", controllers.ApproveInformationWP) //Se debe ingresar el id del wp_post (id en bd wordpress) y tambien el json con los datos de la informacion de nuestra bd

	//////////////////////////////////////////////////////////////////////////////////////
	//Rutas para Revisores, Administradores y SuperAdministradores (debe estar autenticado)
	//////////////////////////////////////////////////////////////////////////////////////
	// router.GET("/services/:id", middleware.RequireAuth, controllers.GetService)
	// router.GET("/servicesByUser/:id", middleware.RequireAuth, controllers.GetServicesByUser)
	// router.GET("/servicesByState/:state", middleware.RequireAuth, controllers.GetServicesByState)
	// router.PUT("/approveService/:id", middleware.RequireAuth, controllers.ApproveService)
	// router.GET("/getUserByService/:id", middleware.RequireAuth, controllers.GetUserByService)
	// router.GET("/categories", middleware.RequireAuth, controllers.GetCategories)
	// router.GET("/informationByTerm/:id", middleware.RequireAuth, controllers.GetInformationsByTerm)
	// router.GET("/informations", middleware.RequireAuth, controllers.GetInformations)
	// router.PUT("/approveServiceWP/:id", middleware.RequireAuth, controllers.ApproveServiceWP)
	// router.PUT("/approveInformationWP/:id", middleware.RequireAuth, controllers.ApproveInformationWP)

	//////////////////////////////////////////////////////////////////////////////////////
	// Rutas Exclusivas para Administradores y SuperAdministradores
	//////////////////////////////////////////////////////////////////////////////////////

	// router.GET("/users", middleware.RequireAuth, middleware.RoleBasedAuth(models.Admin, models.SuperAdmin), controllers.GetUsers)
	// router.GET("/users/:id", middleware.RequireAuth, middleware.RoleBasedAuth(models.Admin, models.SuperAdmin), controllers.GetUser)
	// router.POST("/user", middleware.RequireAuth, middleware.RoleBasedAuth(models.Admin, models.SuperAdmin), controllers.CreateUser)
	// router.GET("/usersByUnity/:unity", middleware.RequireAuth, middleware.RoleBasedAuth(models.Admin, models.SuperAdmin), controllers.GetUsersByUnity)
	// router.GET("/usersByRole/:role", middleware.RequireAuth, middleware.RoleBasedAuth(models.Admin, models.SuperAdmin), controllers.GetUsersByRole)
	// router.PUT("/assignServices", middleware.RequireAuth, middleware.RoleBasedAuth(models.Admin, models.SuperAdmin), controllers.AssignServices)
	// router.PUT("/desassignService/:id", middleware.RequireAuth, middleware.RoleBasedAuth(models.Admin, models.SuperAdmin), controllers.DesassignService)

	router.Run(":8080")
}
