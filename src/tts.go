package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/shaipe/tide/net"
	"log"
	"os"
	"time"
)

/**
	运行:

go run tts.go -c tts.cnf -log xxx.log
build:
linux: GOOS=linux GOARCH=amd64 go build tts.go
windows: GOOS=windows GOARCH=amd64 go build tts.go
build后:

./tts -c tts.cnf -log xxx.log

	在linux下可以使用如下命令让应用在后台运行
	nohup command &

	可以通过log.SetFlags()自定议你想要表达的格式

	设置输出目的地log.SetOutput()

	os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)这是创建log文件.

	1如果log文件不存在，创建一个新的文件os.O_CREATE

	2打开文件的读写os.O_RDWR

	3将log信息写入到log文件，是继承到当前log文件，不是覆盖os.O_APPEND

	3log文件的权限位0666（即所有用户可读写）
 */

var (
	logFileName = flag.String("log", "tts.log", "Log file name")
	confPath = flag.String("c", "tts.cnf", "配置文件路径: 给定运行的配置信息")
)

// 任务配置
type Task struct {
	// 任务名称: name
	Name string
	// 请求的url地址: url
	Url string
	// 请求数据: data
	Data map[string] interface{}
	// 请求方式: method
	Method string
	// 执行时间: time
	ExecuteTime time.Time
}



// 初始化配置
func InitConfig(cnfPath string) []Task {

	var tasks []Task

	cfg, err := ini.Load(cnfPath)

	if err != nil{
		log.Printf("config %v not found", cnfPath)
	}

	sects := cfg.Sections()

	for k := range sects{
		sec := sects[k]
		// 排除默认
		if sec.Name() == "DEFAULT" {
			continue
		}

		// time layout 是代表: go语言2006开始策划, 后面分别为: 1,2,3,4,5   2006-01-02 15:04:05.999999999 -0700 MST
		tmStr := sec.Key("time").String()
		tmStr = time.Now().Format("2006-01-02") + " " + tmStr
		//fmt.Println(tmStr)
		tm, _ := time.Parse("2006-01-02 15:04:05", tmStr)

		// 将post数据转换为字典
		d := sec.Key("data").String()
		var m map[string] interface{}
		err = json.Unmarshal([]byte(d), &m)

		if err != nil {
			fmt.Println("Unmarshal failed, ", err)
			return tasks
		}

		t := Task{
			Name: sec.Key("name").String(),
			Url: sec.Key("url").String(),
			Data: m,
			Method: sec.Key("method").String(),
			ExecuteTime: tm,
		}
		tasks = append(tasks, t)
	}
	return tasks
}


// 循环工作任务
func LoopWorker(tasks []Task){

	for _, task := range tasks{

		start(task)

	}

}

func start(task Task)  {
	fmt.Println(task.Name, task.Url)
	go func() {
		fmt.Println("ewweewew")
		for {
			fmt.Println(task.Url)
			go reqUrl(task.Name, task.Url, task.Method, task.Data)
			now := time.Now()
			next := task.ExecuteTime
			if now.Before(next) {
				next = next.Add(time.Hour * 24)
			}
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}


func reqUrl(name, url, method string, data map[string] interface{}){
	// continue
	html, err := net.Fetch(url, method, data, nil)
	if err != nil{
		defer func() {
			log.Println(err)
		}()
	}
	log.Printf("do worker name: [%v], request result: %v", name, html)
}

// 程序启动入口
func main(){

	flag.Parse()

	// fmt.Println(*logFileName)
	// return

	// fmt.Println(time.Time())

	// 给定配置路径
	tasks := InitConfig(*confPath)
	fmt.Println(tasks)

	// 配置日志设置
	SetLog()

	// 开始执行计划
	LoopWorker(tasks)
}




// 设置日志记录
func SetLog(){
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





