## About

`goredis` 基于[`radix.v2/redis`](https://github.com/mediocregopher/radix.v2/redis) 开发而来，提供一个简单、高可用的redis连接操作基础库。主要特性如下：

* 支持连接池控制；
* 支持连接auth认证；

## Install

```
go get github.com/shideqin/goredis.git
```

## Usage

```golang
//创建redis 连接
client := goredis.Conn("HOST", "PASSWD", poolSzie)

//设置key
str, err := client.Set("key","value",seconds)

//获取key
str, err := client.Get("key")


```
