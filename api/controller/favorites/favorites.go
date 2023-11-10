package favorites

import (
	// "fmt"
	"net/http"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type FavoritesController struct {
	FavoritesUsecase domain.FavoritesUsecase
}



func (l *FavoritesController) Get(c *gin.Context) {
	userID := c.GetUint("x-user-id")
	// fmt.Println(userID)
	favoritess, err := l.FavoritesUsecase.Get(c, userID)
	if err != nil {
		// fmt.Println(userID)
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, favoritess)
}


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






func (p *FavoritesController) Delete(c *gin.Context) {
	lessonID := c.Param("lessonid")

	err := p.FavoritesUsecase.Delete(c, lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "lesson deleted from favorites deleted"})
}

