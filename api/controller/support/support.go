package support

import (
	"net/http"
	"strconv"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type SupportController struct {
	SupportUsecase domain.SupportUsecase
}


// @Summary GetSupport
// @Security ApiKeyAuth
// @Tags support
// @Description get Support by userid
// @ID get-support
// @Produce  json
// @Param        userid   path      int  true  "user ID"
// @Success 200 {object} domain.Support
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /Supports/{userid} [get]
func (l *SupportController) Get(c *gin.Context) {
	userID := strToUint(c.Param("userid"))
	// fmt.Println(userID)
	Supports, err := l.SupportUsecase.Get(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Supports)
}

// @Summary GetSupports
// @Security ApiKeyAuth
// @Tags support
// @Description get all Supports
// @ID get-dupports
// @Produce  json
// @Success 200 {object} []domain.Support
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /supports [get]
func (l *SupportController) GetAll(c *gin.Context) {
	
	Supports, err := l.SupportUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Supports)
}


// @Summary CreateSupport
// @Security ApiKeyAuth
// @Tags support
// @Description create support
// @ID create-support
// @Produce  json
// @Accept   json
// @Param	Support	body  domain.Support	true	"Create Support"
// @Success 200 {object} domain.Support
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /supports [post]
func (l *SupportController) Create(c *gin.Context) {
	userID := c.GetUint("x-user-id")
	var request domain.Support

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	request.UserID = userID
	Support, err := l.SupportUsecase.Create(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, Support)
}





// @Summary AnswerToSupport
// @Security ApiKeyAuth
// @Tags support
// @Description answer support
// @ID answer-support
// @Produce  json
// @Accept   json
// @Param        userid   path      int  true  "user ID"
//@Param	Support	body	domain.Support	true	"answer Support"
// @Success 200 {object} domain.Support
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /supports/{userid} [post]
func (p *SupportController) Answer(c *gin.Context) {
	userID := c.Param("userid")
	request := domain.Support{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "3"})
		return
	}
	support, err := p.SupportUsecase.Answer(c, request.AnswerDescription, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, support)
}



func strToUint(s string) uint {
    i, _ := strconv.Atoi(s)
    return uint(i)
}