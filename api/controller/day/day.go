package day

import (
	"net/http"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type DayController struct {
	DayUsecase domain.DayUsecase
}

// @Summary CreateDay
// @Security ApiKeyAuth
// @Tags Day
// @Description create Day
// @ID create-Day
// @Produce  json
// @Accept   json
// @Param	Day	body  domain.Day	true	"Create Day"
// @Success 200 {object} domain.Day
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /days [post]
func (dc *DayController) Create(c *gin.Context) {
	userID := c.GetUint("x-user-id")
	var request domain.Day
	request.UserID = userID

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	// fmt.Println(request.Time)
	
	day, err := dc.DayUsecase.Create(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, day)
}

// @Summary GetDay
// @Security ApiKeyAuth
// @Tags Day
// @Description get day
// @ID get-Day
// @Produce  json
// @Success 200 {object} domain.Day
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /days [get]
func (dc *DayController) Get(c *gin.Context) {
	userID := c.GetUint("x-user-id")
	// fmt.Println(userID)
	day, err := dc.DayUsecase.Get(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, day)
}

// @Summary UpdateDay
// @Security ApiKeyAuth
// @Tags Day
// @Description update day
// @ID update-day
// @Produce  json
// @Accept       json
//@Param	day	body	domain.Day	true	"update day"
// @Success 200 {object} domain.Day
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /days/update [post]
func (dc *DayController) Update(c *gin.Context) {
	userID := c.GetUint("x-user-id")
	var request domain.Day
	request.UserID = userID

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	request, err = dc.DayUsecase.Update(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, request)
}