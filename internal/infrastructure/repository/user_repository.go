package repository

import (
	"context"
	"errors"
	"food-ordering/internal/domain"
	"food-ordering/internal/domain/repository"
	"strconv"
)

type UserRepo struct {
	users map[string]domain.User
}

func (u *UserRepo) CreateUser(ctx context.Context, user domain.User) (string, error) {
	if _, exists := u.users[user.Username]; exists {
		return "", errors.New("username already exists")
	}

	user.ID = strconv.Itoa(len(u.users) + 1)
	u.users[user.Username] = user
	return user.ID, nil
}

func (u *UserRepo) GetUserByUsername(username string) (*domain.User, error) {
	user, exists := u.users[username]
	if exists {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func NewUserRepository() repository.UserRepository {
	return &UserRepo{users: make(map[string]domain.User)}
}
