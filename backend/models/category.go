package models

type Category struct {
	ID         int    `json:"id"`
	Level1     string `json:"level1"`
	Level2     string `json:"level2"`
	Level3     string `json:"level3"`
	Id_wp_term int    `json:"id_wp_term"`
}

type Categories []Category
