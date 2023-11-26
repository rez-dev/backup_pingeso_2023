package repositories

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rez-dev/backup_pingeso_2023/backend/database"
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() (models.Users, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}
	user := models.User{}
	users := models.Users{}
	for obtenerRegistros.Next() {
		var id int
		var name, email, password, role string
		obtenerRegistros.Scan(&id, &name, &email, &password, &role)
		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.Role = role
		users = append(users, user)
	}
	return users, nil
}

func GetUser(id string) (models.Users, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM user WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	user := models.User{}
	users := models.Users{}
	for obtenerRegistros.Next() {
		var id int
		var name, email, password, role string
		obtenerRegistros.Scan(&id, &name, &email, &password, &role)
		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.Role = role
		users = append(users, user)
	}
	return users, nil
}

func CreateUser(user models.User) error {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO user (name, email, password, role) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	insertarRegistros.Exec(user.Name, user.Email, user.Password, user.Role)
	return nil
}

func UpdateUser(id string, user models.User) error {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	actualizarRegistros, err := conexionEstablecida.Prepare("UPDATE user SET name = ?, email = ?, password = ?, role = ? WHERE id = ?")
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	actualizarRegistros.Exec(user.Name, user.Email, user.Password, user.Role, id)
	return nil
}

func DeleteUser(id string) error {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	eliminarRegistros, err := conexionEstablecida.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return err
	}
	eliminarRegistros.Exec(id)
	return nil
}

func GetUsersByUnity(unity string, role string) (models.Users, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM user WHERE unity = ? AND role = ?", unity, role)
	if err != nil {
		return nil, err
	}
	if role == "" {
		obtenerRegistros, err = conexionEstablecida.Query("SELECT * FROM user WHERE unity = ?", unity)
		if err != nil {
			return nil, err
		}
	}
	user := models.User{}
	users := models.Users{}
	for obtenerRegistros.Next() {
		var id int
		var name, email, password, role string
		obtenerRegistros.Scan(&id, &name, &email, &password, &role)
		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.Role = role
		users = append(users, user)
	}
	return users, nil
}

func GetUserByRole(role string) (models.Users, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT * FROM user WHERE role = ?", role)
	if err != nil {
		return nil, err
	}
	user := models.User{}
	users := models.Users{}
	for obtenerRegistros.Next() {
		var id int
		var name, email, password, role string
		obtenerRegistros.Scan(&id, &name, &email, &password, &role)
		user.ID = id
		user.Name = name
		user.Email = email
		user.Password = password
		user.Role = role
		users = append(users, user)
	}
	return users, nil
}

func Signup(user models.User) error {
	var signupReq struct {
		Email    string
		Password string
		Name     string
		Rut      string
		Role     string
		Unity    string
	}

	signupReq.Email = user.Email
	signupReq.Password = user.Password
	signupReq.Name = user.Name
	signupReq.Rut = user.Rut
	signupReq.Role = user.Role
	signupReq.Unity = user.Unity

	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO user (name, rut, email, password, role, unity) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	insertarRegistros.Exec(signupReq.Name, signupReq.Rut, signupReq.Email, signupReq.Password, signupReq.Role, signupReq.Unity)
	return nil
}

func Login(email string, password string) (string, error) {
	conexionEstablecida := database.ConexionDB()
	// obtener registros que tienen el mismo nombre
	obtenerRegistros, err := conexionEstablecida.Query("SELECT password FROM user WHERE email = ?", email)
	if err != nil {
		return "", err
	}
	// obtener password
	var passwordDB string
	for obtenerRegistros.Next() {
		obtenerRegistros.Scan(&passwordDB)
	}
	// comparar password
	err = bcrypt.CompareHashAndPassword([]byte(passwordDB), []byte(password))
	if err != nil {
		return "", err
	}

	// get user from db
	var user models.User
	err = conexionEstablecida.QueryRow("SELECT * FROM user WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Rut, &user.Email, &user.Password, &user.Role, &user.Unity)
	if err != nil {
		return "", err
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	//mostrar por consola el token
	//fmt.Println(token)

	//sign and get the complete encoded token as a string using the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
