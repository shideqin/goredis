package goredis

import "github.com/mediocregopher/radix.v2/redis"

//Hget 哈希表 key 中给定域 field 的值
func (c *Client) Hget(key string, field interface{}) (string, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	r := conn.Cmd("HGET", key, field)
	if r.IsType(redis.Nil) {
		return "", nil
	}
	return r.Str()
}

//Hset 哈希表 key 中的域 field 的值设为 value
func (c *Client) Hset(key string, field, value interface{}) (int, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("HSET", key, field, value).Int()
}

//HGetAll 哈希表 key 中，所有的域和值
func (c *Client) HGetAll(key string) (map[string]string, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("HGETALL", key).Map()
}

//Hdel 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略
func (c *Client) Hdel(key string, field interface{}) (int, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("HDEL", key, field).Int()
}

//Hlen 哈希表 key 中域的数量
func (c *Client) Hlen(key string) (int, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("HLEN", key).Int()
}

//检查给定域 field 是否存在于哈希表 hash 当中
func (c *Client) Hexists(key, field string) (int, error) {
	conn := c.GetConn()
	defer func() {
		c.PutConn(conn)
	}()
	return conn.Cmd("HEXISTS", key, field).Int()
}
