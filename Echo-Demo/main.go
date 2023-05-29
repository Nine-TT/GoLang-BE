package main

import (
	DB "Echo-Demo/db"
	"Echo-Demo/models"
	"Echo-Demo/routes"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	server := echo.New()

	db, err := DB.ConnectDB()

	if err != nil {
		fmt.Println("Error connect db: ", err)
		return
	} else {
		fmt.Println("connect db success!")
	}

	db.AutoMigrate(&models.User{})

	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "SERVER ON")
	})

	routes.InitRoutes(server, db)

	server.Logger.Fatal(server.Start(":5000"))
}
