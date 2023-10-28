package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sugrado/go-rest-api-template/internal/app/users"
	"github.com/sugrado/go-rest-api-template/pkg/logger"
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
		logger.Logger().Error(err.Error())
		return id, err
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
		logger.Logger().Error(err.Error())
		return nil, err
	}

	row := s.pool.QueryRow(ctx, rawSql, args...)

	var user users.User
	if err := row.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
	); err != nil {
		logger.Logger().Error(err.Error())
		return nil, err
	}

	return &user, nil
}
