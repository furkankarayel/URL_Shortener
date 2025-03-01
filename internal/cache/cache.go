package cache

import (
	"sync"
)

type URLCache struct {
	store sync.Map
}

func NewURLCache() *URLCache {
	return &URLCache{
		store: sync.Map{},
	}
}

func (c *URLCache) Save(shortCode, originalURL string) {

	c.store.Store(shortCode, originalURL)
}

func (c *URLCache) Get(shortCode string) (string, bool) {
	value, ok := c.store.Load(shortCode)
	if !ok {
		return "", false
	}

	return value.(string), true
}

func (c *URLCache) FindValue(longURL string) (string, bool) {
	shortCode := ""
	found := false

	c.store.Range(func(key, value interface{}) bool {
		if value.(string) == longURL {
			shortCode = key.(string)
			found = true
			return false
		}
		return true
	})

	return shortCode, found
}

func (c *URLCache) Delete(shortCode string) {
	c.store.Delete(shortCode)
}
