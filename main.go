package main

import (
	"fmt"
	"todo-app/database"
	"todo-app/pkg/mysql"
	"todo-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	mysql.DatabaseInit()
	database.AutoMigration()
	e := echo.New()

	routes.RouteInit(e.Group("/todolist.api.devcode.gethired.id"))
	fmt.Println("server running localhost:5000 ")
	e.Logger.Fatal(e.Start("localhost:5000")) // delete localhost
}
