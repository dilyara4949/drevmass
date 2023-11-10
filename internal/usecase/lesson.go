package usecase

import (
	"context"
	"time"

	"github.com/dilyara4949/drevmass/internal/domain"
)



type lessonUsecase struct {
	lessonRepository domain.LessonRepository
	contextTimeout time.Duration
}

func NewLessonUsecase(lessonRepository domain.LessonRepository, timeout time.Duration) domain.LessonUsecase {
	return &lessonUsecase{
		lessonRepository: lessonRepository,
		contextTimeout: timeout,
	}
}

func (l *lessonUsecase) Create(c context.Context, lesson domain.Lesson) (domain.Lesson, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.lessonRepository.Create(c, lesson)
}

func (l *lessonUsecase) GetByID(c context.Context, id string) (domain.Lesson, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.lessonRepository.GetByID(c, id)
}

func (l *lessonUsecase) GetAll(c context.Context) ([]domain.Lesson, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.lessonRepository.GetAll(c)
}

func (l *lessonUsecase) Update(c context.Context, lesson domain.Lesson) (domain.Lesson, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.lessonRepository.Update(c, lesson)
}

func (l *lessonUsecase) Delete(c context.Context, id string) ( error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.lessonRepository.Delete(c, id)
}


