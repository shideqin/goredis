package goredis

//Lpush 将一个或多个值 value 插入到列表 key 的表头
func (c *Client) Lpush(key string, value interface{}) (int, error) {
	if connErr != nil {
		return 0, connErr
	}
	return c.pool.Cmd("LPUSH", key, value).Int()
}

//Lpop 移除并返回列表 key 的头元素
func (c *Client) Lpop(key string) (string, error) {
	if connErr != nil {
		return "", connErr
	}
	return c.pool.Cmd("LPOP", key).Str()
}

//Llen 移除并返回列表 key 的头元素
func (c *Client) Llen(key string) (int, error) {
	if connErr != nil {
		return 0, connErr
	}
	return c.pool.Cmd("LLEN", key).Int()
}

//Rpush 将一个或多个值 value 插入到列表 key 的表尾(最右边)
func (c *Client) Rpush(key string, value interface{}) (int, error) {
	if connErr != nil {
		return 0, connErr
	}
	return c.pool.Cmd("RPUSH", key, value).Int()
}

//Rpop 移除并返回列表 key 的尾元素
func (c *Client) Rpop(key string) (string, error) {
	if connErr != nil {
		return "", connErr
	}
	return c.pool.Cmd("RPOP", key).Str()
}

//Lrem 根据参数 count 的值，移除列表中与参数 value 相等的元素
func (c *Client) Lrem(key string, count int, value interface{}) (int, error) {
	if connErr != nil {
		return 0, connErr
	}
	return c.pool.Cmd("LREM", key, count, value).Int()
}

//Lrange 返回列表 key 中指定区间内的元素，区间以偏移量 start 和 stop 指定
func (c *Client) Lrange(key string, start, stop int) ([]string, error) {
	if connErr != nil {
		return []string{}, connErr
	}
	return c.pool.Cmd("LRANGE", key, start, stop).List()
}

//Linsert 将值 value 插入到列表 key 当中，位于值 pivot 之前或之后
func (c *Client) Linsert(key string, pivot, value interface{}) (int, error) {
	if connErr != nil {
		return 0, connErr
	}
	return c.pool.Cmd("LINSERT", key, "BEFORE", pivot, value).Int()
}
