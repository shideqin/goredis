package goredis

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

var (
	//ErrRespNil 数据没有找到
	ErrRespNil = redis.ErrRespNil
)

//Client redis连接结构体
type Client struct {
	pool    *pool.Pool
	connErr error
}

//Conn 连接redis
func Conn(host, passwd string, size int) *Client {
	cli := &Client{}
	df := func(network, addr string) (*redis.Client, error) {
		client, err := redis.Dial(network, addr)
		if err != nil {
			return nil, err
		}
		if passwd != "" {
			if err = client.Cmd("AUTH", passwd).Err; err != nil {
				_ = client.Close()
				return nil, err
			}
		}
		return client, nil
	}
	p, err := pool.NewCustom("tcp", host, size, df)
	if err != nil {
		cli.connErr = err
		return cli
	}
	cli.pool = p
	return cli
}

//Ping 监测数据库连接
func (c *Client) Ping() error {
	if c.connErr != nil {
		return c.connErr
	}
	_, c.connErr = c.pool.Cmd("PING").Str()
	return c.connErr
}
