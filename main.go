package main

import (
	"TechnicalTest/database"
	"TechnicalTest/pkg/postgresql"
	"TechnicalTest/routes"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	godotenv.Load()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
	}))

	postgresql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	// var PORT = os.Getenv("PORT")

	fmt.Println("server running on localhost:5000")
	e.Logger.Fatal(e.Start(":5000"))
}
