package repository

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jxskiss/base62"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRepository_PostLong(t *testing.T) {
	a := assert.New(t)

	err := godotenv.Load(".env.local")
	if err != nil {
		return
	}

	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		host, username, password, databaseName, port)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return
	}
	storage := New(db)

	const vk = "https://vk.com/"
	const shortVk = "https://urlshrt.com/1FMIUYZBzQ/"
	const google = "https://google.com/"

	testCases := []struct {
		name     string
		longUrl  string
		exp      string
		checkErr func(err error, msgAndArgs ...interface{}) bool
	}{
		{
			name:     "new shortUrl created",
			longUrl:  google,
			exp:      "https://urlshrt.com/" + base62.EncodeToString([]byte(google))[:10] + "/",
			checkErr: a.NoError,
		},
		{
			name:     "shortUrl already exists",
			longUrl:  vk,
			exp:      shortVk,
			checkErr: a.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			shortUrl, err := storage.PostLong(tc.longUrl)
			tc.checkErr(err)
			a.Equal(tc.exp, shortUrl)
		})
	}
	err = storage.db.Close()
	if err != nil {
		return
	}
}

func TestRepository_GetLong(t *testing.T) {
	a := assert.New(t)

	err := godotenv.Load(".env.local")
	if err != nil {
		return
	}

	host := os.Getenv("POSTGRES_HOST")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		host, username, password, databaseName, port)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return
	}
	storage := New(db)

	const vk = "https://vk.com/"
	const shortVk = "https://urlshrt.com/1FMIUYZBzQ/"
	const wrongUrl = "https://urlshrt.com/etwWRfas2a/"

	testCases := []struct {
		name     string
		shortUrl string
		exp      string
		checkErr func(err error, msgAndArgs ...interface{}) bool
	}{
		{
			name:     "found a shortUrl",
			shortUrl: shortVk,
			exp:      vk,
			checkErr: a.NoError,
		},
		{
			name:     "not found a shortUrl",
			shortUrl: wrongUrl,
			exp:      "",
			checkErr: a.Error,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			longUrl, err := storage.GetLong(tc.shortUrl)
			tc.checkErr(err)
			a.Equal(tc.exp, longUrl)
		})
	}
	err = storage.db.Close()
	if err != nil {
		return
	}
}
