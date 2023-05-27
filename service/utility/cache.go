package utility

import (
	"context"
	"errors"
	"time"

	"github.com/Zhima-Mochi/go-authentication-service/external"
)

var _ external.Cache = (*Cache)(nil)

type Cache struct {
	memory map[string]interface{}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	c.memory[key] = value
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	if value, ok := c.memory[key]; ok {
		return value, nil
	}

	return nil, errors.New("not found")
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	delete(c.memory, key)
	return nil
}

func NewCache() *Cache {
	return &Cache{
		memory: map[string]interface{}{},
	}
}
