package services

import (
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
	"github.com/rez-dev/backup_pingeso_2023/backend/repositories"
)

func GetInformations() (models.Informations, error) {
	informations, err := repositories.GetInformations()
	if err != nil {
		return nil, err
	}
	return informations, nil
}

func GetInformation(id string) (models.Informations, error) {
	informations, err := repositories.GetInformation(id)
	if err != nil {
		return nil, err
	}
	return informations, nil
}

func CreateInformation(information models.Information) error {
	err := repositories.CreateInformation(information)
	if err != nil {
		return err
	}
	return nil
}

func UpdateInformation(id string, information models.Information) error {
	err := repositories.UpdateInformation(id, information)
	if err != nil {
		return err
	}
	return nil
}

func DeleteInformation(id string) error {
	err := repositories.DeleteInformation(id)
	if err != nil {
		return err
	}
	return nil
}

// func GetInformationsWordPress() (models.Informations, error) {
// 	informations, err := repositories.GetInformationsWordPress()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return informations, nil
// }

func InsertInformations() (models.Informations, error) {
	informations, err := repositories.InsertInformations()
	if err != nil {
		return nil, err
	}
	return informations, nil
}

func GetInformationsByTerm(id string) (models.Informations, error) {
	informations, err := repositories.GetInformationsByTerm(id)
	if err != nil {
		return nil, err
	}
	return informations, nil
}

func ApproveInformationWP(id string, information models.Information) error {
	err := repositories.ApproveInformationWP(id, information)
	if err != nil {
		return err
	}
	return nil
}
