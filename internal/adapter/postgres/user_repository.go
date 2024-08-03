package postgres

import (
	"context"
	"database/sql"

	"github.com/marcelofabianov/my-cash/internal/port/outbound"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, input outbound.CreateUserRepositoryInput) error {
	query := `
		INSERT INTO users (id, document, name, email, password, enabled, created_at, updated_at, deleted_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	u := input.User

	result, err := r.db.ExecContext(ctx, query,
		u.ID.String(),
		u.Document.String(),
		u.Name,
		u.Email.String(),
		u.Password.String(),
		u.Enabled.Bool(),
		u.CreatedAt.String(),
		u.UpdatedAt.String(),
		nil,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return err
	}

	return nil
}

func (r *UserRepository) UserExists(ctx context.Context, input outbound.UserExistsRepositoryInput) (bool, error) {
	query := `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 OR document = $2)
	`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, input.Email, input.Document).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
