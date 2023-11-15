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

// @Summary GetLesson
// @Security ApiKeyAuth
// @Tags Lesson
// @Description get Lesson by id
// @ID get-Lesson
// @Produce  json
// @Param        id   path      int  true  "Lesson ID"
// @Success 200 {object} domain.Lesson
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /lessons/{id} [get]
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

// @Summary GetLessons
// @Security ApiKeyAuth
// @Tags Lesson
// @Description get all Lessons
// @ID get-Lessons
// @Produce  json
// @Success 200 {object} []domain.Lesson
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /lessons [get]
func (l *LessonController) GetAll(c *gin.Context) {
	lessons, err := l.LessonUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, lessons)
}

// @Summary CreateLesson
// @Security ApiKeyAuth
// @Tags Lesson
// @Description create Lesson
// @ID create-Lesson
// @Produce  json
// @Accept       json
//
// @Param	Lesson	body  domain.Lesson	true	"Create Lesson"
//
// @Success 200 {object} domain.Lesson
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /lessons [post]
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

// @Summary UpdateLesson
// @Security ApiKeyAuth
// @Tags Lesson
// @Description update Lesson
// @ID update-Lesson
// @Produce  json
// @Accept       json
// @Param        id   path      int  true  "Lesson ID"
//@Param	Lesson	body	domain.Lesson	true	"update Lesson"
// @Success 200 {object} domain.Lesson
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /lessons/{id} [post]
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



// @Summary DeleteLesson
// @Security ApiKeyAuth
// @Tags Lesson
// @Description Delete Lesson
// @ID Delete-Lesson
// @Produce  json
// @Accept       json
// @Param    id   path      int  true  "Lesson ID"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /lessons/{id} [delete]
func (p *LessonController) Delete(c *gin.Context) {
	LessonID := c.Param("id")

	err := p.LessonUsecase.Delete(c, LessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Lesson deleted"})
}

// @Summary ChangeOrderOfLessons
// @Security ApiKeyAuth
// @Tags Lesson
// @Description change order of lessons
// @ID change-order-lesson
// @Produce  json
// @Accept       json
// @Param    a   path      int  true  "First lesson"
// @Param    b   path      int  true  "Second lesson"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /lessons/change/{a}/{b} [post]
func (l *LessonController) ChangeOrder(c *gin.Context) {
	a := strToUint(c.Param("a"))
	b := strToUint(c.Param("b"))
	err := l.LessonUsecase.ChangeOrder(c, a, b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Order changed"})
}


func strToUint(s string) uint {
    i, _ := strconv.Atoi(s)
    return uint(i)
}