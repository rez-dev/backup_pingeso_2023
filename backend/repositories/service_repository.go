package repositories

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rez-dev/backup_pingeso_2023/backend/database"
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
)

func GetServices() (models.Services, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM service")
	if err != nil {
		return nil, err
	}
	service := models.Service{}
	services := models.Services{}
	for obtenerRegistros.Next() {
		var id int
		var id_user, id_wp_term int
		var name, description, state, last_approval string
		obtenerRegistros.Scan(&id, &name, &description, &state, &id_user, &id_wp_term, &last_approval)
		service.ID = id
		service.Name = name
		service.Description = description
		service.State = state
		service.Id_user = id_user
		service.Id_wp_term = id_wp_term
		service.Last_approval = last_approval
		services = append(services, service)
	}
	return services, nil
}

func GetService(id string) (models.Services, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM service WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	service := models.Service{}
	services := models.Services{}
	for obtenerRegistros.Next() {
		var id int
		var id_user, id_wp_term int
		var name, description, state, last_approval string
		obtenerRegistros.Scan(&id, &name, &description, &state, &id_user, &id_wp_term, &last_approval)
		service.ID = id
		service.Name = name
		service.Description = description
		service.State = state
		service.Id_user = id_user
		service.Id_wp_term = id_wp_term
		service.Last_approval = last_approval
		services = append(services, service)
	}
	return services, nil
}

func CreateService(newService models.Service) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO service(name, description, state, id_user, id_wp_term, last_approval) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	insertarRegistros.Exec(newService.Name, newService.Description, newService.State, newService.Id_user, newService.Id_wp_term, newService.Last_approval)
	return nil
}

func UpdateService(id string, newService models.Service) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("UPDATE service SET name=?, description=?, state=?, id_user=?, id_wp_term=?, last_approval=? WHERE id=?")
	if err != nil {
		return err
	}
	insertarRegistros.Exec(newService.Name, newService.Description, newService.State, newService.Id_user, newService.Id_wp_term, newService.Last_approval, id)
	return nil
}

func DeleteService(id string) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("DELETE FROM service WHERE id=?")
	if err != nil {
		return err
	}
	insertarRegistros.Exec(id)
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
		service.Id_user = 0
		service.Last_approval = "Pendiente"
		services = append(services, service)
	}
	return services, nil
}

// // Funcion que recibe los servicios limpios y los ingresa en la base de datos del middleware
// func InsertServices() (models.Services, error) {
// 	services, err := GetServicesWordPress()
// 	if err != nil {
// 		return nil, err
// 	}
// 	conexionEstablecida := database.ConexionDB()
// 	// insertar registros
// 	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO service(name, description, state, id_user, id_wp_term, last_approval) VALUES(?,?,?,?,?,?)")
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, service := range services {
// 		insertarRegistros.Exec(service.Name, service.Description, service.State, service.Id_user, service.Id_wp_term, service.Last_approval)
// 	}
// 	return services, nil
// }

// Función que recibe los servicios limpios y los ingresa en la base de datos del middleware
func InsertServices() (models.Services, error) {
	services, err := GetServicesWordPress()
	if err != nil {
		return nil, err
	}
	conexionEstablecida := database.ConexionDB()

	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO service(name, description, state, id_user, id_wp_term, last_approval) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return nil, err
	}
	insertedServices := models.Services{} // Lista para almacenar los servicios insertados

	for _, service := range services {
		// Verificar si el servicio ya existe en la base de datos
		if !serviceExists(service.Id_wp_term) {
			// Insertar el servicio si no existe
			_, err := insertarRegistros.Exec(service.Name, service.Description, service.State, service.Id_user, service.Id_wp_term, service.Last_approval)
			if err != nil {
				return nil, err
			}
			insertedServices = append(insertedServices, service)
		}
	}

	return insertedServices, nil
}

// Función para verificar si el servicio ya existe en la base de datos
func serviceExists(id int) bool {
	conexionEstablecida := database.ConexionDB()
	var count int
	err := conexionEstablecida.QueryRow("SELECT COUNT(*) FROM service WHERE id_wp_term = ?", id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func ApproveServiceWP(id string, newService models.Service) error {
	conexionEstablecida := database.ConexionDBWP()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("UPDATE wp_term_taxonomy SET description =? WHERE term_id =?")
	if err != nil {
		return err
	}
	insertarRegistros.Exec(newService.Description, id)
	return nil
}

func ApproveService(id string, id_user int, description string) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("UPDATE service SET description =?, state =?, id_user=?, last_approval=? WHERE id_wp_term =?")
	if err != nil {
		return err
	}
	fecha := time.Now().Format("2006-01-02 15:04:05")
	insertarRegistros.Exec(description, "Aprobado", id_user, fecha, id)
	return nil
}

func GetServicesByUser(id string) (models.Services, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM service WHERE id_user = ?", id)
	if err != nil {
		return nil, err
	}
	service := models.Service{}
	services := models.Services{}
	for obtenerRegistros.Next() {
		var id int
		var id_user, id_wp_term int
		var name, description, state, last_approval string
		obtenerRegistros.Scan(&id, &name, &description, &state, &id_user, &id_wp_term, &last_approval)
		service.ID = id
		service.Name = name
		service.Description = description
		service.State = state
		service.Id_user = id_user
		service.Id_wp_term = id_wp_term
		service.Last_approval = last_approval
		services = append(services, service)
	}
	return services, nil
}

func AssignServices(id_user int, servicios []int) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Prepare("UPDATE service SET id_user =?, state='Asignado' WHERE id =?")
	if err != nil {
		return err
	}
	for _, servicio := range servicios {
		insertarRegistros.Exec(id_user, servicio)
	}
	return nil
}

func DesassignService(id string) error {
	conexionEstablecida := database.ConexionDB()
	// insertar registros
	insertarRegistros, err := conexionEstablecida.Query("UPDATE service SET id_user = ?, state='Pendiente' WHERE id =?", 0, id)
	if err != nil {
		return err
	}
	insertarRegistros.Scan()
	return nil
}

func CountServicesByUser(id string) (int, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT COUNT(*) FROM service WHERE id_user = ? AND state='Asignado'", id)
	if err != nil {
		return 0, err
	}
	var count int
	for obtenerRegistros.Next() {
		obtenerRegistros.Scan(&count)
	}
	return count, nil
}

func GetServicesByState(state string) (models.Services, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM service WHERE state = ?", state)
	if err != nil {
		return nil, err
	}
	service := models.Service{}
	services := models.Services{}
	for obtenerRegistros.Next() {
		var id int
		var id_user, id_wp_term int
		var name, description, state, last_approval string
		obtenerRegistros.Scan(&id, &name, &description, &state, &id_user, &id_wp_term, &last_approval)
		service.ID = id
		service.Name = name
		service.Description = description
		service.State = state
		service.Id_user = id_user
		service.Id_wp_term = id_wp_term
		service.Last_approval = last_approval
		services = append(services, service)
	}
	return services, nil
}

// func GetUserByService(id string) (models.Users, error) {
// 	conexionEstablecida := database.ConexionDB()
// 	db := conexionDBUser() // Make sure this returns *gorm.DB
// 	// obtener id_user de servicio con id = id
// 	obtenerRegistros, _ := conexionEstablecida.Query("SELECT id_user FROM service WHERE id = ?", id)
// 	var id_user int
// 	for obtenerRegistros.Next() {
// 		obtenerRegistros.Scan(&id_user)
// 	}
// 	// obtener registros que tienen el mismo nombre
// 	var user models.User
// 	// Use GORM's First method to find the user by ID
// 	db.Where("id = ?", id_user).First(&user)
// 	var name string = user.Name
// 	c.IndentedJSON(http.StatusOK, gin.H{"name": name})
// }
