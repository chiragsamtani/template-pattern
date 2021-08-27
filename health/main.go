package main

import (
	"flag"
	"fmt"
	server "github.com/chiragsamtani/template-pattern/common/server"
	types "github.com/chiragsamtani/template-pattern/common/types"
	"github.com/chiragsamtani/template-pattern/shared/config"
	"github.com/chiragsamtani/template-pattern/shared/database"

	"github.com/chiragsamtani/template-pattern/health/repository"
	"github.com/chiragsamtani/template-pattern/health/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP Server")
	flag.Parse()

	// setup echo router
	e := echo.New()
	// use default logger as middleware
	e.Use(middleware.Logger())

	// setup mysql connection
	config, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatalf("Cannot read env file, %s", err)
		panic(err)
	}
	db, err := database.OpenMysqlConn(*config)
	db.AutoMigrate(types.Products{})
	if err != nil {
		log.Fatalf("Cannot initialize database connection, %s", err)
		panic(err)
	}

	// repository is a dependency of usecase/service
	mySqlHealthRepo := repository.NewMySQLHealthRepository(db)

	useCase := usecase.NewHealthUsecase(mySqlHealthRepo)

	// register all required product interfaces
	// registering custom handlers will only be supported
	// at a per-usecase/service level
	server.RegisterHandlersWithBaseURL(e, useCase, "/health")

	fmt.Println(e.Routes())

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *port)))
}
