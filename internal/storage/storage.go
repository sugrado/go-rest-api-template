package storage

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/sugrado/go-rest-api-template/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sugrado/go-rest-api-template/internal/config"
	"github.com/sugrado/go-rest-api-template/internal/storage/user"
)

type Storage struct {
	pool *pgxpool.Pool
}

type Database struct {
	db    *pgxpool.Pool
	users UserRepository
}

type Store interface {
	Users() UserRepository
}

type queryTracer struct {
	log *slog.Logger
}

func (t *queryTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	t.log.Debug("Executing command -> ", "sql", data.SQL, "args", data.Args)
	return ctx
}

func (t *queryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
}

func DBConn(cfg config.DatabaseConfiguration) *Storage {
	dbUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)
	c, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse config failed: %v\n", err)
		os.Exit(1)
	}

	c.MinConns = 5
	c.MaxConns = 10

	c.ConnConfig.Tracer = &queryTracer{
		log: logger.Logger(),
	}

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, c)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed connect to database: %v\n", err)
		os.Exit(1)
	}

	if err = CheckDBHealth(ctx, pool); err != nil {
		fmt.Fprintf(os.Stderr, "Failed connect to database: %v\n", err)
		os.Exit(1)
	}

	return &Storage{
		pool: pool,
	}
}

func CheckDBHealth(ctx context.Context, p *pgxpool.Pool) error {
	return p.Ping(ctx)
}

func RegisterRepos(db *Storage) *Database {
	return &Database{
		db:    db.pool,
		users: user.NewRepository(db.pool),
	}
}

func (s *Storage) Close() {
	s.pool.Close()
}

func (db *Database) Users() UserRepository {
	return db.users
}
