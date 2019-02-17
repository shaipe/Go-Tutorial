package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "我运行了")
}


func main(){
	http.HandleFunc("/", sayHello)
	log.Println("我启动了")
	err := http.ListenAndServe(":7000", nil)
	if err !=nil{
		log.Fatal("list 7000")
	}
}