package models

type Service struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	State       string `json:"state"`
	Id_user     int    `json:"id_user"`
	Id_wp_term  int    `json:"id_wp_term"`
}

type Services []Service
