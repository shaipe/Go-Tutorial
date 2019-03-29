package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func connect_redis() *redis.Client  {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.4.136:6380",
		Password: "366ec.redis", // no password set
		DB:       0,  // use default DB
	})

	return client
}

func main (){

	client := connect_redis()
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected result: ", pong)
	// fmt.Println(cache)
	client.Set("name1", "shanhuhai", 0)
	client.Del("name1")
	// fmt.Println(client.())
	// client.Strings
	kyes, _ := client.Keys("*").Result()
	fmt.Println(kyes)
	for _, v := range kyes{
		fmt.Println(v)
	}



	/*
	hash, _ := client.HGetAll("").Result()
	for k, v:= range hash{
		fmt.Printf("key: %v, value: %v ", k, v)
	}
	*/
}