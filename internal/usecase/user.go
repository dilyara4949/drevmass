package usecase

import (
	"context"
	"time"

	"github.com/dilyara4949/drevmass/internal/domain"
)



type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) GetUserByID(c context.Context, id uint) (*domain.User, error) {
	_, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetByID(c, id)
}

func (uu *userUsecase) Update(c context.Context, user domain.User) (*domain.User, error) {
	_, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.Update(c, user)
}

func (uu *userUsecase) Delete(c context.Context, id uint) ( error) {
	_, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.Delete(c, id)
}


