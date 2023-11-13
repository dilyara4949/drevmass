package auth

import (
	"time"

	"github.com/dilyara4949/drevmass/api/controller/auth"
	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewLogoutRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {

	lc := &auth.LogoutController{
		Response: domain.SuccessResponse{Message: ""},
	}
	group.GET("/logout", lc.Logout)
}