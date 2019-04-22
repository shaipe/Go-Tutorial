package main

import (
	"fmt"
	"log"
	"net/http"
)


/**
直接输出文本内容
*/
func mySayHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "我运行了")
}

/**
运行主函数
*/
func main(){
	// http.HandleFunc("/", sayHello)
	http.HandleFunc("/", mySayHello)
	log.Println("我启动了")
	err := http.ListenAndServe(":7000", nil)
	if err !=nil{
		log.Fatal("list 7000")
	}
}