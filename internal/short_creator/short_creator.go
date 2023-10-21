package short_creator

import (
	"github.com/jxskiss/base62"
)

const prefix = "https://urlshrt.com/"

func CreateShortUrl(longUrl string) string {
	val := base62.EncodeToString([]byte(longUrl))
	return prefix + val[:10] + "/"
}
