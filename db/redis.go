// filename redis.go
// author Yoshiya Ito <myon53@gmail.com>
// 2019-03-18
package db

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

var Redis *redis.Pool

// Redis Setup
// !!! You have to call this function after viper setting file
func SetupRedis() {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	maxConnections := viper.GetInt("redis.max_connections")
	Redis = redis.NewPool(func() (redis.Conn, error) {
		conn, err := redis.Dial("tcp", host+":"+port)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return conn, err
	}, maxConnections)
}
