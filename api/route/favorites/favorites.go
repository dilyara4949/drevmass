package favorites

import (
	"time"

	"github.com/dilyara4949/drevmass/api/controller/favorites"
	"github.com/dilyara4949/drevmass/internal/repository"
	"github.com/dilyara4949/drevmass/internal/usecase"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewFavoritesRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	r := repository.NewfavoritesRepository(db)
	c := &favorites.FavoritesController{
		FavoritesUsecase: usecase.NewfavoritesUsecase(r, timeout),
	}
	group.POST("/favorites", c.Create)
	group.GET("/favorites", c.Get)
	group.DELETE("/favorites/:lessonid", c.Delete)

}
