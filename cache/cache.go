package cache

import (
	"sync"
	"time"
)

type item[T any] struct {
	until time.Time
	value *T
}

func (item item[T]) expired() bool {
	return item.until.Before(time.Now())
}

type Cache[T any] struct {
	m   map[string]item[T]
	mx  sync.Mutex
	ttl time.Duration
}

func New[T any](ttl time.Duration) (cache *Cache[T]) {
	return &Cache[T]{
		m:   make(map[string]item[T]),
		mx:  sync.Mutex{},
		ttl: ttl,
	}
}

func (ca *Cache[T]) Put(id string, value *T) {
	ca.mx.Lock()
	defer ca.mx.Unlock()

	ca.m[id] = item[T]{
		until: time.Now().Add(ca.ttl),
		value: value,
	}
}

func (ca *Cache[T]) Get(id string) (value *T) {
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

	return item.value
}

func (ca *Cache[T]) Scrub() {
	ca.mx.Lock()
	defer ca.mx.Unlock()

	for id, item := range ca.m {
		if item.expired() {
			delete(ca.m, id)
		}
	}
}
