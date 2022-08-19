package goredis

//Get 返回 key 所关联的字符串值
func (c *Client) Get(key string) (string, error) {
	if c.connErr != nil {
		return "", c.connErr
	}
	return c.pool.Cmd("GET", key).Str()
}

//Set 哈希表 key 中的域 field 的值设为 value
func (c *Client) Set(key string, value interface{}, seconds int) (string, error) {
	if c.connErr != nil {
		return "", c.connErr
	}
	if seconds > 0 {
		return c.pool.Cmd("SET", key, value, "EX", seconds).Str()
	}
	return c.pool.Cmd("SET", key, value).Str()
}

//SetNX 将 key 的值设为 value ，当且仅当 key 不存在
func (c *Client) SetNX(key string, value interface{}) (int, error) {
	if c.connErr != nil {
		return 0, c.connErr
	}
	return c.pool.Cmd("SETNX", key, value).Int()
}

//GetSet 将给定 key 的值设为 value ，并返回 key 的旧值(old value)
func (c *Client) GetSet(key string, value interface{}) (string, error) {
	if c.connErr != nil {
		return "", c.connErr
	}
	return c.pool.Cmd("GETSET", key, value).Str()
}

//IncrByFloat 为 key 中所储存的值加上浮点数增量 increment
func (c *Client) IncrByFloat(key string, value interface{}) (float64, error) {
	if c.connErr != nil {
		return 0, c.connErr
	}
	return c.pool.Cmd("INCRBYFLOAT", key, value).Float64()
}

//Exists 检查给定 key 是否存在
func (c *Client) Exists(key string) (int, error) {
	if c.connErr != nil {
		return 0, c.connErr
	}
	return c.pool.Cmd("EXISTS", key).Int()
}

//Del 删除给定的一个或多个 key
func (c *Client) Del(key string) (int, error) {
	if c.connErr != nil {
		return 0, c.connErr
	}
	return c.pool.Cmd("DEL", key).Int()
}

//Expire 为给定 key 设置生存时间，当 key 过期时(生存时间为 0 )，它会被自动删除
func (c *Client) Expire(key string, seconds int) (int, error) {
	if c.connErr != nil {
		return 0, c.connErr
	}
	return c.pool.Cmd("EXPIRE", key, seconds).Int()
}
