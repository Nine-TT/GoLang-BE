package models

type User struct {
	Id       uint   `json:"id" form:"id" gorm:"primaryKey" `
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Age      uint8  `json:"age" form:"age"`
}

type Login struct {
	email    string `json:"email"`
	password string `json:"password"`
}
