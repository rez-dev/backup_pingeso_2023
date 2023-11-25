package services

import (
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
	"github.com/rez-dev/backup_pingeso_2023/backend/repositories"
)

func GetServices() (models.Services, error) {
	services, err := repositories.GetServices()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func GetService(id string) (models.Services, error) {
	services, err := repositories.GetService(id)
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

func GetServicesWordPress() (models.Services, error) {
	services, err := repositories.GetServicesWordPress()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func InsertServices() (models.Services, error) {
	services, err := repositories.InsertServices()
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

func ApproveService(id string, id_user int, description string) error {
	err := repositories.ApproveService(id, id_user, description)
	if err != nil {
		return err
	}
	return nil
}

func GetServicesByUser(id string) (models.Services, error) {
	services, err := repositories.GetServicesByUser(id)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func AssignServices(id_user int, servicios []int) error {
	err := repositories.AssignServices(id_user, servicios)
	if err != nil {
		return err
	}
	return nil
}

func DesassignService(id string) error {
	err := repositories.DesassignService(id)
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

func GetServicesByState(state string) (models.Services, error) {
	services, err := repositories.GetServicesByState(state)
	if err != nil {
		return nil, err
	}
	return services, nil
}
