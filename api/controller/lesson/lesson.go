package lesson
import (
	"net/http"
	"strconv"


	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type LessonController struct {
	LessonUsecase domain.LessonUsecase
}


func (l *LessonController) GetOne(c *gin.Context) {
	LessonID := c.Param("id")
	// fmt.Println(LessonID)
	Lesson, err := l.LessonUsecase.GetByID(c, LessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, Lesson)
}

func (l *LessonController) GetAll(c *gin.Context) {
	lessons, err := l.LessonUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, lessons)
}


func (l *LessonController) Create(c *gin.Context) {

	var request domain.Lesson

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	
	lesson, err := l.LessonUsecase.Create(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, lesson)
}


func (l *LessonController) Update(c *gin.Context) {
	lessonID := strToUint( c.Param("id"))
	

	var request domain.Lesson
	

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "3"})
		return
	}
	request.ID = lessonID

	
	
	_, err = l.LessonUsecase.Update(c, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, request)
}



func (p *LessonController) Delete(c *gin.Context) {
	LessonID := c.Param("id")

	err := p.LessonUsecase.Delete(c, LessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Lesson deleted"})
}

func strToUint(s string) uint {
    i, _ := strconv.Atoi(s)
    return uint(i)
}