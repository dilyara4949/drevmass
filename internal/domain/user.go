package domain

import (
	"context"
	"time"
)

type User struct {
	ID       uint      `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Name     string    `json:"name"`	
	Gender   uint      `json:"gender"`
	Height   uint      `json:"height"`
	Weight   uint      `json:"weight"`
	Birth    time.Time `json:"birth"`
	Activity uint      `json:"activity"`
}

type UserInfo struct {
	ID       uint      `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Gender   uint      `json:"gender"`
	Height   uint      `json:"height"`
	Weight   uint      `json:"weight"`
	Birth    time.Time `json:"birth"`
	Activity uint      `json:"activity"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, user User) (*User, error)
	Delete(ctx context.Context, id uint) error
}

type UserUsecase interface {
	GetUserByID(c context.Context, id uint) (*User, error)
	Update(c context.Context, user User) (*User, error)
	Delete(c context.Context, id uint) error
}
