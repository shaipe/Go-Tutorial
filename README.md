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

## 相关学习资料

- [Go语言中文文档](http://docscn.studygolang.com/doc/)
- [Go语言圣经中文版](https://docs.hacknode.org/gopl-zh/index.html)

