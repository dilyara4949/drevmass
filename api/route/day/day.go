package day

import (
	"time"

	"github.com/dilyara4949/drevmass/api/controller/day"
	"github.com/dilyara4949/drevmass/internal/repository"
	"github.com/dilyara4949/drevmass/internal/usecase"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDayRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	dr := repository.NewDayRepository(db)
	dc := &day.DayController{
		DayUsecase: usecase.NewDayUsecase(dr, timeout),
		// Env:          env,
	}
	group.POST("/days", dc.Create)
	group.GET("/days", dc.Get)
	group.POST("/days/update", dc.Update)

}
