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

