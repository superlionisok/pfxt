package helper

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

func RedisSet(key string, value string) (bool, error) {

	conn, err := redis.Dial("tcp", "10.10.15.154:6379")
	if err != nil {
		//fmt.Println("connect redis error :", err)
		return false, err
	}
	defer conn.Close()
	_, err = conn.Do("SET", key, value)
	if err != nil {
		//fmt.Println("redis set error:", err)
		return false, err
	}
	return true, nil

}
func RedisSetByTime(key string, value string, extsecond int) (bool, error) {
	conn, err := redis.Dial("tcp", "10.10.15.154:6379")
	if err != nil {
		//fmt.Println("connect redis error :", err)
		return false, err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, value, "EX", extsecond)
	if err != nil {
		//fmt.Println("redis set error:", err)
		return false, err
	}
	return true, nil
}
func RedisGet(key string) (string, error) {

	conn, err := redis.Dial("tcp", "10.10.15.154:6379")
	if err != nil {
		//fmt.Println("connect redis error :", err)
		return "", err
	}
	defer conn.Close()

	name, err := redis.String(conn.Do("GET", key))
	if err != nil {
		//fmt.Println("redis get error:", err)
		return "", err
	} else {
		// 设置过期时间为24小时

		return name, nil
	}
}

func RedisGetAndResetTime(key string) (string, error) {

	conn, err := redis.Dial("tcp", "10.10.15.154:6379")
	if err != nil {
		//fmt.Println("connect redis error :", err)
		return "", err
	}
	defer conn.Close()

	name, err := redis.String(conn.Do("GET", key))
	if err != nil {
		//fmt.Println("redis get error:", err)
		return "", err
	} else {
		// 设置过期时间为24小时
		n, _ := conn.Do("EXPIRE", key, 1200)
		if n == int64(1) {
			//fmt.Println("success")
		} else {
			beego.Notice("设置Redis过期时间错误")
		}

		return name, nil
	}
}
