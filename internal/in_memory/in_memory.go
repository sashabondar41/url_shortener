package in_memory

import (
	"errors"
	"fmt"
	"sync"
	"urlshortener/internal/in_memory/types"
	"urlshortener/internal/short_creator"
)

var ErrNotFound = errors.New("url is not found")

type inMemory struct {
	urls types.InMemoryStorage
	mu   sync.RWMutex
}

func New() *inMemory {
	mp := make(types.InMemoryStorage)
	return &inMemory{urls: mp}
}

func searchShort(urls *types.InMemoryStorage, longUrl string) (string, error) {
	for short, long := range *urls {
		if longUrl == long {
			return short, nil
		}
	}
	return "", ErrNotFound
}

func searchLong(urls *types.InMemoryStorage, shortUrl string) (string, error) {
	long, ok := (*urls)[shortUrl]
	if !ok {
		return "", ErrNotFound
	}
	return long, nil
}
func (r *inMemory) PostLong(longUrl string) (string, error) {
	r.mu.RLock()
	shortUrl, err := searchShort(&r.urls, longUrl)
	r.mu.RUnlock()
	if err == nil {
		return shortUrl, nil
	}
	for {
		shortUrl = short_creator.CreateShortUrl(longUrl)
		r.mu.RLock()
		_, err = searchLong(&r.urls, shortUrl)
		r.mu.RUnlock()
		if err == ErrNotFound {
			break
		}
	}
	r.mu.Lock()
	r.urls[shortUrl] = longUrl
	fmt.Println(r.urls)
	r.mu.Unlock()
	return shortUrl, nil
}

func (r *inMemory) GetLong(shortUrl string) (string, error) {
	r.mu.RLock()
	longUrl, err := searchLong(&r.urls, shortUrl)
	r.mu.RUnlock()
	if err != nil {
		return "", fmt.Errorf("searchLong: %w", err)
	}
	return longUrl, nil
}
