package main

import (
	"fmt"

	"github.com/shideqin/goredis"
)

func main() {
	//连接
	client := goredis.Conn("127.0.0.1", "", 1)
	fmt.Println(client)

	//检测连接
	err := client.Ping()
	fmt.Println(err)
}
