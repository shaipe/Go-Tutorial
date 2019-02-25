package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CopyFile(srcName, dstName string) (written int64, err error) {
	fmt.Println(srcName, dstName)
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func Walk(srcPath, dstPath string){

	fmt.Println(srcPath, dstPath)
	dirs, err := ioutil.ReadDir(srcPath)

	if err != nil{
		fmt.Println("====", err)
		return
	}

	pthSep := string(os.PathSeparator)

	for _, fi := range dirs{
		fileName := fi.Name()
		if fi.IsDir(){
			if fileName == "__pycache__" || fileName == "node_modules" || fileName == ".git"{
				continue
			}

			Walk(srcPath + pthSep + fileName, dstPath + pthSep + fileName)
		} else{
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".py")
			if ok {
				srcFile := srcPath + pthSep + fileName
				fmt.Println(srcFile)
				_, err1 := os.Stat(dstPath)
				if err1 !=nil{
					os.MkdirAll(dstPath, 0777)
				}
				_, err := CopyFile( srcFile, dstPath + pthSep + fileName)

				if err !=nil{
					fmt.Println(err)
				}
			}
		}

	}
}


func main(){
	srcPath := os.Args[1]
	dstPath := os.Args[2]

	if srcPath == ""{
		fmt.Println("please input source dir")
	}

	if dstPath == ""{
		fmt.Println("please input dist dir")
	}

	_, err := os.Stat(dstPath)

	if err !=nil{
		os.Mkdir(dstPath, 0777)
	}

	Walk(srcPath, dstPath)

}