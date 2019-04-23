package auth

import (
	"errors"
	"sync"
	"time"
)

type object struct {
	data   interface{}
	expire time.Time
}

func (o *object) expired() bool {
	return time.Now().UTC().After(o.expire)
}

// In-memory cache system with expiration feature.
type cache struct {
	expiration time.Duration
	cache      map[string]object
	size       uint64
	mux        sync.RWMutex
}

func newCache(size uint64, exp time.Duration, purge time.Duration) *cache {
	c := &cache{
		expiration: exp,
		cache:      make(map[string]object),
		size:       size,
		mux:        sync.RWMutex{},
	}
	// clear cache periodically.
	go func() {
		for range time.Tick(purge) {
			c.mux.Lock()
			for key, obj := range c.cache {
				if obj.expired() {
					delete(c.cache, key)
				}
			}
			c.mux.Unlock()
		}
	}()
	return c
}

func (c *cache) get(key string) (interface{}, error) {
	obj, err := c.find(key)
	if err != nil {
		return nil, err
	}
	if obj.expired() {
		c.remove(key)
		return nil, errors.New("expired")
	}
	return obj.data, nil
}

func (c *cache) set(key string, data interface{}) error {
	if _, err := c.find(key); err == nil {
		return errors.New("dupicated")
	}
	if err := c.add(key, data); err != nil {
		return err
	}
	return nil
}

func (c *cache) find(key string) (*object, error) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	obj, ok := c.cache[key]
	if !ok {
		return nil, errors.New("not found")
	}
	return &obj, nil
}

func (c *cache) add(key string, data interface{}) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	// Try to find slot that would be available for the new object.
	if len(c.cache) >= int(c.size) {
		slotAvailable := false
		if _, ok := c.cache[key]; ok {
			slotAvailable = true
		} else {
			for key, obj := range c.cache {
				if obj.expired() {
					delete(c.cache, key)
					slotAvailable = true
					break
				}
			}
		}
		if !slotAvailable {
			return errors.New("full")
		}
	}
	c.cache[key] = object{
		data:   data,
		expire: time.Now().UTC().Add(c.expiration),
	}
	return nil
}

func (c *cache) remove(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.cache, key)
}
