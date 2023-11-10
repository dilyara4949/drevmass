package usecase

import (
	"context"
	"time"

	"github.com/dilyara4949/drevmass/internal/domain"
)



type supportUsecase struct {
	supportRepository domain.SupportRepository
	contextTimeout time.Duration
}

func NewsupportUsecase(supportRepository domain.SupportRepository, timeout time.Duration) domain.SupportUsecase {
	return &supportUsecase{
		supportRepository: supportRepository,
		contextTimeout: timeout,
	}
}

func (l *supportUsecase) Create(c context.Context, support domain.Support) (domain.Support, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.supportRepository.Create(c, support)
}


func (l *supportUsecase) Get(c context.Context, userid uint) (domain.Support, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.supportRepository.Get(c, userid)
}
func (l *supportUsecase) GetAll(c context.Context) ([]domain.Support, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.supportRepository.GetAll(c)
}

func (l *supportUsecase) Answer(c context.Context, ans string, userid string) ( domain.Support, error) {
	_, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.supportRepository.Update(c, ans, userid)
}
