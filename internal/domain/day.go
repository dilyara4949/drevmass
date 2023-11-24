package domain

import (
	"context"
)

type Day struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Mon    bool   `json:"mon"`
	Tue    bool   `json:"tue"`
	Wed    bool   `json:"wed"`
	Thu    bool   `json:"thu"`
	Fri    bool   `json:"fri"`
	Sat    bool   `json:"sat"`
	Sun    bool   `json:"sun"`
	Time   string `json:"time"`
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
