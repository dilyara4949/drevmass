package domain

import (
	"context"
)

type Day struct {
	ID     uint
	UserID uint
	Mon    bool
	Tue    bool
	Wed    bool
	Thu    bool
	Fri    bool
	Sat    bool
	Sun    bool
	Time   string
}

type DayRepository interface {
	Create(ctx context.Context, day Day) (Day, error)
	Get(ctx context.Context, userID uint) (Day, error)
	Update(ctx context.Context, day Day) (Day, error)
}

type DayUsecase interface {
	Create(c context.Context, day Day) (Day, error)
	Get(c context.Context, userID uint) (Day, error)
	Update(c context.Context, day Day) (Day, error)
}