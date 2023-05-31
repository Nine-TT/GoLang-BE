package handlers

import (
	"Echo-Demo/models"
	"github.com/labstack/echo/v4"
)

func (c *UserHandler) Login(ctx echo.Context) {
	account := new(models.Login)

	ctx.Bind(account)

}
