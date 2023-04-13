package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	ActivityRoutes(e)
	TodoRoutes(e)
}
