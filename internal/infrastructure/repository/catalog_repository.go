package repository

import (
	"context"
	"database/sql"
	"fmt"
	"food-ordering/internal/domain"
	"food-ordering/internal/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) FetchAllMenu(ctx context.Context) ([]domain.Menu, error) {
	query := "SELECT id, name, description, price, currency FROM menus"
	rows, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch menus: %w", err)
	}
	defer rows.Close()

	var menus []domain.Menu
	for rows.Next() {
		var menu domain.Menu
		if err := rows.Scan(&menu.ID, &menu.Name, &menu.Description, &menu.Price, &menu.Currency); err != nil {
			return nil, fmt.Errorf("failed to scan menu: %w", err)
		}

		menus = append(menus, menu)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return menus, nil
}

func NewCatalogRepository(db *sql.DB) repository.CatalogRepository {
	return &userRepository{db: db}
}
