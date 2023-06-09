package routes

import (
	"Echo-Demo/handlers"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	userHandler := handlers.NewUserHandler(db)
	loginHandler := handlers.LoginHandler(db)

	e.POST("/login", loginHandler.Login)

	usersGroup := e.Group("/users")
	usersGroup.POST("/create_user", userHandler.CreateUser)
	usersGroup.GET("/all_users", userHandler.GetAllUsers)
	usersGroup.GET("/getbyid", userHandler.GetUserById)
	usersGroup.PUT("/update", userHandler.UpdateUser)
	usersGroup.DELETE("/delete/:id", userHandler.DeleteUser)

}
