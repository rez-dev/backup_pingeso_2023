package services

import (
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
	"github.com/rez-dev/backup_pingeso_2023/backend/repositories"
)

func GetUsers() (models.Users, error) {
	users, err := repositories.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(id string) (models.Users, error) {
	users, err := repositories.GetUser(id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user models.User) error {
	err := repositories.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(id string, user models.User) error {
	err := repositories.UpdateUser(id, user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	err := repositories.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func GetUsersByUnity(unity string, role string) (models.Users, error) {
	users, err := repositories.GetUsersByUnity(unity, role)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByRole(role string) (models.Users, error) {
	users, err := repositories.GetUserByRole(role)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func Signup(user models.User) error {
	err := repositories.Signup(user)
	if err != nil {
		return err
	}
	return nil
}

func Login(email string, password string) (string, error) {
	token, err := repositories.Login(email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
