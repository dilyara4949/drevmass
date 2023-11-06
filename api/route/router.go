package route

import (
	"time"

	"github.com/dilyara4949/drevmass/internal/postgres"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/dilyara4949/drevmass/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(env *pkg.Env, timeout time.Duration, db postgres.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
}