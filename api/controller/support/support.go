package support

import (
	// "fmt"
	"net/http"
	"strconv"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type SupportController struct {
	SupportUsecase domain.SupportUsecase
}



func (l *SupportController) Get(c *gin.Context) {
	userID := strToUint( c.Param("userid"))
	// fmt.Println(userID)
	Supports, err := l.SupportUsecase.Get(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Supports)
}

func (l *SupportController) GetAll(c *gin.Context) {
	
	Supports, err := l.SupportUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Supports)
}

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






func (p *SupportController) Answer(c *gin.Context) {
	userID := c.Param("userid")
	// support 
	request := domain.Support{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "3"})
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