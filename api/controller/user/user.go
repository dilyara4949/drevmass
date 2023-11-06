package user 

import (
	"net/http"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.User
}

func (uc *UserController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	user, err := uc.UserUsecase.
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, nil)
}