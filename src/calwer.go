package main

/**
	包引用:
		github.com/PuerkitoBio/goquery
	安装包:
		goquery包中引用的golang.org/x/net包,并使用了此包中的html包,在安装goquery包之前需要先安装golang.org/x/net包
		因为golang.org在大陆地区不能访问,直接使用go get方式安装不能成功,此处采用手动安装的方式进行:

		1. cd $GOPATH/src	# 进入go的工作区目录
		2. mkdir -p golang.org/x # 批量创建包对应的目录
		3. git clone https://github.com/golang/net	# 从git上获取包源码, golang.org提供的包在github上的地址为: https://github.com/golang
		4. git install net	# 把下载好的源码进行安装
		5. git test golang.org/x/net/html	# 测试包是否安装成功

		通过以上5个步骤可以完成golang.org/x/net包的安装,并测试其中的html包是否可用,接下来安装 goquery 包

		go get github.com/PuerkitoBio/goquery	# 安装goquery

*/


import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main()  {

	// 定义需要采集的页面地址
	url := "https://chengdu.anjuke.com/sale/b62/?kw=%E4%B8%87%E8%BE%B0%E4%B9%90%E5%B1%85&k_comm_id=879254&kw_type=3"

	// 获取远程地址的内容,并赋值给res对象和记录错误对象 err
	res, err := http.Get(url)

	// 如果获取发生错误记录输出日志
	if err != nil{
		log.Fatal(err)
	}

	//
	defer res.Body.Close()

	// 判断返回的资源对象的response状态是否为200,即成功对象,失败刚输出日志
	if res.StatusCode != 200{
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// 通过goquery根据获取的文本流创建一个文档对象
	doc, err := goquery.NewDocumentFromReader(res.Body)

	// 判断创建对象是否发生异常
	if err != nil{
		log.Fatal(err)
	}

	// 查找页面上的房源列表元素
	ul := doc.Find("#houselist-mod-new")

	// 查找价格对象并循环输出
	ul.Find(".price-det").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})


	fmt.Println(ul.Find(".price-det").Html())

}
