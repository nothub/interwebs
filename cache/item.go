package cache

import "time"

type item struct {
	until time.Time
	value any
}

func (item item) expired() bool {
	return item.until.Before(time.Now())
}
