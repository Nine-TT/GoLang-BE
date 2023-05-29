package handlers

import (
	"Echo-Demo/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (c *UserHandler) CreateUser(ctx echo.Context) error {

	user := new(models.User)
	hashPass, _ := HashPassword(user.Password)
	user.Password = hashPass
	if err := ctx.Bind(user); err != nil {
		return err
	}

	c.db.Create(&user)
	return ctx.JSON(http.StatusCreated, user)
}

func (c *UserHandler) GetAllUsers(ctx echo.Context) error {
	var users []models.User
	c.db.Find(&users)
	return ctx.JSON(http.StatusOK, users)
}
