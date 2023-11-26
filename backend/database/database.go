package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConexionDB() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := "mysql"
	Nombre := "bd_test1"
	// Tipo_conexion := "127.0.0.1"

	// Cadena de conexión
	cadenaConexion := fmt.Sprintf("%s:%s@tcp(db:3306)/", Usuario, Contrasenia)

	conexion, err := sql.Open(Driver, cadenaConexion)

	if err != nil {
		panic(err.Error())
	}

	// crear base de datos
	_, err = conexion.Exec("CREATE DATABASE IF NOT EXISTS " + Nombre)
	if err != nil {
		panic(err.Error())
	}

	// usar base de datos
	_, err = conexion.Exec("USE " + Nombre)
	if err != nil {
		panic(err.Error())
	}
	// crear todas las tablas si no existen
	_, err = conexion.Exec(`CREATE TABLE IF NOT EXISTS service (
		id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		state VARCHAR(255),
		id_user INT NOT NULL,
		id_wp_term INT NOT NULL,
		last_approval VARCHAR(255)
	)`)
	if err != nil {
		panic(err.Error())
	}
	_, err = conexion.Exec(`CREATE TABLE IF NOT EXISTS category (
		id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		level1 VARCHAR(255) NOT NULL,
		level2 VARCHAR(255),
		level3 VARCHAR(255),
		id_wp_term INT NOT NULL
		)`)
	if err != nil {
		panic(err.Error())
	}
	_, err = conexion.Exec(`CREATE TABLE IF NOT EXISTS information (
        id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT,
        position INT,
        author VARCHAR(255),
        id_wp_post INT NOT NULL,
        id_post_meta INT NOT NULL,
        id_author INT NOT NULL,
        id_wp_term INT NOT NULL
    )`)
	if err != nil {
		panic(err.Error())
	}
	_, err = conexion.Exec(`CREATE TABLE IF NOT EXISTS user (
		id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		rut VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(255) NOT NULL,
		unity VARCHAR(255) NOT NULL
	)`)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

// func ConexionDBWP() (conexion *sql.DB) {
// 	Driver := "mysql"
// 	Usuario := "root"
// 	Contrasenia := "mysql"
// 	Nombre := "test_usachatiendebd"
// 	// Tipo_conexion := "127.0.0.1"
// 	// Cadena de conexión
// 	cadenaConexion := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3333)/%s", Usuario, Contrasenia, Nombre)
// 	conexion, err := sql.Open(Driver, cadenaConexion)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return conexion
// }

func ConexionDBWP() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "etl"
	Contrasenia := "etl"
	Nombre := "test_usachatiendebd"
	// Cadena de conexión
	// cadenaConexion := fmt.Sprintf("%s:%s@tcp(172.23.64.1:3333)/%s", Usuario, Contrasenia, Nombre)
	cadenaConexion := fmt.Sprintf("%s:%s@tcp(172.23.64.1:3333)/%s", Usuario, Contrasenia, Nombre)
	conexion, err := sql.Open(Driver, cadenaConexion)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}
