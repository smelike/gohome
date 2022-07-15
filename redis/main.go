package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))
	// ctx := context.TODO()
	// connectRedis(ctx)

	connect()
	/* setToRedis("name", "redis-test")
	setToRedis("name2", "redis-test-2")

	val := getFromRedis("name")

	fmt.Println("First value with name key: ", val)

	values := getAllKeys("name*")

	fmt.Printf("All values: %v \n", values) */
}

func connectRedis(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	redisClient = client
}

func setToRedis(key, val string) {
	err := redisClient.Set(key, val, 20000000).Err()

	if err != nil {
		fmt.Println(err)
	}
}

func getFromRedis(key string) string {
	val, err := redisClient.Get(key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

func getAllKeys(key string) []string {
	keys := []string{}

	iter := redisClient.Scan(0, key, 0).Iterator()

	for iter.Next() {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	return keys
}
