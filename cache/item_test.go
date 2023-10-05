package cache

import (
	"testing"
	"time"
)

func Test_item_expired(t *testing.T) {
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
		it := item{
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
