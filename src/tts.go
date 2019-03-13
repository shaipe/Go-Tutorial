package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/shaipe/tide/net"
	"log"
	"os"
	"path/filepath"
	"time"
)

/**
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
)

// 任务配置
type Task struct {
	// 任务名称
	Name string
	// 请求的url地址
	Url string
	// 请求数据
	Data map[string] interface{}
	// 请求方式
	Method string
	// 执行时间
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

		tm, _ := sec.Key("time").Time()

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
	i := 0
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	//go func() {
		for {
			select {
			case <- ticker.C:
				i++
				DoWorker(i, tasks)
			}
		}
	//}()

}

// 工作任务执行
func DoWorker(i int, tasks []Task){

	for _, val := range tasks{

		fmt.Println(time.Now())
		// continue
		html, err := net.Fetch(val.Url, val.Method, val.Data, nil)
		if err != nil{
			log.Println(err)
		}
		log.Printf("do worker index [%v] name: [%v], request result: %v", i, val.Name, html)

		// 停止2秒再执行
		time.Sleep(20 * time.Second)
	}
	return

}

// 程序启动入口
func main(){

	// fmt.Println(time.Time())

	// 给定配置路径
	confPath, _ := filepath.Abs(os.Args[1])
	tasks := InitConfig(confPath)
	fmt.Println(tasks)

	flag.Parse()

	// fmt.Println(*logFileName)

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





