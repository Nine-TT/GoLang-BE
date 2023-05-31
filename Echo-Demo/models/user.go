package models

import "time"

type User struct {
	Id       uint   `json:"id" form:"id" gorm:"primaryKey" `
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Age      uint8  `json:"age" form:"age"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Work struct {
	Id            int    `json:"id"`
	UserId        int    `json:"user_id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	DateCreated   *time.Time
	DateComplated *time.Time
	IsComplated   bool `json:"is_complated"`
	IsWarming     bool `json:"is_warming"`
	IsPending     bool `json:"is_pending"`
}
