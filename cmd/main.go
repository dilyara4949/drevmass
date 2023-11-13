package main

import (
	"log"
	"time"
	_ "github.com/dilyara4949/drevmass/docs"
	"github.com/dilyara4949/drevmass/api/route"
	// "github.com/dilyara4949/drevmass/api/controller/route"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8081
//	@BasePath	/api
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app, err := pkg.App()

	if (err != nil) {
		log.Fatal(err)
	}

	env := app.Env
	db := app.Pql
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
// defer Close(app.Pql)