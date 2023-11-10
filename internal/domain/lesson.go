package domain

import (
	"context"
	// "time"
)

type Lesson struct {
	ID uint
	Name string
	Title string
	Description string
	ImageSrc string
	VideoSrc string
	Duration float64
	Created_at string
	Updated_at string
}

type LessonRepository interface {
	Create(ctx context.Context, lesson Lesson) (Lesson, error)
	GetByID(ctx context.Context, id string) (Lesson, error) 
	GetAll(ctx context.Context, ) ([]Lesson, error) 
	Update(ctx context.Context,  lesson Lesson) (Lesson, error) 
	Delete(ctx context.Context, id string) ( error)
}

type LessonUsecase interface {
	Create(c context.Context, lesson Lesson) (Lesson, error)
	GetByID(c context.Context, id string) (Lesson, error)
	Update(c context.Context, lesson Lesson) (Lesson, error)
	Delete(c context.Context, id string) ( error)
	GetAll(c context.Context) ([]Lesson, error)
}