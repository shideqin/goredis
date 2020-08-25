package goredis

import (
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

//Client redis连接结构体
type Client struct {
	pool       *pool.Pool
	ErrRespNil error
}

var connErr error

//Conn 连接redis
func Conn(host, passwd string, size int) *Client {
	defer func() {
		if r := recover(); r != nil {
			connErr = r.(error)
		}
	}()
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
		panic(err)
	}
	connErr = nil
	return &Client{
		pool:       p,
		ErrRespNil: redis.ErrRespNil,
	}
}

func Ping() error {
	return connErr
}
