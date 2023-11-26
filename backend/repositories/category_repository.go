package repositories

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rez-dev/backup_pingeso_2023/backend/database"
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
)

func GetCategories() (models.Categories, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	category := models.Category{}
	categories := models.Categories{}
	for obtenerRegistros.Next() {
		var id, id_wp_term int
		var level1, level2, level3 string
		obtenerRegistros.Scan(&id, &level1, &level2, &level3, &id_wp_term)
		category.ID = id
		category.Level1 = level1
		category.Level2 = level2
		category.Level3 = level3
		category.Id_wp_term = id_wp_term
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategory(id string) (models.Categories, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM category WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	category := models.Category{}
	categories := models.Categories{}
	for obtenerRegistros.Next() {
		var id int
		var id_wp_term int
		var level1, level2, level3 string
		obtenerRegistros.Scan(&id, &level1, &level2, &level3, &id_wp_term)
		category.ID = id
		category.Level1 = level1
		category.Level2 = level2
		category.Level3 = level3
		category.Id_wp_term = id_wp_term
		categories = append(categories, category)
	}
	return categories, nil
}

func CreateCategory(newCategory models.Category) error {
	conexionEstablecida := database.ConexionDB()
	// insertar datos
	insertarDatos, err := conexionEstablecida.Prepare("INSERT INTO category(level1, level2, level3, id_wp_term) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	insertarDatos.Exec(newCategory.Level1, newCategory.Level2, newCategory.Level3, newCategory.Id_wp_term)
	return nil
}

func UpdateCategory(id string, newCategory models.Category) error {
	conexionEstablecida := database.ConexionDB()
	// actualizar datos
	actualizarDatos, err := conexionEstablecida.Prepare("UPDATE category SET level1 = ?, level2 = ?, level3 = ?, id_wp_term = ? WHERE id = ?")
	if err != nil {
		return err
	}
	actualizarDatos.Exec(newCategory.Level1, newCategory.Level2, newCategory.Level3, newCategory.Id_wp_term, id)
	return nil
}

func DeleteCategory(id string) error {
	conexionEstablecida := database.ConexionDB()
	// eliminar datos
	eliminarDatos, err := conexionEstablecida.Prepare("DELETE FROM category WHERE id = ?")
	if err != nil {
		return err
	}
	eliminarDatos.Exec(id)
	return nil
}

// Funcion que obtiene las categorias de los servicios de WordPress de UN servicio y las entrega en formato lista
func GetParentsWordPress(id int) ([3]string, error) {
	conexionEstablecida := database.ConexionDBWP()
	// obtener registros que tienen el mismo nombre
	obtenerPadre1, err := conexionEstablecida.Query("SELECT parent_term.term_id AS id_categoria_padre, parent_term.name AS nombre_categoria_padre FROM wp_term_taxonomy AS child_taxonomy JOIN wp_terms AS child_term ON child_taxonomy.term_id = child_term.term_id LEFT JOIN wp_term_taxonomy AS parent_taxonomy ON child_taxonomy.parent = parent_taxonomy.term_id LEFT JOIN wp_terms AS parent_term ON parent_taxonomy.term_id = parent_term.term_id WHERE child_term.term_id = ?", id)
	if err != nil {
		return [3]string{}, err
	}
	var padre1, padre2, padre3 string
	var id_padre1, id_padre2, id_padre3 int
	for obtenerPadre1.Next() {
		obtenerPadre1.Scan(&id_padre1, &padre1)
	}
	obtenerPadre2, err := conexionEstablecida.Query("SELECT parent_term.term_id AS id_categoria_padre, parent_term.name AS nombre_categoria_padre FROM wp_term_taxonomy AS child_taxonomy JOIN wp_terms AS child_term ON child_taxonomy.term_id = child_term.term_id LEFT JOIN wp_term_taxonomy AS parent_taxonomy ON child_taxonomy.parent = parent_taxonomy.term_id LEFT JOIN wp_terms AS parent_term ON parent_taxonomy.term_id = parent_term.term_id WHERE child_term.term_id = ?", id_padre1)
	if err != nil {
		return [3]string{}, err
	}
	for obtenerPadre2.Next() {
		obtenerPadre2.Scan(&id_padre2, &padre2)
	}
	obtenerPadre3, err := conexionEstablecida.Query("SELECT parent_term.term_id AS id_categoria_padre, parent_term.name AS nombre_categoria_padre FROM wp_term_taxonomy AS child_taxonomy JOIN wp_terms AS child_term ON child_taxonomy.term_id = child_term.term_id LEFT JOIN wp_term_taxonomy AS parent_taxonomy ON child_taxonomy.parent = parent_taxonomy.term_id LEFT JOIN wp_terms AS parent_term ON parent_taxonomy.term_id = parent_term.term_id WHERE child_term.term_id = ?", id_padre2)
	if err != nil {
		return [3]string{}, err
	}
	for obtenerPadre3.Next() {
		obtenerPadre3.Scan(&id_padre3, &padre3)
	}
	var datos [3]string
	datos[0] = padre3
	datos[1] = padre2
	datos[2] = padre1

	return datos, nil
}

// // Funcion que obtiene las categorias de los servicios de WordPress y las inserta en la base de datos del middleware
// func InsertCategories() (models.Categories, error) {
// 	conexionEstablecida := database.ConexionDBWP()
// 	// obtener registros que tienen el mismo nombre
// 	obtenerRegistros, err := conexionEstablecida.Query("SELECT wp_terms.term_id as term_id FROM wp_terms INNER JOIN wp_term_taxonomy ON wp_terms.term_id = wp_term_taxonomy.term_id WHERE taxonomy = 'category' AND LENGTH(description) > 0")
// 	if err != nil {
// 		return nil, err
// 	}
// 	var terms []int
// 	for obtenerRegistros.Next() {
// 		var term_id int
// 		obtenerRegistros.Scan(&term_id)
// 		terms = append(terms, term_id)
// 	}
// 	category := models.Category{}
// 	categories := models.Categories{}
// 	var padres [3]string
// 	for i := 0; i < len(terms); i++ {
// 		padres, err = GetParentsWordPress(terms[i])
// 		if err != nil {
// 			return nil, err
// 		}
// 		category.Level1 = padres[0]
// 		category.Level2 = padres[1]
// 		category.Level3 = padres[2]
// 		category.Id_wp_term = terms[i]
// 		categories = append(categories, category)
// 		// fmt.Println(padres)
// 	}

// 	conexionEstablecida2 := database.ConexionDB()
// 	insertarRegistros, err := conexionEstablecida2.Prepare("INSERT INTO category(level1, level2, level3, id_wp_term) VALUES(?,?,?,?)")
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, category := range categories {
// 		insertarRegistros.Exec(category.Level1, category.Level2, category.Level3, category.Id_wp_term)
// 	}
// 	return categories, nil
// }

func InsertCategories() (models.Categories, error) {
	conexionEstablecida := database.ConexionDBWP()
	obtenerRegistros, err := conexionEstablecida.Query("SELECT wp_terms.term_id as term_id FROM wp_terms INNER JOIN wp_term_taxonomy ON wp_terms.term_id = wp_term_taxonomy.term_id WHERE taxonomy = 'category' AND LENGTH(description) > 0")
	if err != nil {
		return nil, err
	}
	var terms []int
	for obtenerRegistros.Next() {
		var term_id int
		obtenerRegistros.Scan(&term_id)
		terms = append(terms, term_id)
	}
	category := models.Category{}
	categories := models.Categories{}
	var padres [3]string
	for i := 0; i < len(terms); i++ {
		padres, err = GetParentsWordPress(terms[i])
		if err != nil {
			return nil, err
		}
		category.Level1 = padres[0]
		category.Level2 = padres[1]
		category.Level3 = padres[2]
		category.Id_wp_term = terms[i]

		// Verificar si el registro ya existe en la base de datos
		if !categoryExists(category.Id_wp_term) {
			categories = append(categories, category)
		}
	}

	conexionEstablecida2 := database.ConexionDB()
	insertarRegistros, err := conexionEstablecida2.Prepare("INSERT INTO category(level1, level2, level3, id_wp_term) VALUES(?,?,?,?)")
	if err != nil {
		return nil, err
	}
	for _, category := range categories {
		_, err := insertarRegistros.Exec(category.Level1, category.Level2, category.Level3, category.Id_wp_term)
		if err != nil {
			return nil, err
		}
	}
	return categories, nil
}

// Función para verificar si la categoría ya existe en la base de datos
func categoryExists(id int) bool {
	conexionEstablecida := database.ConexionDB()
	var count int
	err := conexionEstablecida.QueryRow("SELECT COUNT(*) FROM category WHERE id_wp_term = ?", id).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func GetCategoriesByTerm(id string) (models.Categories, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM category WHERE id_wp_term = ?", id)
	if err != nil {
		return nil, err
	}
	category := models.Category{}
	categories := models.Categories{}
	for obtenerRegistros.Next() {
		var id, id_wp_term int
		var level1, level2, level3 string
		obtenerRegistros.Scan(&id, &level1, &level2, &level3, &id_wp_term)
		category.ID = id
		category.Level1 = level1
		category.Level2 = level2
		category.Level3 = level3
		category.Id_wp_term = id_wp_term
		categories = append(categories, category)
	}
	return categories, nil
}
