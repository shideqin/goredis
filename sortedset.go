package goredis

//Zrange 返回有序集 key 中，指定区间内的成员
func (c *Client) Zrange(key string, start, stop int) ([]string, error) {
	if connErr != nil {
		return []string{}, connErr
	}
	return c.pool.Cmd("ZRANGE", key, start, stop).List()
}

//Zadd 将一个或多个 member 元素及其 score 值加入到有序集 key 当中
func (c *Client) Zadd(key string, score int, member interface{}) (int, error) {
	if connErr != nil {
		return 0, connErr
	}
	return c.pool.Cmd("ZADD", key, score, member).Int()
}

//Zrem 移除有序集 key 中的一个或多个成员，不存在的成员将被忽略
func (c *Client) Zrem(key string, member interface{}) (int, error) {
	if connErr != nil {
		return 0, connErr
	}
	return c.pool.Cmd("ZREM", key, member).Int()
}

//Zrank 返回有序集 key 中成员 member 的排名
func (c *Client) Zrank(key string, member interface{}) (int, error) {
	if connErr != nil {
		return 0, connErr
	}
	return c.pool.Cmd("ZRANK", key, member).Int()
}

//Zcount 返回有序集 key 中， score 值在 min 和 max 之间(默认包括 score 值等于 min 或 max )的成员的数量
func (c *Client) Zcount(key string, min, max int) (int, error) {
	if connErr != nil {
		return 0, connErr
	}
	return c.pool.Cmd("ZCOUNT", key, min, max).Int()
}
