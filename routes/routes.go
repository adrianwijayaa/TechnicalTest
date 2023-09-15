package routes

import "github.com/labstack/echo"

func RouteInit(e *echo.Group) {
	ToDoListRoute(e)
}
