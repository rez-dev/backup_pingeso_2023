package models

type Information struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Position     int    `json:"position"`
	Author       string `json:"author"`
	Id_wp_post   int    `json:"id_wp_post"`
	Id_post_meta int    `json:"id_post_meta"`
	Id_author    int    `json:"id_author"`
	Id_wp_term   int    `json:"id_wp_term"`
}

type Informations []Information
