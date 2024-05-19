package main

import (
	"echo-mysql-default/internal/routes"
	"echo-mysql-default/internal/util"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"os"
)

func main() {
	config := util.GetEnvironmentConfig()
	log.Printf("Environment Variables => %+v\n", config)

	db, err := util.OpenDatabase(config)

	if err != nil {
		log.Fatalln(util.GetStackTrace(err))
	}

	util.InitDB(db, "./migrations")

	logFile, _ := os.OpenFile(config.LogFile, os.O_SYNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logFile,
		Format: middleware.DefaultLoggerConfig.Format,
	}))

	database := goqu.New("mysql", db)

	routes.ProductsRoutes(e, database)

	defer e.Logger.Fatal(e.Start(config.ServerAddr))
	defer logFile.Close()
}
