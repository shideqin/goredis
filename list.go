package goredis

import (
	"github.com/mediocregopher/radix.v2/redis"
)

//Lpush 将一个或多个值 value 插入到列表 key 的表头
func (c *Client) Lpush(key string, value interface{}) (int, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("LPUSH", key, value).Int()
}

//Lpop 移除并返回列表 key 的头元素
func (c *Client) Lpop(key string) (string, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	r := conn.Cmd("LPOP", key)
	if r.IsType(redis.Nil) {
		return "", nil
	}
	return r.Str()
}

//Llen 移除并返回列表 key 的头元素
func (c *Client) Llen(key string) (int, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("LLEN", key).Int()
}

//Rpush 将一个或多个值 value 插入到列表 key 的表尾(最右边)
func (c *Client) Rpush(key string, value interface{}) (int, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("RPUSH", key, value).Int()
}

//Rpop 移除并返回列表 key 的尾元素
func (c *Client) Rpop(key string) (string, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	r := conn.Cmd("RPOP", key)
	if r.IsType(redis.Nil) {
		return "", nil
	}
	return r.Str()
}

//Lrem 根据参数 count 的值，移除列表中与参数 value 相等的元素
func (c *Client) Lrem(key string, count int, value interface{}) (int, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("LREM", key, count, value).Int()
}

//Lrange 返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定
func (c *Client) Lrange(key string, start, stop int) ([]string, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("LRANGE", key, start, stop).List()
}
