package main

import (
	"immersive-dashboard/app/config"
	"immersive-dashboard/app/database"
	"immersive-dashboard/app/router"
	"immersive-dashboard/middlewares"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBMySql(*cfg)
	database.InitialMigration(db)

	e := echo.New()
	middlewares.LogMiddlewares(e)
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
