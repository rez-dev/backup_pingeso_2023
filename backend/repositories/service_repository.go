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
	services = models.Services{}
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
	return services, nil
}

func GetService(id string) (services []models.Service, err error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, _ := conexionEstablecida.Query("SELECT * FROM service WHERE id = ?", id)
	service := models.Service{}
	services = models.Services{}
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
	return services, nil
}

func CreateService(newService models.Service) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO service(name, description, state, id_user, id_wp_term) VALUES(?,?,?,?,?)")
	insertarRegistros.Exec(newService.Name, newService.Description, newService.State, newService.Id_user, newService.Id_wp_term)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func UpdateService(id string, newService models.Service) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("UPDATE service SET name=?, description=?, state=?, id_user=?, id_wp_term=? WHERE id=?")
	insertarRegistros.Exec(newService.Name, newService.Description, newService.State, newService.Id_user, newService.Id_wp_term, id)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func DeleteService(id string) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("DELETE FROM service WHERE id=?")
	insertarRegistros.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

// Funcion que obtiene los servicios limpios de WordPress y los entrega en formato json
func GetServicesWordPress() (models.Services, error) {
	conexionEstablecida := database.ConexionDBWP()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT name, description, wp_terms.term_id as term_id FROM wp_terms INNER JOIN wp_term_taxonomy ON wp_terms.term_id = wp_term_taxonomy.term_id WHERE taxonomy = 'category' AND LENGTH(description) > 0")
	if err != nil {
		return nil, err
	}
	service := models.Service{}
	services := models.Services{}
	for obtenerRegistros.Next() {
		var name, description string
		var term_id int
		obtenerRegistros.Scan(&name, &description, &term_id)
		service.Name = name
		service.Description = description
		service.Id_wp_term = term_id
		service.State = "Pendiente"
		services = append(services, service)
	}
	// fmt.Println(services)
	return services, nil
}

// Funcion que recibe los servicios limpios y los ingresa en la base de datos del middleware
func InsertServices() (models.Services, error) {
	services, err := GetServicesWordPress()
	if err != nil {
		panic(err.Error())
	}
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO service(name, description, state, id_user, id_wp_term) VALUES(?,?,?,?,?)")
	for _, service := range services {
		insertarRegistros.Exec(service.Name, service.Description, service.State, 0, service.Id_wp_term)
		if err != nil {
			panic(err.Error())
		}
	}
	return services, nil
}

// Funcion que actualiza la descripcion de un servicio en la base de datos de WordPress segun su id_wp_term
func ApproveServiceWP(id string, newService models.Service) error {
	conexionEstablecida := database.ConexionDBWP()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("UPDATE wp_term_taxonomy SET description =? WHERE term_id =?")
	insertarRegistros.Exec(newService.Description, id)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func GetServicesByUser(id string) (models.Services, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM service WHERE id_user = ?", id)
	service := models.Service{}
	services := models.Services{}
	for obtenerRegistros.Next() {
		if err != nil {
			panic(err.Error())
		}
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
	return services, nil
}

func AssignServices(id_user int, services []int) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("UPDATE service SET id_user =? WHERE id =?")
	for _, service := range services {
		insertarRegistros.Exec(id_user, service)
		if err != nil {
			panic(err.Error())
		}
	}
	return nil
}

func CountServicesByUser(id string) (int, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT COUNT(*) FROM service WHERE id_user = ?", id)
	var count int
	for obtenerRegistros.Next() {
		obtenerRegistros.Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	}
	return count, nil
}
