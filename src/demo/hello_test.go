package demo

/**
	要想此文件能正常的运行,在命令行中进入此文件夹,然后再运行test
cd src/demo
go test -v

*/

import "testing"

func TestAdd(t *testing.T) {
	Add("this is a test")
}