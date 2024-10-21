package cache

import (
	"context"
	"sync"
	"time"
)

const (
	janitorTTL = 1
)

type Cacher interface {
	Get(context.Context, string) (bool, interface{})
	Put(string, interface{})
	Delete(string)
	Janitor()
}

var _ Cacher = &Cache{}

type CacheItem struct {
	data interface{}
	ttl  time.Time
}

type Cache struct {
	m        *sync.RWMutex
	cache    map[string]CacheItem
	cacheTTL time.Duration
}

func InitCacheBucket(cacheTTL time.Duration) *Cache {
	var c = Cache{
		m:        &sync.RWMutex{},
		cache:    make(map[string]CacheItem, 0),
		cacheTTL: cacheTTL,
	}
	go c.Janitor()

	return &c
}

func (c *Cache) Get(ctx context.Context, k string) (isCached bool, data interface{}) {
	c.m.Lock()
	defer c.m.Unlock()

	var cI CacheItem

	cI, ok := c.cache[k]
	if !ok {
		return false, []byte{}
	}

	if cI.ttl.Before(time.Now()) {
		return false, []byte{}
	}

	return true, cI.data
}

func (c *Cache) Put(k string, data interface{}) {
	c.m.Lock()
	defer c.m.Unlock()
	c.cache[k] = CacheItem{
		data: data,
		ttl:  time.Now().Add(c.cacheTTL),
	}
}

func (c *Cache) Delete(k string) {
	c.m.Lock()
	defer c.m.Unlock()
	delete(c.cache, k)
}

func (c *Cache) Janitor() {
	t := time.NewTicker(time.Hour * janitorTTL)
	defer t.Stop()
	for range t.C {
		c.m.Lock()
		for k, v := range c.cache {
			if v.ttl.Before(time.Now()) {
				delete(c.cache, k)
			}
		}
		c.m.Unlock()
	}
}
