package usecase

import (
	"context"
	"time"

	"github.com/dilyara4949/drevmass/internal/domain"
)



type favoritesUsecase struct {
	favoritesRepository domain.FavoritesRepository
	contextTimeout time.Duration
}

func NewfavoritesUsecase(favoritesRepository domain.FavoritesRepository, timeout time.Duration) domain.FavoritesUsecase {
	return &favoritesUsecase{
		favoritesRepository: favoritesRepository,
		contextTimeout: timeout,
	}
}

func (l *favoritesUsecase) Create(c context.Context, favorites domain.Favorites) (domain.Favorites, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.favoritesRepository.Create(c, favorites)
}


func (l *favoritesUsecase) Get(c context.Context, userid uint) ([]domain.Lesson, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.favoritesRepository.Get(c, userid)
}

func (l *favoritesUsecase) Delete(c context.Context, lessonid string) ( error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.favoritesRepository.Delete(c, lessonid)
}


