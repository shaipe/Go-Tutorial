Go语言学习教程
---

Go 编程语言是一个开源项目，它使程序员更具生产力。

Go 语言具有很强的表达能力，它简洁、清晰而高效。得益于其并发机制， 用它编写的程序能够非常有效地利用多核与联网的计算机，其新颖的类型系统则使程序结构变得灵活而模块化。 Go 代码编译成机器码不仅非常迅速，还具有方便的垃圾收集机制和强大的运行时反射机制。 它是一个快速的、静态类型的编译型语言，感觉却像动态类型的解释型语言。


## 开发环境搭建

#### 下载Go发行版本
[官方二进制发行版](官方二进制发行版) 支持 FreeBSD（8-STABLE 发行版及以上）、Linux、Mac OS X（Snow Leopard 及以上）和 Windows 操作系统以及32位（386）和64位（amd64）的 x86 处理器架构。

#### Mac OS X安装包
Mac平台下载pkg文件根据提示安装

打开此包文件 并跟随提示来安装Go工具。该包会将Go发行版安装到 /usr/local/go 中。

此包应该会将 /usr/local/go/bin 目录放到你的 PATH 环境变量中。 要使此更改生效，你需要重启所有打开的终端回话。

**工作区环境设置**
```bash
vim ~/.bash_profile

# 添加入下代码
export GOPATH=/Users/UserName/workDir 
# UserName为计算机登录用户名, workDir为想要设置的go语言的工作空间目录

# 让配置生效
source ~/.bash_profile
```

#### Windows Msi安装
打开此MSI文件 并跟随提示来安装Go工具。默认情况下，该安装程序会将Go发行版放到 c:\Go 中。

此安装程序应该会将 c:\Go\bin 目录放到你的 PATH 环境变量中。 要使此更改生效，你需要重启所有打开的命令行。
Windows平台可以下载msi文件根据提示进行安装,windows平台安装完成以后需要重启系统,系统的环境变量才会生效

#### Linux 安装

到go语言的官方网站下载最新版本的二进制文件[Go语言官方下载地址](http://golang.org/dl)选择下载对应的操作系统以及cpu架构,如(centos_amd64)表示为Centos操作系统,amd64架构

比如下载以下文件: https://dl.google.com/go/go1.11.5.linux-armv6l.tar.gz

1. 使用tar命令进持解析,官方建议解压到/usr/local目录下
```bash
tar -zxf go1.11.5.linux-armv6l.tar.gz -C /usr/local
# 具命令执行需要使用root权限
```
2. 进入/usr/local 目录查看是否有 `go` 目录
```bash
# 进入local目录
cd /usr/local
# 查看local目录下的文件或目录
ls
# 进入go目录
cd go
# 运行go命令查看当前go的版本信息
go version
# go version go1.11.5 linux/amd64 显示以上信息表示正确
```

3. 环境变量的设置

**GOROOT, GOPATH, GOBIN, PATH**
为了让每次启动机器让这些环境变量都生效,可以把这信息添加到`profile' 中, (``~/.bash_profile``<单一用户> 或 ``/etc/profile``<所有用户>)
添加内容:

```shell
# 设置go开发环境的根目录到环境变量
export GOROOT=/usr/local/go
# 设置go语言的工作区目录,下面是设置了二个工作区目录
export GOPATH=~/workspace/go:~/goproject
# GOBIN 是存放编译后的可执行文件目录
export GOBIN=~/gobin
# 为了方便使用GO命令和go程序可执行文件在系统的任何目录下输入命令都可以执行编译安装好的可以执行文件
export PATH=$PATH:$GOROOT/bin/:$GOPATH/bin:$GOBIN
# 注意此处必须在原PATH上进行追加... 
```
4. 让配置环境变量生效

```bash
source <某个profile文件>
# 执行完后在任一目录下,输入go version 查看是否显示开始执行以上命令相同的显示即可验证是否安装正确
```

#### 测试你的安装

通过构建一个简单的程序来检查Go的安装是否正确，具体操作如下：

首先创建一个名为 hello.go 的文件，并将以下代码保存在其中：

```go
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```
接着通过 go 工具运行它：
```bash
# 安装好以后可以用以下命令检测环境是否安装好,查看go语言的版本
go version

$ go run hello.go
hello, world
# 若你看到了“hello, world”信息，那么你的Go已被正确安装。

```

## 交流社区

### 中文社区

 - [Golang 中国](http://www.golangtc.com/)：国内较早的 Go 社区，汇聚各类信息与服务
 - [Study Golang](http://studygolang.com/)：国内 Go 社区先驱，同样汇聚各类信息与服务
 - [Revel 交流论坛](http://gorevel.cn/)：[Revel](https://github.com/revel/revel) 框架的中文社区
 - [GoCN Forum](https://gocn.vip/)：Go 语言爱好者中文交流论坛

 ### 英文社区

- [Go Forum](https://forum.golangbridge.org/)：Go 语言爱好者英文交流论坛
- [golang-nuts 邮件列表](https://groups.google.com/forum/#!forum/golang-nuts)：Go 语言官方指定邮件列表讨论区
- [Go+ 社区](https://plus.google.com/u/0/communities/114112804251407510571)：Go 语言官方指定 G+ 社区

## 网址导航

- 官方：
    - [Go 中国站点](https://golang.google.cn/): Go 语言中国官方站点(无需翻墙)
	- [Playground](http://play.golang.org)：Go 语言代码在线运行
- 国内镜像：
	- [Go 指南国内镜像](http://tour.golangtc.com/)
	- [Go 语言国内下载镜像](http://www.golangtc.com/download)
	- [Go 官方网站国内镜像](http://docs.studygolang.com/)
- Web 框架：
	- [Macaron](https://go-macaron.com/)：模块化 Web 框架
	- [Beego](http://beego.me/)：重量级 Web 框架
	- [Revel](https://github.com/revel/revel)：较早成熟的重量级 Web 框架
	- [Martini](https://github.com/go-martini/martini): 一个强大为了编写模块化 Web 应用而生的 Go 语言框架
	- [Echo](https://echo.labstack.com/): 功能模块齐全, 上手容易, 文档示例齐全
	- [Gin](https://github.com/gin-gonic/gin)：轻量级 HTTP Web 框架
- ORM 以及数据库驱动：
	- [xorm](https://github.com/go-xorm/xorm)：支持 MySQL、PostgreSQL、SQLite3 以及 MsSQL
	- [mgo](http://labix.org/mgo)：MongoDB 官方推荐驱动
	- [gorm](https://github.com/jinzhu/gorm): 全功能 ORM (无限接近) 支持 MySQL、PostgreSQL、SQLite3 以及 MsSQL
- 辅助站点：
	- [Go Walker](https://gowalker.org)：Go 语言在线 API 文档
	- [gobuild.io](http://gobuild.io/)：Go 语言在线二进制编译与下载
	- [Rego](http://regoio.herokuapp.com/)：Go 语言正则在线测试
	- [gopm.io](https://gopm.io)：科学下载第三方包
    - [Json To Go struct](https://mholt.github.io/json-to-go/):Convert JSON to Go struct在线工具
- 开发工具：
    - [Emacs24](http://ftp.gnu.org/gnu/emacs/)：[配置脚本](https://github.com/wackonline/hack/blob/master/install-mint-dev/install-emacs.d.sh) / [(中文社区)](http://emacser.com/)
	- [LiteIDE](https://github.com/visualfc/liteide)
	- [Sublime Text 2/3](http://sublimetext.com)：[配置教程](http://my.oschina.net/Obahua/blog/110767)
	- [GoLand](https://www.jetbrains.com/go/?fromMenu)
    - [Atom](https://atom.io)：[配置插件](https://atom.io/packages/go-plus)（感觉还不错，类似 Sublime，配置比较简单）
    - [VIM](http://www.vim.org)：[配置插件](https://github.com/humiaozuzu/dot-vimrc)（嫌 vim 配置麻烦的童鞋可以直接用这个）
- 学习站点：
	- [Go by Example](https://gobyexample.com/)
	- [Go database/sql tutorial](http://go-database-sql.org/)
- 支持 Go 的云平台：
	- [Koding](https://koding.com/)
	- [Nitrous.IO](https://www.nitrous.io/)
	- [Get up and running with Go on Google Cloud Platform]( https://cloud.google.com/go/)
	- [AWS SDK for Go - Developer Preview](http://aws.amazon.com/cn/sdk-for-go/):=>[github](https://github.com/aws/aws-sdk-go)
	- [azure sdk for go](https://godoc.org/github.com/Azure/azure-sdk-for-go):=>[github](https://github.com/Azure/azure-sdk-for-go)
        - [How to Use CoreOS on Azure](https://azure.microsoft.com/zh-cn/documentation/articles/virtual-machines-linux-coreos-how-to/)
        - [Create Azure Web app with GoLang extension](https://azure.microsoft.com/zh-cn/documentation/templates/101-webapp-with-golang/)
    - [Qiniu](https://www.qiniu.com)
    	- [Qiniu SDK for Go](http://developer.qiniu.com/docs/v6/sdk/go-sdk.html):=>[github](https://github.com/qiniu/api.v6)
- 其它站点：
	- [Golang 杂志](https://flipboard.com/section/the-golang-magazine-bJ1GqB)：[订阅说明](http://bbs.go-china.org/post/476)
	- [Reddit](http://www.reddit.com/r/golang/)
	- [Newspaper.IO](http://www.newspaper.io/golang)：Golang 新闻订阅
	- [Go Newsletter](http://www.golangweekly.com/)：Golang 新闻订阅

## 资料汇总

### 中文资料

- 书籍：
	- [《深入解析Go》](https://github.com/tiancaiamao/go-internals)
	- [《Go实战开发》](https://github.com/astaxie/Go-in-Action)
	- [《Go入门指南》](https://github.com/Unknwon/the-way-to-go_ZH_CN)
	- [《Go语言标准库》](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
	- [《Go Web 编程》](https://github.com/astaxie/build-web-application-with-golang)
	- [《Go语言博客实践》](https://github.com/achun/Go-Blog-In-Action)
	- [《Go语言学习笔记》](https://github.com/qyuhen/book)
    - [Go语言中文文档](http://docscn.studygolang.com/doc/)
    - [Go语言圣经中文版](https://docs.hacknode.org/gopl-zh/index.html)
- 翻译：
	- [Effective Go](https://golang.org/doc/effective_go.html) 英文版
	- [The Way to Go](https://github.com/Unknwon/the-way-to-go_ZH_CN) 中文版
	- [《Learning Go》](https://github.com/miekg/gobook)英文版:=>[《Learning Go》](https://github.com/mikespook/Learning-Go-zh-cn) 中文版
- 教程：
	- [《Go编程基础》](https://github.com/Unknwon/go-fundamental-programming)
	- [《Go Web基础》](https://github.com/Unknwon/go-web-foundation)
	- [《Go名库讲解》](https://github.com/Unknwon/go-rock-libraries-showcases)
	- [Go 命令教程](https://github.com/hyper-carrot/go_command_tutorial)

### 英文资料

- 文档：
	- [Go Code Review Comments](https://code.google.com/p/go-wiki/wiki/CodeReviewComments)：Go 语言代码风格指导
	- [Go Code Convention](https://github.com/Unknwon/go-code-convention)：无闻的 Go 语言编码规范
	- [GopherCon 2014](https://github.com/gophercon/2014-talks)
	- [GopherCon 2015](https://github.com/gophercon/2015-talks)
	- [GopherCon 2016](https://github.com/gophercon/2016-talks)
	- [GopherCon 2017](https://github.com/gophercon/2017-talks)
- 书籍：
	- [Network programming with Go](http://jan.newmarch.name/go/)：[中文版](https://github.com/astaxie/NPWG_zh)
	- [Practical Cryptography With Go](https://leanpub.com/gocrypto/read#leanpub-auto-select-bibliography)
	- [An Introduction to Programming in Go](http://www.golang-book.com/)
	- [Go Bootcamp](http://www.golangbootcamp.com/book)
	- [Mastering Concurrency in Go(July 2014)Nathan Kozyra](https://www.packtpub.com/application-development/mastering-concurrency-go)
	- [Go Programming Blueprints(January 23,2015)](https://www.packtpub.com/application-development/go-programming-blueprints)
	- [The Go Programming Language(Published Oct 30, 2015,Not Yet Published)](http://www.gopl.io/)

### 视频资料

- 基础：
	- Go Slices and Bytes - Shakeel Mahate：[优酷视频](http://v.youku.com/v_show/id_XNjkzMjM1Mjg4.html) - [Youtube](http://www.youtube.com/watch?v=dKlNSIUSfz0)
- COSCUP 2013:
	- Golang & ORM - 林佑安：[优酷视频](http://v.youku.com/v_show/id_XNjkzMTQ1MjYw.html) - [Youtube](http://www.youtube.com/watch?v=VwAtYGyjTks)
- GopherCon：
	- 2014：[Youtube](https://www.youtube.com/playlist?list=PL2ntRZ1ySWBcD_BiJiDJUcyrb2w3bTulF)
	- 2015：[Youtube](https://www.youtube.com/playlist?list=PL2ntRZ1ySWBf-_z-gHCOR2N156Nw930Hm)
	- 2016：[Youtube](https://www.youtube.com/watch?v=KINIAgRpkDA&list=PL2ntRZ1ySWBdliXelGAItjzTMxy2WQh0P)
	- 2017：[Youtube](https://www.youtube.com/watch?v=ha8gdZ27wMo&list=PL2ntRZ1ySWBdD9bru6IR-_WXUgJqvrtx9)
- Golang UK Conference：
	- 2015：[Youtube](https://www.youtube.com/playlist?list=PLDWZ5uzn69ezRJYeWxYNRMYebvf8DerHd)
- GopherCon India 2015: [Youtube](https://www.youtube.com/playlist?list=PLxFC1MYuNgJTY3uQ5Ja4F5Sz305nnrBOq)





