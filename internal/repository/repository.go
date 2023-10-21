package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"urlshortener/internal/repository/types"
	"urlshortener/internal/short_creator"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) *repository {
	return &repository{db}
}

func searchShort(db *sql.DB, longUrl string) (string, error) {
	var res string
	q := `SELECT short FROM urls WHERE long = $1`
	err := db.QueryRow(q, longUrl).Scan(&res)
	if err != nil {
		return "", err
	}
	return res, nil
}

func searchLong(db *sql.DB, shortUrl string) (string, error) {
	var res string
	q := `SELECT long FROM urls WHERE short = $1`
	err := db.QueryRow(q, shortUrl).Scan(&res)
	if err != nil {
		return "", err
	}
	return res, nil
}

func insertNew(db *sql.DB, urLs *types.URLs) error {
	q := `INSERT INTO urls (long, short) values ($1, $2)`
	_, err := db.Exec(q, urLs.Long, urLs.Short)
	if err != nil {
		return fmt.Errorf("insert operation failed: %w", err)
	}
	return nil
}

func (r *repository) PostLong(longUrl string) (string, error) {
	var err error
	urls := new(types.URLs)
	urls.Long = longUrl
	urls.Short, err = searchShort(r.db, urls.Long)
	if err == nil {
		return string(urls.Short), nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("shortURL search failed: %w", err)
	}
	for {
		urls.Short = short_creator.CreateShortUrl(longUrl)
		_, err = searchLong(r.db, urls.Short)
		if errors.Is(err, sql.ErrNoRows) {
			break
		}
		if err != nil {
			return "", fmt.Errorf("longURL search failed: %w", err)
		}
	}
	if err = insertNew(r.db, urls); err != nil {
		return "", err
	}
	return urls.Short, nil
}

func (r *repository) GetLong(shortUrl string) (string, error) {
	longUrl, err := searchLong(r.db, shortUrl)
	if err != nil {
		return "", fmt.Errorf("no such short URL: %w", err)
	}
	return longUrl, nil
}
