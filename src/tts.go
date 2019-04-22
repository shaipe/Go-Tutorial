package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/shaipe/tide/net"
	"log"
	"net/http"
	"os"
	"strings"
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
	// oper = flag.String("s", "", "输入参数,用于重启或停止")
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
	// 间隔
	Interval int
	// 重复定义, None-不重复, day-按天, hour-小时, minute-分, second-秒
	Repeat string

}



// 初始化配置
func InitConfig(cnfPath string) []Task {
	// fmt.Println(confPath)
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
		tmStr := sec.Key("executeTime").String()
		// fmt.Println(tmStr)
		tmStr = time.Now().Format("2006-01-02") + " " + tmStr
		// fmt.Println(tmStr)
		tm, _ := time.Parse("2006-01-02 15:04:05", tmStr)

		// 将post数据转换为字典
		d := sec.Key("data").String()
		var m map[string] interface{}
		err = json.Unmarshal([]byte(d), &m)

		if err != nil {
			fmt.Println("Unmarshal failed, ", err)
			return tasks
		}

		interval, _ := sec.Key("interval").Int()

		// 初始化任务对象
		t := Task{
			Name: sec.Key("name").String(),
			Url: sec.Key("url").String(),
			Data: m,
			Method: sec.Key("method").String(),
			ExecuteTime: tm,
			Repeat: sec.Key("repeat").String(),
			Interval: interval,
		}

		fmt.Println(t)

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

// 获取下次执行时间
func getNext(task Task) time.Time {
	//fmt.Println(task)
	// 获取当前时间
	now := time.Now()
	// 默认执行时间为下次执行时间
	next := task.ExecuteTime

	if next.Before(now){
		switch strings.ToLower(task.Repeat) {
		// 只执行一次
		case "none":
			next = now
		case "day":
			next = next.Add(time.Hour * 24)
		case "hour":
			next = now.Add(time.Hour * time.Duration(task.Interval))
		case "minute":
			next = now.Add(time.Minute * time.Duration(task.Interval))
		case "second":
			next = now.Add(time.Minute * time.Duration(task.Interval))
		}
	}
	// fmt.Println(next)
	return next
}

// 启动
func start(task Task)  {
	//fmt.Println(task.Name, task.Url)
	go func() {
		// fmt.Println("ewweewew")
		for {
			now := time.Now()
			next := getNext(task)
			if next.Before(now) {
				break
			}
			// go reqUrl(task.Name, task.Url, task.Method, task.Data)
			log.Printf("next execute: %v", next)
			t := time.NewTimer(next.Sub(now))
			<- t.C
			go reqUrl(task.Name, task.Url, task.Method, task.Data)
		}
	}()
}


func reqUrl(name, url, method string, data map[string] interface{}){
	log.Printf("request url: %v method: %v data: %v", url, method, data)
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

	//if *oper == "stop"{
	//	os.Exit(1)
	//} else if *oper == "reload"{
	//}

	// fmt.Println(*logFileName)
	// return

	// fmt.Println(time.Time())

	// 给定配置路径
	tasks := InitConfig(*confPath)
	// fmt.Println(tasks)

	// 配置日志设置
	SetLog()

	// 开始执行计划
	go LoopWorker(tasks)

	// 禁止 main 函数退出
	// defer func() { select {} }()

	starSite()
}



func starSite() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
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





