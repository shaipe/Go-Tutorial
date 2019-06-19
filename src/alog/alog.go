package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
	"context"
)

/**
删除语句执行情况
*/
type AppLog struct {
	// 平台
	Platform string 			// DevicePlatform(android/ios)
	// 崩溃时间(采集时间)
	CaptureTime time.Time		// CrashTime
	// 执行语句
	LogText string				// CrashLog
	// 记录时间
	AddTime time.Time			// CreateTime 服务器的记录时间
	// 程序名称
	AppName string
	// 接口域名
	ApiDomain string			// AppPCDomain
	// 版本号
	Version string				// VersionCode
	// 构建版本号
	BuildNumber string
	// 包名
	BundleId string
	// 应用Id
	AppID string
	// H5域名
	H5Domain string				// AppH5Domain
	// 设备序列号
	DeviceSerial string
	// 设备信息
	DeviceInfo string
	// 运营主体
	Proprietor string
	// 运营主体Id
	ProprietorId string
	// 登录角色
	LoginRole string
	// 登录账号
	LoginAccount string
	// 登录名称
	LoginName string

}

/**
Mongo数据库连接配置
*/
type MongoConfig struct {
	// 服务器地址
	server string
	// 端口
	port int
	// 用户名
	username string
	// 密码
	password string
	// 数据库
	database string
}

/**
数据库配置
 */
var (
	DbConfig MongoConfig
	confPath = flag.String("c", "/bin/alog/alog.conf", "配置文件路径: 给定运行的配置信息")
)
/**
启动函数
 */
func main(){

	flag.Parse()

	// 读取配置文件路径
	loadConfig(*confPath)
	// fmt.Println(DbConfig)

	http.HandleFunc("/", index)
	http.HandleFunc("/rec", logRoute)

	fmt.Println("start server", time.Now().UnixNano())

	err := http.ListenAndServe(":7007", nil)
	if err !=nil{
		log.Fatal("list 7007")
	}
}

/**
加载配置文件
 */
func loadConfig(confPath string){
	cfg, err := ini.Load(confPath)

	if err != nil{
		log.Printf("config %v not found", confPath)
	}

	sects := cfg.Sections()

	for k := range sects {
		sec := sects[k]
		// 排除默认
		if sec.Name() == "DEFAULT" {
			continue
		}else if sec.Name() == "mongo" {

			p, _ := sec.Key("port").Int()
			c := MongoConfig{
				server:sec.Key("server").String(),
				port: p,
				username:sec.Key("username").String(),
				password:sec.Key("password").String(),
				database:sec.Key("database").String(),
			}
			// fmt.Println(c)
			DbConfig = c
		}
	}
}


/**
接口首页
*/
func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "a-log")
}

/**
记录日志
 */
func logRoute(w http.ResponseWriter, r *http.Request){
	//接受post请求，然后打印表单中key和value字段的值
	if r.Method == "POST" {
		// 对Form表单进行解析
		r.ParseForm()
		// fmt.Println(len(r.Form))
		// 判断是否post为表单
		if len(r.Form) > 0{
			go parseFormLog(r.Form)

		} else {	// 对传入的json字符进行处理
			//获取post的json数据
			content, _ := ioutil.ReadAll(r.Body)
			//
			go parseJsonLog(content)
			// fmt.Println(url.QueryUnescape(string(content)))
		}

		fmt.Fprintf(w, "received")
	} else{
		fmt.Fprintf(w, "log")
	}
}

/**
解析表单数据到对象
 */
func parseFormLog(frm url.Values){

	// fmt.Println(frm)

	tmStr := frm.Get("CaptureTime")
	// fmt.Println(tmStr)
	tm, _ := time.Parse("2006-01-02 15:04:05", tmStr)

	appLog := AppLog{
		ApiDomain: frm.Get("ApiDomain"),
		LoginAccount: frm.Get("LoginAccount"),
		Version:  frm.Get("Version"),
		DeviceInfo:  frm.Get("DeviceInfo"),
		DeviceSerial:  frm.Get("DeviceSerial"),
		LoginName: frm.Get("LoginName"),
		LoginRole: frm.Get("LoginRole"),
		Platform: frm.Get("Platform"),
		CaptureTime: tm,
		H5Domain: frm.Get("H5Domain"),
		AppID: frm.Get("AppID"),
		AppName: frm.Get("AppName"),
		Proprietor: frm.Get("Proprietor"),
		ProprietorId: frm.Get("ProprietorId"),
		BuildNumber: frm.Get("BuildNumber"),
		BundleId: frm.Get("BundleId"),
		LogText: frm.Get("LogText"),
		AddTime: time.Now(),
	}
	// for k, v := range frm {
	//	fmt.Println(k, v)
	// }
	// 写入数据库
	RecordLog(appLog)
}

/**
解析Json数据到实体
 */
func parseJsonLog(jsonByte []byte){
	var appLog AppLog
	json.Unmarshal(jsonByte, &appLog)
	// 给定添加时间
	appLog.AddTime = time.Now()

	// 写入数据库
	RecordLog(appLog)
}


/**
记录日志
 */
func RecordLog(appLog AppLog){

	url := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", DbConfig.username, DbConfig.password,
		DbConfig.server, DbConfig.port, DbConfig.database)

	fmt.Println(appLog)

	// 设置连接选项
	clientOptions := options.Client().ApplyURI(url)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("success")

	collection := client.Database("alogs").Collection("log_crash_logs")

	//
	// d := AnalyzeResult{ time.Now(), "TableName", "Sql"}
	// 写入单个文件
	// res, err := collection.InsertOne(context.TODO(), d)

	res, err := collection.InsertOne(context.TODO(), appLog)

	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(res)
}