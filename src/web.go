package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Todo struct {
	Title string
	Done  bool
}

/**
定列TodoPage数据类型
 */
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}


func sumWeb(x,y int,c chan int){
	time.Sleep(5*time.Second)
	c <- x + y
}

func anotherWeb(c chan int){
	fmt.Println(<-c)      //管道有数据了直接继续执行，相当于异步通知
	//do something else...
}

/**
直接输出文本内容
 */
func sayHello(w http.ResponseWriter, r *http.Request){
	c := make (chan int)
	go sumWeb(24,18,c)
	go anotherWeb(c)
	fmt.Fprintf(w, "我运行了")
}

/**
	从文本模板中输出内容
 */
func renderTemplate(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("src/template/index.html",
		"src/template/header.html",
		"src/template/footer.html")

	if err != nil {
		errStr := err.Error()

		errStr += getCurrentDirectory()

		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

/**
	获取子目录路径
 */
func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

/**
	获取指定目录的父级目录地址
 */
func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

/**
	获取当前运行的目录地址
 */
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}


/**
	运行主函数
 */
func main(){
	// http.HandleFunc("/", sayHello)
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/tpl", renderTemplate)
	log.Println("我启动了")
	err := http.ListenAndServe(":7000", nil)
	if err !=nil{
		log.Fatal("list 7000")
	}
}