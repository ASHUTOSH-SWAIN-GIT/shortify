package db

import (
	"context"
	"database/sql"
	"errors"
	"shortify/internals/db/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrNotFound = errors.New("short code not found")
)

type URLStore interface {
	Save(ctx context.Context, url *models.URL) error
	Get(ctx context.Context, shortCode string) (*models.URL, error)
	Exists(ctx context.Context, shortCode string) (bool, error)
}

type MySqlStore struct {
	DB *sql.DB
}

func NewMySqlStore(dataSourceName string) (*MySqlStore, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &MySqlStore{DB: db}, nil
}

// inserts a url to the database
func (s *MySqlStore) Save(ctx context.Context, url *models.URL) error {
	query := "INSERT INTO urls (short_code, long_url) VALUES (?, ?)"
	_, err := s.DB.ExecContext(ctx, query, url.ShortCode, url.LongURL)
	return err
}

// retreives a url from the databaes using the shortcode
func (s *MySqlStore) Get(ctx context.Context, shortCode string) (*models.URL, error) {
	url := &models.URL{}
	query := "SELECT short_code , long_url , created_at FROM urls WHERE short_code = ?"

	err := s.DB.QueryRowContext(ctx, query, shortCode).Scan(&url.ShortCode, &url.LongURL, &url.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return url, nil
}

// checks if the shortcode already exists in the db or not
func (s *MySqlStore) Exists(ctx context.Context, shortCode string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM urls WHERE short_code = ?)"

	err := s.DB.QueryRowContext(ctx, query, shortCode).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
