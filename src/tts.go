package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

/**
	在linux下可以使用如下命令让应用在后台运行
	nohup command &
 */

var (
	logFileName = flag.String("log", "tts.log", "Log file name")
)

func main(){
	// 配置日志设置
	setLog()
	// 开始执行计划
	loopWorker()
}

// 设置日志记录
func setLog(){
	// 定义日志文件和创建日志文件
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	// 日志文件创建判断
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "tts start Failed")
		os.Exit(1)
	}
	// 设置日志的输出模式
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// write log
	log.Printf("tts service start! %v \n", "start log file")

}

// 循环工作任务
func loopWorker(){
	i := 0
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	//go func() {
		for {
			select {
			case <- ticker.C:
				i++

				doWorker(i)
			}
		}
	//}()

}


// 工作任务执行
func doWorker(i int){
	time.Sleep(2 * time.Second)
	log.Printf("start do worker i: %v ;", i)
}


func curl(url string){

}


