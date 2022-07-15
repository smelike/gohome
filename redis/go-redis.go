package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func connect() {
	// var rdb []map[int]interface{}
	for i := 0; i < 100; i++ {
		client := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       i,
		})
		// fmt.Println("Number: %v \n", i)
		fmt.Printf("Redis client: %v", client)
	}
	// fmt.Printf("rdb is %v", rdb)
}
