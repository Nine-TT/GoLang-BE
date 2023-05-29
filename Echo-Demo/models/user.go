package models

import "gorm.io/gorm"

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	gorm.Model
}