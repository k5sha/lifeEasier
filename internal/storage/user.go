package storage

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/k5sha/lifeEasier/internal/model"
	"time"
)

type UserPostgresStorage struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *UserPostgresStorage {
	return &UserPostgresStorage{db: db}
}

func (u *UserPostgresStorage) UserById(ctx context.Context, id int64) (*model.User, error) {
	conn, err := u.db.Connx(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var user dbUser
	if err := conn.GetContext(ctx, &user, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return nil, err
	}

	return (*model.User)(&user), nil
}

func (u *UserPostgresStorage) Add(ctx context.Context, user model.User) (int64, error) {
	conn, err := u.db.Connx(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var id int64
	row := conn.QueryRowxContext(
		ctx,
		`INSERT INTO users (username, chat_id) VALUES ($1, $2) RETURNING id;`, user.Username, user.ChatId,
	)

	if err := row.Err(); err != nil {
		return 0, err
	}

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

type dbUser struct {
	Id        int64     `db:"id"`
	Username  string    `db:"username"`
	ChatId    int64     `db:"chat_id"`
	CreatedAt time.Time `db:"created_at"`
}
