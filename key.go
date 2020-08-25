package goredis

//Keys 查找所有符合给定模式 pattern 的 key
func (c *Client) Keys(pattern string) ([]string, error) {
	if connErr != nil {
		return []string{}, connErr
	}
	return c.pool.Cmd("KEYS", pattern).List()
}
