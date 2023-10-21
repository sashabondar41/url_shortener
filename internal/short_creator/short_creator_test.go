package short_creator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortCreator_CreateShortUrl(t *testing.T) {
	a := assert.New(t)

	const vk = "https://vk.com/"
	const google = "https://google.com/"

	testCases := []struct {
		name    string
		longUrl string
		expLen  int
	}{
		{
			name:    "correct length for google",
			longUrl: google,
			expLen:  len(prefix) + 11,
		},
		{
			name:    "correct length for vk",
			longUrl: vk,
			expLen:  len(prefix) + 11,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			shortUrl := CreateShortUrl(tc.longUrl)
			a.Equal(tc.expLen, len(shortUrl))
		})
	}
}
