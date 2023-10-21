package in_memory

import (
	"github.com/jxskiss/base62"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemory_PostLong(t *testing.T) {
	a := assert.New(t)
	storage := New()

	const vk = "https://vk.com/"
	const shortVk = "https://urlshrt.com/1FMIUYZBzQ/"
	const google = "https://google.com/"

	storage.urls[shortVk] = vk

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
}

func TestInMemory_GetLong(t *testing.T) {
	a := assert.New(t)
	storage := New()

	const vk = "https://vk.com/"
	const shortVk = "https://urlshrt.com/1FMIUYZBzQ/"
	const wrongUrl = "https://urlshrt.com/etwWRfas2a/"

	storage.urls[shortVk] = vk

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
}
