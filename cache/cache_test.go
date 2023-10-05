package cache

import (
	"testing"
	"time"
)

func Test_Cache_expiry(t *testing.T) {
	c := New[bool](10 * time.Millisecond)
	var v = true

	c.Put("foo", &v)
	time.Sleep(1 * time.Millisecond)
	foo := c.Get("foo")
	if foo == nil {
		t.Fatal("item expired but should not be")
	}

	c.Put("bar", &v)
	time.Sleep(10 * time.Millisecond)
	bar := c.Get("bar")
	if bar != nil {
		t.Fatal("item not expired but should be")
	}
}

func Test_item_expiry(t *testing.T) {
	for _, d := range []struct {
		validForDura time.Duration
		shouldExpire bool
	}{
		{
			validForDura: 0 * time.Minute,
			shouldExpire: true,
		},
		{
			validForDura: 1 * time.Minute,
			shouldExpire: false,
		},
	} {
		it := item[any]{
			until: time.Now().Add(d.validForDura),
		}
		time.Sleep(1 * time.Millisecond)
		if d.shouldExpire && !it.expired() {
			t.Fatal("item not expired but should be")
		}
		if !d.shouldExpire && it.expired() {
			t.Fatal("item expired but should not be")
		}
	}
}
