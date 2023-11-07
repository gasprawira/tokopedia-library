package rediscluster

import (
	"fmt"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
)

// Del do DEL command to redis
// Cluster does not support multi command, will invesitage later
func (c *Cluster) Del(key string) (int, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return 0, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return 0, err
	}

	resp, err := redis.Int(rc.Do("DEL", key))
	if err != nil {
		return 0, err
	}

	if resp < 0 {
		return 0, fmt.Errorf("Unexpected redis response %d", resp)
	}

	return resp, nil
}

// Command do Free command to redis
func (c *Cluster) Command(command string, args ...interface{}) (interface{}, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return nil, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return nil, err
	}

	return rc.Do(command, args...)
}

// Get do GET command to redis
func (c *Cluster) Get(key string) (string, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return "", client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return "", err
	}

	resp, err := redis.String(rc.Do("GET", key))
	if err != nil && err != redis.ErrNil {
		return "", err
	}

	return resp, nil
}

// Setex do SETEX command to redis
func (c *Cluster) Setex(key string, value string, ttl int) error {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return err
	}

	// set default TTL
	if ttl == 0 {
		ttl = 3600
	}

	resp, err := redis.String(rc.Do("SET", key, value, "EX", ttl))
	if err != nil {
		return err
	}

	if !strings.EqualFold("OK", resp) {
		return fmt.Errorf("Unexpected redis response %s", resp)
	}

	return nil
}

// HGet do HGET command to redis
func (c *Cluster) HGet(key, field string) (string, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return "", client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return "", err
	}

	resp, err := redis.String(rc.Do("HGET", key, field))
	if err != nil && err != redis.ErrNil {
		return "", err
	}

	return resp, nil
}

// HSet do HSET command to redis
func (c *Cluster) HSet(key string, field string, value string, ttl int) error {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return err
	}

	// do HSET
	resp, err := redis.Int(rc.Do("HSET", key, field, value))
	if err != nil {
		return err
	}

	if resp < 0 || resp > 1 {
		return fmt.Errorf("Unexpected redis response %d", resp)
	}

	// set default TTL
	if ttl == 0 {
		ttl = 3600
	}

	// do EXPIRE
	respExpire, err := redis.Int(rc.Do("EXPIRE", key, ttl))
	if err != nil {
		return err
	}

	if respExpire < 0 || respExpire > 1 {
		return fmt.Errorf("Unexpected redis response %d", respExpire)
	}

	return nil
}

// HMGet do HMGET command to redis
func (c *Cluster) HMGet(key string, fields []string) ([]string, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return []string{}, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return []string{}, err
	}

	args := make([]interface{}, len(fields)+1)
	args[0] = key
	for i, v := range fields {
		args[i+1] = v
	}

	resp, err := redis.Strings(rc.Do("HMGET", args...))
	if err != nil && err != redis.ErrNil {
		return []string{}, err
	}

	return resp, nil
}

// HMSet do HMSET command to redis
func (c *Cluster) HMSet(key string, ttl int, m map[string]string) error {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return err
	}

	// do HMSET
	resp, err := redis.String(rc.Do("HMSET", redis.Args{}.Add(key).AddFlat(m)...))
	if err != nil {
		return err
	}

	if !strings.EqualFold("OK", resp) {
		return fmt.Errorf("Unexpected redis response %s", resp)
	}

	// set default TTL
	if ttl == 0 {
		ttl = 3600
	}

	// do EXPIRE
	respExpire, err := redis.Int(rc.Do("EXPIRE", key, ttl))
	if err != nil {
		return err
	}

	if respExpire < 0 || respExpire > 1 {
		return fmt.Errorf("Unexpected redis response %d", respExpire)
	}

	return nil
}

// HMSet Without TTL do HMSET command to redis without TTL
func (c *Cluster) HMSetWithoutTTL(key string, m map[string]string) error {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return err
	}

	// do HMSET
	resp, err := redis.String(rc.Do("HMSET", redis.Args{}.Add(key).AddFlat(m)...))
	if err != nil {
		return err
	}

	if !strings.EqualFold("OK", resp) {
		return fmt.Errorf("Unexpected redis response %s", resp)
	}

	return nil
}

// IncrBy do INCRBY command
func (c *Cluster) IncrBy(key string, amount int) (int, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return 0, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return 0, err
	}

	// do INCRBY
	resp, err := redis.Int(rc.Do("INCRBY", key, amount))
	if err != nil {
		return 0, err
	}

	return resp, nil
}

// DecrBy do DECRBY command
func (c *Cluster) DecrBy(key string, amount int) (int, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return 0, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return 0, err
	}

	// do DECRBY
	resp, err := redis.Int(rc.Do("DECRBY", key, amount))
	if err != nil {
		return 0, err
	}

	return resp, nil
}

// Expire to EXPIRE command to redis
func (c *Cluster) Expire(key string, ttl int) (bool, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return false, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return false, err
	}

	resp, err := redis.Bool(rc.Do("EXPIRE", key, ttl))
	if err != nil {
		return false, err
	}

	return resp, nil
}

// Exists to EXISTS command to redis
func (c *Cluster) EXISTS(key string) (bool, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return false, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return false, err
	}

	resp, err := redis.Bool(rc.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return resp, nil
}

// HExists determines if a hash field exists
func (c *Cluster) HExists(key string, field string) (bool, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return false, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return false, err
	}

	resp, err := redis.Bool(rc.Do("HEXISTS", key, field))
	if err != nil {
		return false, err
	}

	return resp, nil
}

// Setnx do SETNX command to redis
func (c *Cluster) Setnx(key string, value string, ttl int) error {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return err
	}

	// set default TTL
	if ttl == 0 {
		ttl = 3600
	}

	resp, err := redis.String(rc.Do("SET", key, value, "EX", ttl, "NX"))
	if err != nil {
		return err
	}

	if !strings.EqualFold("OK", resp) {
		return fmt.Errorf("Unexpected redis response %s", resp)
	}

	return nil
}

// HGetall do HGETALL command to redis
func (c *Cluster) HGetAll(key string) ([]string, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return []string{}, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return []string{}, err
	}

	resp, err := redis.Strings(rc.Do("HGETALL", key))
	if err != nil && err != redis.ErrNil {
		return []string{}, err
	}

	return resp, nil
}

func (c *Cluster) MGet(key ...string) ([]string, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return []string{}, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return []string{}, err
	}

	args := make([]interface{}, len(key))
	for i := range key {
		args[i] = key[i]
	}

	resp, err := redis.Strings(rc.Do("MGET", args...))
	if err != nil && err != redis.ErrNil {
		return []string{}, err
	}

	return resp, nil
}

// ZAddPairs do ZADD command to redis using pairs of map[string]float64 data, will add EXPIRE if ttl > 0
func (c *Cluster) ZAddPairs(key string, pairs map[string]float64, ttl int) (int, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return 0, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return 0, err
	}

	var pairsValInterface []interface{}
	pairsValInterface = append(pairsValInterface, key)
	for k, v := range pairs {
		pairsValInterface = append(pairsValInterface, v, k)
	}

	resp, err := redis.Int64(rc.Do("ZADD", pairsValInterface...))
	if err != nil {
		return int(resp), err
	}

	if ttl > 0 {
		// do EXPIRE
		respExpire, err := redis.Int(rc.Do("EXPIRE", key, ttl))
		if err != nil {
			return 0, err
		}

		if respExpire < 0 || respExpire > 1 {
			return respExpire, fmt.Errorf("Unexpected redis response %d", respExpire)
		}
	}

	return 0, err
}

// ZRange do ZRANGE command to redis
func (c *Cluster) ZRange(key string, start, stop int) ([]string, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return []string{}, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return []string{}, err
	}

	result, err := redis.Strings(rc.Do("ZRANGE", key, start, stop))
	if err != nil {
		return []string{}, err
	}

	return result, err
}

// ZRevRange do ZREVRANGE command to redis
func (c *Cluster) ZRevRange(key string, start, stop int) ([]string, error) {
	client := c.ClusterPool.Get()
	defer client.Close()

	if client.Err() != nil {
		return []string{}, client.Err()
	}

	rc, err := redisc.RetryConn(client, c.RetryCount, c.RetryDuration*time.Millisecond)
	if err != nil {
		return []string{}, err
	}

	result, err := redis.Strings(rc.Do("ZREVRANGE", key, start, stop))
	if err != nil {
		return []string{}, err
	}

	return result, err
}
