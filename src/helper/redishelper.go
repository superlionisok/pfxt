package helper

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func RedisSet(key string, value string) (bool, error) {

	conn, err := redis.Dial("tcp", "10.10.15.154:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return false, err
	}
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set error:", err)
		return false, err
	}
	return true, nil

}
func RedisSetByTime(key string, value string, extsecond int) (bool, error) {

	conn, err := redis.Dial("tcp", "10.10.15.154:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return false, err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value, "EX", extsecond)
	if err != nil {
		fmt.Println("redis set error:", err)
		return false, err
	}
	return true, nil

}
func RedisGet(key string) (string, error) {

	conn, err := redis.Dial("tcp", "10.10.15.154:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return "", err
	}
	defer conn.Close()

	name, err := redis.String(conn.Do("GET", key))
	if err != nil {
		fmt.Println("redis get error:", err)
		return "", err
	} else {
		return name, nil
	}
}
