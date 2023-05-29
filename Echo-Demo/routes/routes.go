package routes

import (
	"Echo-Demo/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {

	userHandler := handlers.NewUserHandler(db)
	usersGroup := e.Group("/users")

	usersGroup.POST("/create_user", userHandler.CreateUser)
	usersGroup.GET("/all_users", userHandler.GetAllUsers)

}
