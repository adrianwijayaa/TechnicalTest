package routes

import (
	"TechnicalTest/handlers"
	"TechnicalTest/pkg/postgresql"
	"TechnicalTest/repositories"

	"github.com/labstack/echo"
)

func ToDoListRoute(e *echo.Group) {
	toDoListRepository := repositories.RepositoryToDoLists(postgresql.DB)
	h := handlers.HandlerToDoLists(toDoListRepository)

	e.POST("/toDoList", h.CreateToDoList)
	e.GET("/findToDoList", h.FindToDoList)
	e.GET("/toDoList/:id", h.GetToDoListById)
	e.PATCH("/toDoList/:id", h.UpdateToDoListById)
	e.DELETE("/toDoList/:id", h.DeleteToDoList)
}
