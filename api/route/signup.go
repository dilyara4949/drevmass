package route

import (
	"time"

	"github.com/dilyara4949/drevmass/api/controller"
	"github.com/dilyara4949/drevmass/internal/postgres"
	"github.com/dilyara4949/drevmass/internal/repository"
	"github.com/dilyara4949/drevmass/internal/usecase"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *pkg.Env, timeout time.Duration, db postgres.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}