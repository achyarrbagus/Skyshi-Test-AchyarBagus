package routes

import (
	"todo-app/handlers"
	"todo-app/pkg/mysql"
	repositories "todo-app/repository"

	"github.com/labstack/echo/v4"
)

func ActivityRoutes(e *echo.Group) {
	activityRepository := repositories.RepositoryActivity(mysql.DB)
	h := handlers.HandlerActivity(activityRepository)

	e.POST("/activity-groups", h.CreateActivity)
	e.GET("/activity-groups/:id", h.GetActivity)
	e.DELETE("/activity-groups/:id", h.DeleteActivity)
	e.PATCH("/activity-groups/:id", h.UpdateActivity)
	e.GET("/activity-groups", h.FindActivity)

}
