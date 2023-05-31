package handlers

import (
	"Echo-Demo/models"
	"fmt"
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

	if err := ctx.Bind(user); err != nil {
		return err
	}

	hashPass, _ := HashPassword(user.Password)
	user.Password = hashPass

	fmt.Println("============> ", user)

	if user.Name == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "missing name!",
		})
	} else if user.Age == 0 {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "missing name!",
		})
	} else {
		c.db.Where(models.User{Email: user.Email}).FirstOrCreate(&user)

		//c.db.Create(&user)
		return ctx.JSON(http.StatusCreated, map[string]any{
			"message": "create success",
			"user":    user,
		})
	}

	return nil
}

func (c *UserHandler) GetAllUsers(ctx echo.Context) error {
	var users []models.User
	c.db.Omit("password").Find(&users)
	return ctx.JSON(http.StatusOK, users)
}

func (c *UserHandler) GetUserById(ctx echo.Context) error {
	id := ctx.FormValue("id")

	var user models.User

	if err := c.db.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found!")
	}
	return ctx.JSON(http.StatusOK, user)
}

func (c *UserHandler) UpdateUser(ctx echo.Context) error {
	id := ctx.FormValue("id")

	var user models.User

	if err := c.db.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found!")
	}

	if err := ctx.Bind(&user); err != nil {
		return err
	}

	c.db.Omit("password").Save(&user)

	return ctx.JSON(http.StatusOK, map[string]any{
		"message": "Update user success",
	})
}

func (c *UserHandler) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")

	var user models.User
	if err := c.db.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{
			"message": "User not found",
		})
	}

	c.db.Delete(&user)
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Delete user success!",
	})

	return nil
}
