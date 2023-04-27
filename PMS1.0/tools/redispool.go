package tools

import "github.com/garyburd/redigo/redis"

func GetRedisPool() (c redis.Conn) {
	var pool *redis.Pool
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库的额最大链接数，0 表示没有限制
		IdleTimeout: 100, //最大空闲时间	（规定时间内没有操作的链接将返回连接池）
		Dial: func() (redis.Conn, error) { //初始化链接，链接哪个ip的redis
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	//defer pool.Close() //关闭之后就不能继续读取链接

	c = pool.Get() //从链接池中取出一个链接
	return c
}
