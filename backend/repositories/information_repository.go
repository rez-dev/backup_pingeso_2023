package repositories

import (
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rez-dev/backup_pingeso_2023/backend/database"
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
)

func GetInformations() (models.Informations, error) {
	conexionEstablecida := database.ConexionDB()
	defer conexionEstablecida.Close() // Asegúrate de cerrar la conexión cuando hayas terminado

	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM information")
	if err != nil {
		return nil, err
	}
	information := models.Information{}
	informations := models.Informations{}
	for obtenerRegistros.Next() {
		var id, id_wp_post, id_post_meta, id_author, id_wp_term, position int
		var title, content, author string
		obtenerRegistros.Scan(&id, &title, &content, &position, &author, &id_wp_post, &id_post_meta, &id_author, &id_wp_term)
		information.ID = id
		information.Title = title
		information.Content = content
		information.Position = position
		information.Author = author
		information.Id_wp_post = id_wp_post
		information.Id_post_meta = id_post_meta
		information.Id_author = id_author
		information.Id_wp_term = id_wp_term
		informations = append(informations, information)
	}
	return informations, nil
}

func GetInformation(id string) (models.Informations, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM information WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	information := models.Information{}
	informations := models.Informations{}
	for obtenerRegistros.Next() {
		var id, id_wp_post, id_post_meta, id_author, id_wp_term, position int
		var title, content, author string
		obtenerRegistros.Scan(&id, &title, &content, &position, &author, &id_wp_post, &id_post_meta, &id_author, &id_wp_term)
		information.ID = id
		information.Title = title
		information.Content = content
		information.Position = position
		information.Author = author
		information.Id_wp_post = id_wp_post
		information.Id_post_meta = id_post_meta
		information.Id_author = id_author
		information.Id_wp_term = id_wp_term
		informations = append(informations, information)
	}
	return informations, nil
}

func CreateInformation(newInformation models.Information) error {
	conexionEstablecida := database.ConexionDB()
	// insertar datos
	insertarDatos, err := conexionEstablecida.Prepare("INSERT INTO information(title, content, position, author, id_wp_post, id_post_meta, id_author, id_wp_term) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	insertarDatos.Exec(newInformation.Title, newInformation.Content, newInformation.Position, newInformation.Author, newInformation.Id_wp_post, newInformation.Id_post_meta, newInformation.Id_author, newInformation.Id_wp_term)
	return nil
}

func UpdateInformation(id string, newInformation models.Information) error {
	conexionEstablecida := database.ConexionDB()
	// actualizar datos
	actualizarDatos, err := conexionEstablecida.Prepare("UPDATE information SET title=?, content=?, position=?, author=?, id_wp_post=?, id_post_meta=?, id_author=?, id_wp_term=? WHERE id=?")
	if err != nil {
		return err
	}
	actualizarDatos.Exec(newInformation.Title, newInformation.Content, newInformation.Position, newInformation.Author, newInformation.Id_wp_post, newInformation.Id_post_meta, newInformation.Id_author, newInformation.Id_wp_term, id)
	return nil
}

func DeleteInformation(id string) error {
	conexionEstablecida := database.ConexionDB()
	// eliminar datos
	eliminarDatos, _ := conexionEstablecida.Prepare("DELETE FROM information WHERE id=?")
	eliminarDatos.Exec(id)
	return nil
}

// Funcion que obtiene las informaciones limpias de WordPress y los entrega en un json de informations
func GetInformationWordPress() models.Informations {
	conexionEstablecida := database.ConexionDBWP()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT wp_posts.ID AS id_wp_post, wp_postmeta.meta_id AS id_post_meta, wp_users.ID AS id_author, wp_terms.term_id AS id_term, post_title, post_content, meta_value, user_nicename FROM wp_posts INNER JOIN wp_postmeta ON wp_posts.ID = wp_postmeta.post_id INNER JOIN wp_users ON wp_posts.post_author = wp_users.ID INNER JOIN wp_term_relationships ON wp_posts.ID = wp_term_relationships.object_id INNER JOIN wp_term_taxonomy ON wp_term_relationships.term_taxonomy_id = wp_term_taxonomy.term_taxonomy_id INNER JOIN wp_terms ON wp_term_taxonomy.term_id = wp_terms.term_id WHERE wp_postmeta.meta_key = 'orden'")
	if err != nil {
		panic(err.Error())
	}
	information := models.Information{}
	informations := models.Informations{}
	for obtenerRegistros.Next() {
		var id_wp_post, id_post_meta, id_author, id_term int
		var post_title, post_content, meta_value, user_nicename string
		obtenerRegistros.Scan(&id_wp_post, &id_post_meta, &id_author, &id_term, &post_title, &post_content, &meta_value, &user_nicename)
		information.Id_wp_post = id_wp_post
		information.Id_post_meta = id_post_meta
		information.Id_author = id_author
		information.Id_wp_term = id_term
		information.Title = post_title
		information.Content = post_content
		information.Position, _ = strconv.Atoi(meta_value)
		information.Author = user_nicename
		informations = append(informations, information)
	}
	return informations
}

// // Funcion que obtiene las informaciones limpias de WordPress y las inserta en la base de datos del Middleware
// func InsertInformations() (models.Informations, error) {
// 	conexionEstablecida := database.ConexionDB()
// 	informations := GetInformationWordPress()
// 	// insertar datos
// 	for _, information := range informations {
// 		insertarDatos, err := conexionEstablecida.Prepare("INSERT INTO information(title, content, position, author, id_wp_post, id_post_meta, id_author, id_wp_term) VALUES(?,?,?,?,?,?,?,?)")
// 		if err != nil {
// 			return nil, err
// 		}
// 		insertarDatos.Exec(information.Title, information.Content, information.Position, information.Author, information.Id_wp_post, information.Id_post_meta, information.Id_author, information.Id_wp_term)
// 	}
// 	return informations, nil
// }

// Función que obtiene las informaciones limpias de WordPress y las inserta en la base de datos del Middleware
func InsertInformations() (models.Informations, error) {
	conexionEstablecida := database.ConexionDB()
	informations := GetInformationWordPress()
	insertedInformation := models.Informations{} // Lista para almacenar las informaciones insertadas

	// Iterar sobre la información obtenida
	for _, information := range informations {
		// Verificar si la información ya existe en la base de datos
		if !informationExists(information.Id_wp_post) {
			// Insertar la información si no existe
			insertarDatos, err := conexionEstablecida.Prepare("INSERT INTO information(title, content, position, author, id_wp_post, id_post_meta, id_author, id_wp_term) VALUES(?,?,?,?,?,?,?,?)")
			if err != nil {
				return nil, err
			}
			_, err = insertarDatos.Exec(information.Title, information.Content, information.Position, information.Author, information.Id_wp_post, information.Id_post_meta, information.Id_author, information.Id_wp_term)
			if err != nil {
				return nil, err
			}
			insertedInformation = append(insertedInformation, information)
		}
	}
	return insertedInformation, nil
}

// Función para verificar si la información ya existe en la base de datos
func informationExists(id int) bool {
	conexionEstablecida := database.ConexionDB()
	var count int
	err := conexionEstablecida.QueryRow("SELECT COUNT(*) FROM information WHERE id_wp_post = ?", id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

// Funcion que obtiene todas las informaciones de UN servicio y la entrega en formato json
func GetInformationsByTerm(id string) (models.Informations, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM information WHERE id_wp_term = ?", id)
	if err != nil {
		return nil, err
	}
	information := models.Information{}
	informations := models.Informations{}
	for obtenerRegistros.Next() {
		var id, id_wp_post, id_post_meta, id_author, id_wp_term, position int
		var title, content, author string
		obtenerRegistros.Scan(&id, &title, &content, &position, &author, &id_wp_post, &id_post_meta, &id_author, &id_wp_term)
		information.ID = id
		information.Title = title
		information.Content = content
		information.Position = position
		information.Author = author
		information.Id_wp_post = id_wp_post
		information.Id_post_meta = id_post_meta
		information.Id_author = id_author
		information.Id_wp_term = id_wp_term
		informations = append(informations, information)
	}
	return informations, nil
}

// Funcion que actualiza los valores de post_content, post_title y post_status de wp_posts segun id_wp_post y actualiza meta_value de wp_postmeta segun id_wp_post
func ApproveInformationWP(id string, updatedInformation models.Information) error {
	conexionEstablecida := database.ConexionDBWP()
	// Actualiza post_content, post_title y post_status de wp_posts segun id_wp_post
	insertarRegistros, err := conexionEstablecida.Prepare("UPDATE wp_posts SET post_content=?, post_title=?, post_status = 'aprobado' WHERE ID=?")
	if err != nil {
		return err
	}
	insertarRegistros.Exec(updatedInformation.Content, updatedInformation.Title, id)

	// Actualiza meta_value de wp_postmeta

	insertarRegistros, _ = conexionEstablecida.Prepare("UPDATE wp_postmeta SET meta_value=? WHERE post_id=? AND meta_key='orden'")
	insertarRegistros.Exec(updatedInformation.Position, id)
	return nil
}
