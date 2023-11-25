package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rez-dev/backup_pingeso_2023/backend/database"
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
)

func GetServices() (services []models.Service, err error) {
	conexionEstablecida := database.ConexionDB()
	if err != nil {
		return nil, err
	}
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, _ := conexionEstablecida.Query("SELECT * FROM service")
	service := models.Service{}
	services := models.Services{}
	for obtenerRegistros.Next() {
		var id int
		var id_user, id_wp_term int
		var name, description, state string
		obtenerRegistros.Scan(&id, &name, &description, &state, &id_user, &id_wp_term)
		service.ID = id
		service.Name = name
		service.Description = description
		service.State = state
		service.Id_user = id_user
		service.Id_wp_term = id_wp_term
		services = append(services, service)
	}
}
