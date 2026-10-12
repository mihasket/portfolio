package cache

import (
	"sync"
	"time"
)

type entry[T any] struct {
	data      T
	expiresAt time.Time
}

type Cache[T any] struct {
	mu    sync.RWMutex
	ttl   time.Duration
	data  map[string]entry[T]
	fetch func(key string) (T, error)
}

func New[T any](ttl time.Duration, fetchFn func(key string) (T, error)) *Cache[T] {
	return &Cache[T]{
		ttl:   ttl,
		data:  make(map[string]entry[T]),
		fetch: fetchFn,
	}
}

func (c *Cache[T]) Get(key string) (T, error) {
	c.mu.RLock()
	e, ok := c.data[key]
	c.mu.RUnlock()

	if ok && time.Now().Before(e.expiresAt) {
		return e.data, nil
	}

	result, err := c.fetch(key)
	if err != nil {
		var zero T
		return zero, err
	}

	c.mu.Lock()
	c.data[key] = entry[T]{
		data:      result,
		expiresAt: time.Now().Add(c.ttl),
	}
	c.mu.Unlock()

	return result, nil
}
