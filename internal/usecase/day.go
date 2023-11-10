package usecase

import (
	"context"
	"time"

	"github.com/dilyara4949/drevmass/internal/domain"
)


type dayUsecase struct {
	dayRepository domain.DayRepository
	contextTimeout time.Duration
}

func NewDayUsecase(dayRepository domain.DayRepository, timeout time.Duration) domain.DayUsecase {
	return &dayUsecase{
		dayRepository: dayRepository,
		contextTimeout: timeout,
	}
}

func (d *dayUsecase) Create(c context.Context, day domain.Day) (domain.Day, error) {
	_, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()
	return d.dayRepository.Create(c, day)
}

func (d *dayUsecase) Get(c context.Context, userID uint) (domain.Day, error) {
	_, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()
	return d.dayRepository.Get(c, userID)
}

func (d *dayUsecase) Update(c context.Context, day domain.Day) ( domain.Day, error) {
	_, cancel := context.WithTimeout(c, d.contextTimeout)
	defer cancel()
	return d.dayRepository.Update(c, day)
}
