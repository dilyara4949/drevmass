package favorites

import (
	"net/http"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type FavoritesController struct {
	FavoritesUsecase domain.FavoritesUsecase
}

// @Summary GetFavorites
// @Security ApiKeyAuth
// @Tags favorites
// @Description get Favorites
// @ID get-favorites
// @Produce  json
// @Success 200 {object} []domain.Lesson
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /favorites [get]
func (l *FavoritesController) Get(c *gin.Context) {
	userID := c.GetUint("x-user-id")
	favoritess, err := l.FavoritesUsecase.Get(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, favoritess)
}

// @Summary Createfavorites
// @Security ApiKeyAuth
// @Tags favorites
// @Description create favorites
// @ID create-favorites
// @Produce  json
// @Accept   json
// @Param	favorites	body  domain.Favorites	true	"Create favorites"
// @Success 200 {object} domain.Favorites
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /favorites [post]
func (l *FavoritesController) Create(c *gin.Context) {
	userID := c.GetUint("x-user-id")
	var request domain.Favorites

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	request.UserID = userID
	favorites, err := l.FavoritesUsecase.Create(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, favorites)
}

// @Summary Delete lesson from Favorites
// @Security ApiKeyAuth
// @Tags favorites
// @Description Delete favorites
// @ID Delete-favorites
// @Produce  json
// @Accept   json
// @Param    lessonid   path      int  true  "lesson ID"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /favorites/{lessonid} [delete]
func (p *FavoritesController) Delete(c *gin.Context) {
	lessonID := c.Param("lessonid")

	err := p.FavoritesUsecase.Delete(c, lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "lesson deleted from favorites deleted"})
}

