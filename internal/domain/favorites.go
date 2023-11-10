package domain

import "context"

type Favorites struct {
	ID       uint
	UserID   uint
	LessonID uint
}


type FavoritesRepository interface {
	Create(ctx context.Context, f Favorites) (Favorites, error)
	Get(ctx context.Context, userid uint) ([]Lesson, error) 
	Delete(ctx context.Context, lessonid string) ( error)
}

type FavoritesUsecase interface {
	Create(ctx context.Context, f Favorites) (Favorites, error)
	Get(ctx context.Context, userid uint) ([]Lesson, error) 
	Delete(ctx context.Context, lessonid string) ( error)
}