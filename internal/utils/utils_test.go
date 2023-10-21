package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils_ValidateUrl(t *testing.T) {
	a := assert.New(t)

	const vk = "https://vk.com/"
	const totallyWrongUrl = "aefwdasrgrsgrefa"
	const slightlyWrongUrl = "htps:/google.com/"

	testCases := []struct {
		name     string
		Url      string
		expErr   error
		checkErr func(err error, msgAndArgs ...interface{}) bool
	}{
		{
			name:     "correct url",
			Url:      vk,
			expErr:   ErrWrUrl,
			checkErr: a.NoError,
		},
		{
			name:     "totally incorrect url",
			Url:      totallyWrongUrl,
			expErr:   ErrWrUrl,
			checkErr: a.NoError,
		},
		{
			name:     "slightly incorrect url",
			Url:      slightlyWrongUrl,
			expErr:   ErrWrUrl,
			checkErr: a.NoError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_ = ValidateUrl(tc.Url)
		})
	}
}
