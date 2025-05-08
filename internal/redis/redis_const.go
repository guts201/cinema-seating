package redis

import "time"

const (
	cacheLocalMaxCount = 1000
	cacheLocalTime     = time.Minute * 5
)

const (
	// CacheKeyPrefix is the prefix for all cache keys.
	CacheKeyPrefix = "cache:"
	// CacheKeySeparator is the separator for cache keys.
	CacheKeySeparator = ":"
	// CacheKeyExpireTime is the default expiration time for cache keys.
	CacheKeyExpireTime = time.Minute * 5
	// CacheKeyExpireTimeShort is the short expiration time for cache keys.
	CacheKeyExpireTimeShort = time.Second * 30
	// CacheKeyExpireTimeLong is the long expiration time for cache keys.
	CacheKeyExpireTimeLong = time.Hour * 24 * 7
	// CacheKeyExpireTimeNever is the expiration time for cache keys that never expire.
	CacheKeyExpireTimeNever = 0
)
