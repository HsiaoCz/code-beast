package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	lock sync.RWMutex
	data map[string][]byte
}

func New() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Hash(key []byte) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	_, ok := c.data[string(key)]
	return ok
}

func (c *Cache) Delete(key []byte) error {
	c.lock.RLock()
	defer c.lock.RUnlock()
	delete(c.data, string(key))
	return nil
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	val, ok := c.data[string(key)]
	if !ok {
		return nil, fmt.Errorf("key (%s) not found", string(key))
	}
	return val, nil
}

func (c *Cache) Set(key, value []byte, ttl time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, ok := c.data[string(key)]
	if !ok {
		return errors.New("the key is exists")
	}
	c.data[string(key)] = value
	return nil
}
