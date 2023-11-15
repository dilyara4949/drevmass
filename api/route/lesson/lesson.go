package lesson

import (
	"time"

	"github.com/dilyara4949/drevmass/api/controller/lesson"
	"github.com/dilyara4949/drevmass/internal/repository"
	"github.com/dilyara4949/drevmass/internal/usecase"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewLessonRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	r := repository.NewLessonRepository(db)
	c := &lesson.LessonController{
		LessonUsecase: usecase.NewLessonUsecase(r, timeout),
	}
	group.POST("/lessons", c.Create)
	group.GET("/lessons", c.GetAll)
	group.GET("/lessons/:id", c.GetOne)
	group.POST("/lessons/:id", c.Update)
	group.DELETE("/lessons/:id", c.Delete)
	group.POST("/lessons/change/:a/:b", c.ChangeOrder)
}
