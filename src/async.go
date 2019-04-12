package main

import (
	"fmt"
	"time"
)

func sum(x,y int,c chan int){
	time.Sleep(5*time.Second)
	c <- x + y
}

func another(c chan int){
	fmt.Println(<-c)      //管道有数据了直接继续执行，相当于异步通知
	//do something else...
}

func main(){
	c := make (chan int)
	go sum(24,18,c)
	go another(c)
	fmt.Println("继续执行")
	//do something else...
	time.Sleep(60*time.Second)
}
