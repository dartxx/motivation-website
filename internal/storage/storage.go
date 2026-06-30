package storage

import (
	"context"
	"fmt"
	"go-fullstack/internal/models/card"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, connStr string) (*Storage, error) {
	const op = "storage.New"

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to create pool: %w", op, err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("%s: failed to ping database: %w", op, err)
	}

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS cards(
			id SERIAL PRIMARY KEY,
			title VARCHAR(50) NOT NULL,
			content TEXT NOT NULL,
			author VARCHAR(30) NOT NULL,
			created_at TIMESTAMP DEFAULT now()
		);
	`)

	if err != nil {
		pool.Close()
		return nil, fmt.Errorf("%s: failed to create table: %w", op, err)
	}

	return &Storage{pool}, nil
}

func (s *Storage) Close() {
	s.pool.Close()
}

func (storage *Storage) CreateCard(ctx context.Context, card *card.Card) error {
	const op = "storage.CreateCard"

	row := storage.pool.QueryRow(ctx,
		"INSERT INTO cards (title, content, author) VALUES ($1, $2, $3) RETURNING id, created_at", card.Title, card.Content, card.Author)

	err := row.Scan(&card.ID, &card.CreatedAt)
	if err != nil {
		return fmt.Errorf("%s: failed to create card: %w", op, err)
	}

	return nil
}

func (storage *Storage) GetAllCards(ctx context.Context) ([]card.Card, error) {
	const op = "storage.GetAllCards"

	rows, err := storage.pool.Query(ctx, "SELECT id, title, content, author, created_at FROM cards ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("%s: failed get cards: %w", op, err)
	}
	defer rows.Close()

	var cards []card.Card
	for rows.Next() {
		var c card.Card
		if err := rows.Scan(&c.ID, &c.Title, &c.Content, &c.Author, &c.CreatedAt); err != nil {
			return nil, fmt.Errorf("%s: failed get all fields: %w", op, err)
		}
		cards = append(cards, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows iteration error: %w", op, err)
	}

	return cards, nil
}
