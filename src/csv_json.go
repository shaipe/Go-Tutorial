//
// 运行方式: go run src/csv_json.go -csv="/Users/shaipe/documents/kdn.csv" -json="xx.json"

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)


var (
	csvPath = flag.String("csv", "", "csv路径")
	jsonPath = flag.String("json", "res.json", "Json存放路径")
)

func main(){

	flag.Parse()

	if *csvPath == ""{
		fmt.Println("请通过-csv参数的方式给定csv路径")
		return
	}

	fmt.Println(*csvPath)

	if !exist(*csvPath){
		fmt.Println("未找到csv文件")
		return
	}

	csv, err := os.Open(*csvPath)

	if err != nil{
		fmt.Println(err)
		return
	}

	defer csv.Close()

	br := bufio.NewReader(csv)

	jmap := make(map[string]string)

	for{
		l, _, c := br.ReadLine()
		// 结束即退出循环
		if c == io.EOF{
			break
		}

		sl := strings.Split(string(l), ",")
		jmap[strings.TrimSpace(sl[0])] = strings.TrimSpace(sl[1])
	}

	js, err := json.Marshal(jmap)
	if err != nil{
		fmt.Println(err)
		return
	}
	writeFile(*jsonPath, string(js))

}

// 将内容写入文件中
func writeFile(path string, content string){
	f, err := os.Create(path)

	if err != nil{
		fmt.Println(err)
		return
	}

	defer f.Close()

	f.WriteString(content)

}


// 判断文件或文件夹是否存在
func exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil{
		if os.IsExist(err){
			return true
		}
		return false
	}
	return true
}