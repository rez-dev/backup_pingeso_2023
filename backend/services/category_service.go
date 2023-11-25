package services

import (
	"github.com/rez-dev/backup_pingeso_2023/backend/models"
	"github.com/rez-dev/backup_pingeso_2023/backend/repositories"
)

func GetCategories() (models.Categories, error) {
	categories, err := repositories.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategory(id string) (models.Categories, error) {
	categories, err := repositories.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func CreateCategory(newCategory models.Category) error {
	err := repositories.CreateCategory(newCategory)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(id string, newCategory models.Category) error {
	err := repositories.UpdateCategory(id, newCategory)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id string) error {
	err := repositories.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}

func GetParentsWordPress(id int) ([3]string, error) {
	parents, err := repositories.GetParentsWordPress(id)
	if err != nil {
		return [3]string{}, err
	}
	return parents, nil
}

func InsertCategories() (models.Categories, error) {
	categories, err := repositories.InsertCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoriesByTerm(id string) (models.Categories, error) {
	categories, err := repositories.GetCategoriesByTerm(id)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
