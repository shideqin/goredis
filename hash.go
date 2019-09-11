package goredis

//Hget 哈希表 key 中给定域 field 的值
func (c *Client) Hget(key string, field interface{}) (string, error) {
	return c.pool.Cmd("HGET", key, field).Str()
}

//Hset 哈希表 key 中的域 field 的值设为 value
func (c *Client) Hset(key string, field, value interface{}) (int, error) {
	return c.pool.Cmd("HSET", key, field, value).Int()
}

//HGetAll 哈希表 key 中，所有的域和值
func (c *Client) HGetAll(key string) (map[string]string, error) {
	return c.pool.Cmd("HGETALL", key).Map()
}

//Hdel 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略
func (c *Client) Hdel(key string, field interface{}) (int, error) {
	return c.pool.Cmd("HDEL", key, field).Int()
}

//Hlen 哈希表 key 中域的数量
func (c *Client) Hlen(key string) (int, error) {
	return c.pool.Cmd("HLEN", key).Int()
}
