package auth

import (
	"time"

	"github.com/dilyara4949/drevmass/api/controller/auth"
	"github.com/dilyara4949/drevmass/internal/repository"
	"github.com/dilyara4949/drevmass/internal/usecase"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewLoginRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := &auth.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}