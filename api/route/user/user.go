package user

import (
	"time"

	"github.com/dilyara4949/drevmass/api/controller/user"
	"github.com/dilyara4949/drevmass/internal/repository"
	"github.com/dilyara4949/drevmass/internal/usecase"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewUserRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := &user.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
		// Env:          env,
	}
	group.GET("/user", lc.UserBasic)
	group.POST("/user", lc.UserUpdate)
	group.GET("/user/information", lc.UserInfo)
	group.POST("/user/delete", lc.UserDelete)
	group.POST("/forget-password", lc.ForgetPassword)
}