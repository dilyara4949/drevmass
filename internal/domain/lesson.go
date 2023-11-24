package domain

import (
	"context"
	// "time"
)

type Lesson struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageSrc    string  `json:"image_src"`
	VideoSrc    string  `json:"video_src"`
	Duration    float64 `json:"duration"`
	Created_at  string  `json:"created_at"`
	Updated_at  string  `json:"updated_at"`
}

type LessonRepository interface {
	Create(ctx context.Context, lesson Lesson) (Lesson, error)
	GetByID(ctx context.Context, id string) (Lesson, error)
	GetAll(ctx context.Context) ([]Lesson, error)
	Update(ctx context.Context, lesson Lesson) (Lesson, error)
	Delete(ctx context.Context, id string) error
	ChangeOrder(ctx context.Context, a uint, b uint) error
}

type LessonUsecase interface {
	Create(c context.Context, lesson Lesson) (Lesson, error)
	GetByID(c context.Context, id string) (Lesson, error)
	Update(c context.Context, lesson Lesson) (Lesson, error)
	Delete(c context.Context, id string) error
	GetAll(c context.Context) ([]Lesson, error)
	ChangeOrder(ctx context.Context, a uint, b uint) error
}
