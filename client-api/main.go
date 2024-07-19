package main

import (
	"log"

	"github.com/felipematthew/go-learnings/client-api/api/config"
	"github.com/felipematthew/go-learnings/client-api/api/routes"
	_ "github.com/felipematthew/go-learnings/client-api/docs"
	"github.com/labstack/echo/v4"
)

// @title Client API
// @version 1.0
// @description Essa é a API oficial do projeto Client API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @schemes https
// @Produces  application/json
// @Consumes  application/json

// @SecurityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @Scheme bearer
// @BearerFormat JWT
func main() {
	e := echo.New()
	routes.Generate(e)

	config.DatabaseInit()

	log.Fatal(e.Start(":8080"))
}
