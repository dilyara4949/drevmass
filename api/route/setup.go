package route

import (
	"time"

	"github.com/dilyara4949/drevmass/api/route/day"
	"github.com/dilyara4949/drevmass/api/route/support"
	"github.com/dilyara4949/drevmass/api/route/favorites"
	"github.com/dilyara4949/drevmass/api/route/lesson"
	"github.com/dilyara4949/drevmass/api/route/product"
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
	day.NewDayRouter(env, timeout, db, protectedRouter)
	product.NewProductRouter(env, timeout, db, protectedRouter)
	lesson.NewLessonRouter(env, timeout, db, protectedRouter)
	favorites.NewFavoritesRouter(env, timeout, db, protectedRouter)
	support.NewSupportRouter(env, timeout, db, protectedRouter)
}