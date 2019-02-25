package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)


//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	fmt.Println("===", dirPth)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			if fi.Name() == "node_modules" || fi.Name() == ".git"{
				continue
			}
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			go GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			if fi.Name() == "node_modules" || fi.Name() == ".git"{
				continue
			}
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			go GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}


func main(){

	dirPath := os.Args[1]

	fmt.Println(dirPath)

	files, dirs, err := GetFilesAndDirs(dirPath)

	if err != nil{
		fmt.Println(err)
	}

	for _, dir := range dirs {
		fmt.Printf("获取的文件夹为[%s]\n", dir)
	}

	for _, table := range dirs {
		temp, _, _ := GetFilesAndDirs(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	for _, table1 := range files {
		fmt.Printf("获取的文件为[%s]\n", table1)
	}

	fmt.Printf("=======================================\n")
	xfiles, _ := GetAllFiles(dirPath)
	for _, file := range xfiles {
		fmt.Printf("获取的文件为[%s]\n", file)
	}
}