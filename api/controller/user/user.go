package user

import (
	"fmt"
	"net/http"
	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}


func (uc *UserController) UserInfo(c *gin.Context) {
	userID := c.GetUint("x-user-id")
	// fmt.Println(userID)
	user, err := uc.UserUsecase.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	userClient := userInfoToClient(*user)
	c.JSON(http.StatusOK, userClient)
}

func (uc *UserController) UserBasic(c *gin.Context) {
	
	userID := c.GetUint("x-user-id")
	fmt.Println(userID)

	user, err := uc.UserUsecase.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	userClient := userToBasic(*user)
	c.JSON(http.StatusOK, userClient)
	
}


func (uc *UserController) UserUpdate(c *gin.Context) {
	
}
func (uc *UserController) UserDelete(c *gin.Context) {
	
}
func (uc *UserController) ForgetPassword(c *gin.Context) {
	
}

