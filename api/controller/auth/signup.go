package auth

import (
	"net/http"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *pkg.Env
}

// @Summary SignUp
// @Tags auth
// @Description signup
// @ID signup
// @Accept  json
// @Produce  json
// @Param input body domain.Signup true "signup"
// @Success 200 {object} domain.SignupResponse{} 
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /signup [post]
func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.Signup
	

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "3"})
		return
	}
	// fmt.Println(request.Email, request.Name)
	if request.Email == "" || request.Password == "" || request.Name =="" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Required fields are missing or invalid."})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "9"})
		return
	}
	

	u, err := sc.SignupUsecase.GetUserByEmail(c, request.Email)
	// fmt.Print(u)
	// if (domain.User{} == u) {
	if (u != nil) { 
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User with given email already exists"})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "4"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "2"})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		// ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	_, err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: err.Error()})
		// c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "1"})
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "5"})
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "6"})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}