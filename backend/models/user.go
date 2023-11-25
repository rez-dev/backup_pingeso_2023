package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//ID       int    `json:"id"`
	Name     string `json:"name"`
	Rut      string `json:"rut"`
	Email    string `gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
	//Hash      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Users []User
