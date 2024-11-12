package user

import (
	"context"
	"errors"
	"food-ordering/internal/domain"
	"food-ordering/internal/domain/repository"
	"food-ordering/internal/domain/usecase"
	"food-ordering/pkg/utils"
)

type userUseCaseImpl struct {
	userRepo repository.UserRepository
}

func (u *userUseCaseImpl) Register(ctx context.Context, user domain.User) (string, error) {
	if len(user.Password) < 6 {
		return "", errors.New("password too short")
	}

	userID, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return utils.GenerateJWT(userID)
}

func NewUserUseCase(repo repository.UserRepository) usecase.UserUseCase {
	return &userUseCaseImpl{userRepo: repo}

}
