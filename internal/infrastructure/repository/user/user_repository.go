package user

import (
	"context"
	"database/sql"
	"fmt"
	"food-ordering/internal/domain"
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

func (u *userRepository) OrderMenu(ctx context.Context, order domain.Order) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *sql.DB) User {
	return &userRepository{db: db}
}
