package rediscluster

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"

	rc "github.com/gasprawira/tokopedia-library/dependency/rediscluster"
)

var (
	redisCfg *rc.Config
)

type Cluster struct {
	ClusterPool   *redisc.Cluster
	RetryCount    int
	RetryDuration time.Duration
}

// InitializeRedis create new redis connection. nb: Idle timeout is in second
func InitializeRedisCluster(address, password string, maxActive, maxIdle int) (*Cluster, error) {
	var err error

	if maxActive <= 0 {
		maxActive = 600
	}
	if maxIdle <= 0 {
		maxIdle = 3
	}

	rediscCfg := rc.Config{
		Host:               address,
		RetryCount:         100,
		RetryDuration:      1,
		MaxActive:          maxActive,
		MaxIdle:            maxIdle,
		IdleTimeout:        240,
		DialConnectTimeout: 5,
		DialPassword:       password,
	}

	rc, err := initCluster(&rediscCfg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rc, err
}

// InitCluster initialize cluster client
func initCluster(cfg *rc.Config) (*Cluster, error) {
	redisCfg = cfg

	pool := redisc.Cluster{
		StartupNodes: []string{redisCfg.Host},
		DialOptions:  genDialOption(),
		CreatePool:   createPool,
	}

	if err := pool.Refresh(); err != nil {
		return nil, err
	}

	cluster := &Cluster{
		ClusterPool:   &pool,
		RetryCount:    redisCfg.RetryCount,
		RetryDuration: time.Duration(redisCfg.RetryDuration),
	}

	return cluster, nil
}

func genDialOption() []redis.DialOption {
	options := []redis.DialOption{}

	if redisCfg.DialConnectTimeout != 0 {
		options = append(options, redis.DialConnectTimeout(time.Duration(redisCfg.DialConnectTimeout)*time.Second))
	}

	if redisCfg.DialWriteTimeout != 0 {
		options = append(options, redis.DialWriteTimeout(time.Duration(redisCfg.DialWriteTimeout)*time.Second))
	}

	if redisCfg.DialReadTimeout != 0 {
		options = append(options, redis.DialReadTimeout(time.Duration(redisCfg.DialReadTimeout)*time.Second))
	}

	if redisCfg.DialDatabase != 0 {
		options = append(options, redis.DialDatabase(redisCfg.DialDatabase))
	}

	if redisCfg.DialKeepAlive != 0 {
		options = append(options, redis.DialKeepAlive(time.Duration(redisCfg.DialKeepAlive)*time.Second))
	}

	if redisCfg.DialPassword != "" {
		options = append(options, redis.DialPassword(redisCfg.DialPassword))
	}

	return options
}

func createPool(address string, opts ...redis.DialOption) (*redis.Pool, error) {
	rpool := &redis.Pool{
		MaxActive:   redisCfg.MaxActive,
		MaxIdle:     redisCfg.MaxIdle,
		IdleTimeout: time.Duration(redisCfg.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address, opts...)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Second {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	if _, err := rpool.Dial(); err != nil {
		rpool.Close()
		return nil, err
	}

	return rpool, nil
}
