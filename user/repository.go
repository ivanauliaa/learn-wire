package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ivanauliaa/learn-wire/domain"
)

type repository struct {
	db *sql.DB
}

func (r *repository) FetchByUsername(ctx context.Context, username string) (*domain.UserEntity, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM users WHERE username = ?")
	if err != nil {
		return nil, fmt.Errorf("error while preparing sql statement. %s", err.Error())
	}
	defer stmt.Close()

	var user domain.UserEntity

	row := stmt.QueryRowContext(ctx, username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		return nil, fmt.Errorf("error while retrieving user. %s", err.Error())
	}

	return &user, nil
}
