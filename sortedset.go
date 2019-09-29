package goredis

//Zrange 返回有序集 key 中，指定区间内的成员
func (c *Client) Zrange(key string, start, stop int) ([]string, error) {
	return c.pool.Cmd("ZRANGE", key, start, stop).List()
}

//Zadd 将一个或多个 member 元素及其 score 值加入到有序集 key 当中
func (c *Client) Zadd(key string, score int, member interface{}) (int, error) {
	return c.pool.Cmd("ZADD", key, score, member).Int()
}

//Zrem 移除有序集 key 中的一个或多个成员，不存在的成员将被忽略
func (c *Client) Zrem(key string, member interface{}) (int, error) {
	return c.pool.Cmd("ZREM", key, member).Int()
}

//Zrank 返回有序集 key 中成员 member 的排名
func (c *Client) Zrank(key string, member interface{}) (int, error) {
	return c.pool.Cmd("ZRANK", key, member).Int()
}
