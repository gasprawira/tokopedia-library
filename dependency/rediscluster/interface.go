package rediscluster

type Client interface {
	Del(key string) (int, error)
	Command(command string, args ...interface{}) (interface{}, error)
	Get(key string) (string, error)
	Setex(key string, value string, ttl int) error
	HGet(key, field string) (string, error)
	HSet(key string, field string, value string, ttl int) error
	HMGet(key string, fields []string) ([]string, error)
	HMSet(key string, ttl int, m map[string]string) error
	HMSetWithoutTTL(key string, m map[string]string) error
	IncrBy(key string, amount int) (int, error)
	DecrBy(key string, amount int) (int, error)
	Expire(key string, ttl int) (bool, error)
	EXISTS(key string) (bool, error)
	HExists(key string, field string) (bool, error)
	Setnx(key string, value string, ttl int) error
	HGetAll(key string) ([]string, error)
	MGet(key ...string) ([]string, error)
	ZAddPairs(key string, pairs map[string]float64, ttl int) (int, error)
	ZRange(key string, start, stop int) ([]string, error)
	ZRevRange(key string, start, stop int) ([]string, error)
}
