package support

import (
	"time"

	"github.com/dilyara4949/drevmass/api/controller/support"
	"github.com/dilyara4949/drevmass/internal/repository"
	"github.com/dilyara4949/drevmass/internal/usecase"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewSupportRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	r := repository.NewsupportRepository(db)
	c := &support.SupportController{
		SupportUsecase: usecase.NewsupportUsecase(r, timeout),
	}
	group.GET("/supports", c.GetAll)
	group.GET("/supports/:userid", c.Get)
	group.POST("/supports", c.Create)
	group.POST("/supports/:userid", c.Answer)

}