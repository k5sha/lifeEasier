package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/k5sha/lifeEasier/internal/model"
	"time"
)

type LinkPostgresStorage struct {
	db *sqlx.DB
}

func NewLinkStorage(db *sqlx.DB) *LinkPostgresStorage {
	return &LinkPostgresStorage{db: db}
}
func (l *LinkPostgresStorage) LinkById(ctx context.Context, id int64) (*model.Link, error) {
	conn, err := l.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var link dbLink
	if err := conn.GetContext(ctx, &link, `SELECT * FROM links WHERE id = $1`, id); err != nil {
		return nil, err
	}

	return (*model.Link)(&link), nil
}

func (l *LinkPostgresStorage) Add(ctx context.Context, link model.Link) (int64, error) {
	conn, err := l.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var id int64
	row := conn.QueryRowxContext(
		ctx,
		`INSERT INTO links (link, message, user_id) VALUES ($1, $2, $3) RETURNING id;`, link.Link, link.Message, link.UserId,
	)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

type dbLink struct {
	Id        int64     `db:"id"`
	Link      string    `db:"link"`
	Message   string    `db:"message"`
	Used      bool      `db:"used"`
	UserId    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}
