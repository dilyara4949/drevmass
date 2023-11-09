package route

import (
	"time"

	"github.com/dilyara4949/drevmass/api/route/user"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/dilyara4949/drevmass/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Setup(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	user.NewUserRouter(env, timeout, db, protectedRouter)
}