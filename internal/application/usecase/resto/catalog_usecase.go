package resto

import (
	"context"
	"food-ordering/internal/domain"
	"food-ordering/internal/domain/repository"
	"food-ordering/internal/domain/usecase"
)

type catalogUseCase struct {
	repository repository.CatalogRepository
}

func (u *catalogUseCase) GetCatalogMenu(ctx context.Context) ([]domain.Menu, error) {
	menus, err := u.repository.FetchAllMenu(ctx)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func NewCatalogUseCase(repository repository.CatalogRepository) usecase.CatalogUseCase {
	return &catalogUseCase{repository: repository}

}
