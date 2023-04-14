package main

import (
	"fmt"
	"todo-app/database"
	"todo-app/pkg/mysql"
	"todo-app/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	mysql.DatabaseInit()
	database.AutoMigration()
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type"},
	}))

	routes.RouteInit(e.Group("/todolist.api.devcode.gethired.id"))
	fmt.Println("server running localhost:3030 ")
	e.Logger.Fatal(e.Start("localhost:3030")) // delete localhost
}
