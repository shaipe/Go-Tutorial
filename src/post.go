package main

import (
	"encoding/json"
	"github.com/shaipe/tide/net"
	"log"
)

func main(){

	url := "http://172.16.1.66:8000/das"
	method := "POST"
	dataString  := "{\"databaseId\":44,\"appid\":\"tts_app\",\"method\":\"rows.del.update\", \"tablename\": \"sto_shop\"}"
	var data map[string] interface{}
	err := json.Unmarshal([]byte(dataString), &data)
	// continue
	html, err := net.Fetch(url, method, data, nil)
	if err != nil{
		defer func() {
			log.Println(err)
		}()
	}
	log.Printf("do worker name: [%v], request result: %v", method, html)
}