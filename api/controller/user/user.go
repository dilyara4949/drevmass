package user

import (
	// "fmt"
	"net/http"
	"time"

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
	// fmt.Println()
	userID := c.GetUint("x-user-id")
	// fmt.Println(userID)

	user, err := uc.UserUsecase.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	userClient := userToBasic(*user)
	c.JSON(http.StatusOK, userClient)
	
}


func (uc *UserController) UserUpdate(c *gin.Context) {
	userID := c.GetUint("x-user-id")
	// fmt.Println(userID)
	user, err := uc.UserUsecase.GetUserByID(c, userID)


	var request domain.UserInfo
	

	err = c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "3"})
		return
	}
	t := time.Time{}
	if request.Name != "" {
		user.Name = request.Name
	}
	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Gender != 0 {
		user.Gender = request.Gender
	}
	if request.Height != 0 {
		user.Height = request.Height
	}
	if request.Weight != 0 {
		user.Weight = request.Weight
	}
	if request.Birth != t {
		user.Birth = request.Birth
	}
	if request.Activity != 0 {
		user.Activity = request.Activity
	}
	

	_, err = uc.UserUsecase.Update(c, *user)
	userInfo := userInfoToClient(*user)
	c.JSON(http.StatusOK, userInfo)
}



func (uc *UserController) UserDelete(c *gin.Context) {
	userID := c.GetUint("x-user-id")

	err := uc.UserUsecase.Delete(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "User deleted"})
}
