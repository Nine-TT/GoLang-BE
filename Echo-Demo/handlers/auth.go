package handlers

import (
	"Echo-Demo/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func LoginHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type jwtCustomClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func (c *UserHandler) Login(ctx echo.Context) error {
	account := new(models.Login)
	if err := ctx.Bind(account); err != nil {
		return err
	}

	var user models.User

	result := c.db.Where("email = ?", account.Email).First(&user)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to query database",
		})
	} else if result.RowsAffected == 0 {
		return ctx.JSON(http.StatusNotFound, echo.Map{
			"message": "User not found Or Email incorrect",
		})
	}

	match := CheckPasswordHash(account.Password, user.Password)
	if match == false {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{
			"message": "Password incorrect!",
		})
	}

	claims := &jwtCustomClaims{
		user.Id,
		user.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	if match == true {
		return ctx.JSON(http.StatusOK, echo.Map{
			"message": "Login success",
			"token":   t,
		})
	}

	return nil
}
