package auth

import (
	"net/http"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)


type LogoutController struct {
	Response domain.SuccessResponse
}

func (lc *LogoutController) Logout(c *gin.Context) {
	
	cookie := &http.Cookie{
        Name:   "token",
        MaxAge: -1,
    }
    http.SetCookie(c.Writer, cookie)
	lc.Response.Message = "Logged out"

	c.JSON(http.StatusOK, lc.Response)
}