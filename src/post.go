package main

import (
	"encoding/json"
	"github.com/shaipe/tide/net"
	"log"
)

func main(){

	url := "http://127.0.0.1:7007/rec"
	method := "POST"
	dataString  := "{\"ApiDomain\":\"http%3A//www.320.366ec.cn\",\"appid\":\"tts_app\",\"AppName\":\"rows.del.update\", \"tablename\": \"sto_shop\"}"
	// dataString = "ApiDomain=http%3A//www.320.366ec.cn&AppID&AppName=%E5%BF%AB%E9%A9%AC%E6%89%B9%E5%8F%91%E6%B5%8B%E8%AF%95&BuildNumber=322&BundleId=com.366EC.ECRtest&CaptureTime=2019-06-19 12:00:12"

	html:= postForm(url, dataString)
	// html := postJson(url, method, dataString)
	log.Printf("do worker name: [%v], request result: %v", method, html)
}

func postForm(url string, dataString string) string{
	var data map[string] interface{}
	err := json.Unmarshal([]byte(dataString), &data)
	// continue
	html, err := net.PostForm(url, data, nil)
	if err != nil{
		defer func() {
			log.Println(err)
		}()
	}

	return html
}

func postJson(url string, method string,dataString string) string{
	var data map[string] interface{}
	err := json.Unmarshal([]byte(dataString), &data)
	// continue
	html, err := net.Fetch(url, method, data, nil)
	if err != nil{
		defer func() {
			log.Println(err)
		}()
	}

	return html
}