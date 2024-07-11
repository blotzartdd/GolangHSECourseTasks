package main

import (
	"HSECourse/Homework2/accounts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	accountsHandler := accounts.New()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)

	e.POST("/account/create", accountsHandler.CreateAccount)
	e.POST("/account/delete", accountsHandler.DeleteAccount)
	e.POST("/account/changeName", accountsHandler.ChangeAccountName)
	e.POST("/account/changeAmount", accountsHandler.ChangeAccountAmount)

	e.Logger.Fatal(e.Start(":1323"))
}
