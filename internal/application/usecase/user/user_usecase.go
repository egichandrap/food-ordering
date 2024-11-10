package user

import (
	"context"
	"food-ordering/internal/domain"
	"food-ordering/internal/infrastructure/repository/user"
)

type useCase struct {
	repository user.User
}

func (u *useCase) GetListMenu(ctx context.Context) ([]domain.Menu, error) {
	menus, err := u.repository.FetchAllMenu(ctx)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func NewUC(repository user.User) UC {
	return &useCase{repository: repository}

}
