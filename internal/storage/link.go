package storage

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/k5sha/lifeEasier/internal/model"
	"github.com/samber/lo"
	"time"
)

type LinkPostgresStorage struct {
	db *sqlx.DB
}

func NewLinkStorage(db *sqlx.DB) *LinkPostgresStorage {
	return &LinkPostgresStorage{db: db}
}

func (l *LinkPostgresStorage) AllNotPosted(ctx context.Context, limit uint64) ([]model.Link, error) {
	conn, err := l.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var links []dbLink

	if err := conn.SelectContext(
		ctx,
		&links,
		`SELECT DISTINCT ON (chat_id) id, link, message, chat_id, posted_at
         FROM links
         WHERE posted_at IS NULL
         AND scheduled_at <= CURRENT_TIMESTAMP 
         ORDER BY chat_id LIMIT $1;`,
		limit,
	); err != nil {
		return nil, err
	}

	return lo.Map(links, func(link dbLink, _ int) model.Link {
		return model.Link{
			Id:        link.Id,
			Link:      link.Link,
			Message:   link.Message,
			ChatId:    link.ChatId,
			CreatedAt: link.CreatedAt,
		}
	}), nil
}

func (l *LinkPostgresStorage) MarkAsPosted(ctx context.Context, id int64) error {
	conn, err := l.db.Connx(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(
		ctx,
		`UPDATE links SET posted_at = $1::timestamp WHERE id = $2;`,
		time.Now().UTC().Format(time.RFC3339),
		id,
	); err != nil {
		return err
	}

	return nil
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
		`INSERT INTO links (link, message, chat_id, scheduled_at) VALUES ($1, $2, $3, $4) RETURNING id;`,
		link.Link,
		link.Message,
		link.ChatId,
		link.ScheduledAt.UTC().Format(time.RFC3339),
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
	Id          int64        `db:"id"`
	Link        string       `db:"link"`
	Message     string       `db:"message"`
	ChatId      int64        `db:"chat_id"`
	PostedAt    sql.NullTime `db:"posted_at"`
	ScheduledAt time.Time    `db:"scheduled_at"`
	CreatedAt   time.Time    `db:"created_at"`
}
