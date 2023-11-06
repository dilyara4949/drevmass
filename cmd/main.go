package main

import (
	"time"

	"github.com/dilyara4949/drevmass/api/route"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	app := pkg.App()
	env := app.Env
	db := app.Pql
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}