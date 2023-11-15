package route

import (
	"time"

	"github.com/dilyara4949/drevmass/api/route/auth"
	"github.com/dilyara4949/drevmass/api/route/day"
	"github.com/dilyara4949/drevmass/api/route/favorites"
	"github.com/dilyara4949/drevmass/api/route/lesson"
	"github.com/dilyara4949/drevmass/api/route/product"
	"github.com/dilyara4949/drevmass/api/route/support"
	"github.com/dilyara4949/drevmass/api/route/user"
	_ "github.com/dilyara4949/drevmass/docs"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/dilyara4949/drevmass/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)


func Setup(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, gin *gin.Engine) {
	// eng := engine.Default()
	publicRouter := gin.Group("/api")

    // All Public APIs
	publicRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth.NewSignupRouter(env, timeout, db, publicRouter)
	auth.NewLoginRouter(env, timeout, db, publicRouter)

	
	protectedRouter := gin.Group("/api")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	
	
	// All Private APIs
	user.NewUserRouter(env, timeout, db, protectedRouter)
	day.NewDayRouter(env, timeout, db, protectedRouter)
	product.NewProductRouter(env, timeout, db, protectedRouter)
	lesson.NewLessonRouter(env, timeout, db, protectedRouter)
	favorites.NewFavoritesRouter(env, timeout, db, protectedRouter)
	support.NewSupportRouter(env, timeout, db, protectedRouter)
	auth.NewLogoutRouter(env, timeout, db, protectedRouter)

}

//  r := gin.Default()
//     eng := engine.Default()
//     cfg := config.Config{
//         Databases: config.DatabaseList{
//             "default": {
//                 Host:       "127.0.0.1",
//                 Port:       "5432",
//                 User:       "postgres",
//                 Pwd:        "12345",
//                 Name:       "drevmass",
//                 Driver:     config.DriverPostgresql,
//             },
//         },
//         UrlPrefix: "admin",
//         Store: config.Store{
//             Path:   "./uploads",
//             Prefix: "uploads",
//         },
//         Language: language.CN,
//     }

//     // AddGenerator can also be used to load the Generator, like:
//     // eng.AddGenerator("user", GetUserTable)

//     eng.AddConfig(&cfg).
//         AddGenerators(tables.Generators).  // 加载插件
//         Use(r)

//     r.Run(":9033")