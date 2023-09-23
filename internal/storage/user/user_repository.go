package user

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sugrado/tama-server/internal/app/users"
	"time"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{pool: db}
}

func (s *Repository) Save(firstName, lastName, email string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	q := `INSERT INTO users (first_name, last_name, email) VALUES ($1::varchar, $2::varchar, $3::varchar) RETURNING id`

	var id int
	row := s.pool.QueryRow(ctx, q, firstName, lastName, email)
	err := row.Scan(&id)
	if err != nil {
		return id, fmt.Errorf("could not query users: %w", err)
	}

	return id, nil
}

func (s *Repository) Find(id int) (*users.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rawSql, args, err := psql.Select(
		"first_name",
		"last_name",
		"email").
		From("users").
		Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %w", err)
	}

	row := s.pool.QueryRow(ctx, rawSql, args...)

	var user users.User
	if err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	); err != nil {
		return nil, fmt.Errorf("could not query prepare query: %w", err)
	}

	return &user, nil
}
