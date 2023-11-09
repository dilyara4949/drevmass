package domain

import (
	"context"
	"time"

	
)



type User struct {
	ID        uint   
	Email     string 
	Password  string 
	Name      string 
	Gender    uint 
	Height    uint 
	Weight    uint 
	Birth     time.Time
	Activity  uint 
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error) 
	GetByID(ctx context.Context, id uint) (*User, error) 
}

type UserUsecase interface {
	GetUserByID(c context.Context, id uint) (*User, error)
}