package goredis

import (
	"log"
	"os"

	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
)

//Client redis连接结构体
type Client struct {
	pool *pool.Pool
}

//Conn 连接redis
func Conn(host, passwd string, size int) *Client {
	defer func() {
		if r := recover(); r != nil {
			log.Println("redis recover in ", r)
			os.Exit(0)
		}
	}()
	df := func(network, addr string) (*redis.Client, error) {
		client, err := redis.Dial(network, addr)
		if err != nil {
			return nil, err
		}
		if passwd != "" {
			if err = client.Cmd("AUTH", passwd).Err; err != nil {
				client.Close()
				return nil, err
			}
		}
		return client, nil
	}
	p, err := pool.NewCustom("tcp", host, size, df)
	if err != nil {
		panic("pool " + err.Error())
	}
	return &Client{
		pool: p,
	}
}

//GetConn 从连接池获取连接
func (c *Client) GetConn() *redis.Client {
	defer func() {
		if r := recover(); r != nil {
			log.Println("redis recover in ", r)
		}
	}()
	conn, err := c.pool.Get()
	if err != nil {
		conn.Close()
		panic("conn " + err.Error())
	}
	return conn
}

//PutConn 放回连接池
func (c *Client) PutConn(conn *redis.Client) {
	c.pool.Put(conn)
}
