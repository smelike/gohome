package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func connect() {
	rdb := redis.NewClient(&redis.Options{
		PoolSize: 6,
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	fmt.Printf("rdb is %v", rdb)
}
