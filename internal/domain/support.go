package domain

import "context"

type Support struct {
	ID uint
	ProblemDescription string
	AnswerDescription string
	UserID uint
}

type SupportRepository interface {
	Create(ctx context.Context, s Support) (Support, error)
	Get(ctx context.Context, userid uint) (Support, error) 
	GetAll(ctx context.Context ) ([]Support, error) 
	Update(ctx context.Context, ans string, userid string) (Support, error) 
}

type SupportUsecase interface {
	Create(ctx context.Context, s Support) (Support, error)
	Get(ctx context.Context, userid uint) (Support, error) 
	Answer(ctx context.Context, ans string, userid string) (Support, error) 
	GetAll(ctx context.Context ) ([]Support, error)
	
}