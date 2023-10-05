package cache

import (
	"sync"
	"time"
)

type m map[string]item
type Cache struct {
	m
	mx  sync.Mutex
	ttl time.Duration
}

func New(ttl time.Duration) (cache *Cache) {
	return &Cache{
		m:   make(map[string]item),
		mx:  sync.Mutex{},
		ttl: ttl,
	}
}

func (ca *Cache) Put(id string, value any) {
	ca.mx.Lock()
	defer ca.mx.Unlock()

	ca.m[id] = item{
		until: time.Now().Add(ca.ttl),
		value: value,
	}
}

func (ca *Cache) Get(id string) (value any) {
	ca.mx.Lock()
	defer ca.mx.Unlock()

	item, ok := ca.m[id]
	if !ok {
		return nil
	}

	if item.expired() {
		delete(ca.m, id)
		return nil
	}

	return item
}

func (ca *Cache) Scrub() {
	ca.mx.Lock()
	defer ca.mx.Unlock()

	for id, item := range ca.m {
		if item.expired() {
			delete(ca.m, id)
		}
	}
}
