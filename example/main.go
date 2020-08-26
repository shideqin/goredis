package main

import (
	"fmt"

	"github.com/shideqin/goredis"
)

func main() {
	//连接
	client := goredis.Conn("127.0.0.1", "", 10)
	fmt.Println(client)

	//检测连接
	err := goredis.Ping()
	fmt.Println(err)
}
