package scache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var cache *gocache.Cache

const cacheDuration = time.Second * 60
const cacheDurationLong = time.Minute * 5

func init() {
	cache = gocache.New(3*time.Minute, 10*time.Minute)
}

func Set(k string, x interface{}, d time.Duration) {
	cache.Set(k, x, d)
}

func Get(k string) (interface{}, bool) {
	return cache.Get(k)
}
