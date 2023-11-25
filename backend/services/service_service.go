package services

import (
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
	"github.com/rez-dev/backup_pingeso_2023/backend/repositories"
)

func GetServices() (services []models.Service, err error) {
	services, err = repositories.GetServices()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func GetService(id string) (services []models.Service, err error) {
	services, err = repositories.GetService(id)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func CreateService(newService models.Service) error {
	err := repositories.CreateService(newService)
	if err != nil {
		return err
	}
	return nil
}

func UpdateService(id string, newService models.Service) error {
	err := repositories.UpdateService(id, newService)
	if err != nil {
		return err
	}
	return nil
}

func DeleteService(id string) error {
	err := repositories.DeleteService(id)
	if err != nil {
		return err
	}
	return nil
}

func GetServicesWordPress() (services []models.Service, err error) {
	services, err = repositories.GetServicesWordPress()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func InsertServices() (services []models.Service, err error) {
	services, err = repositories.InsertServices()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func ApproveServiceWP(id string, newService models.Service) error {
	err := repositories.ApproveServiceWP(id, newService)
	if err != nil {
		return err
	}
	return nil
}

func GetServicesByUser(id string) (services []models.Service, err error) {
	services, err = repositories.GetServicesByUser(id)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func AssignServices(id int, services []int) error {
	err := repositories.AssignServices(id, services)
	if err != nil {
		return err
	}
	return nil
}

func CountServicesByUser(id string) (int, error) {
	count, err := repositories.CountServicesByUser(id)
	if err != nil {
		return 0, err
	}
	return count, nil
}
