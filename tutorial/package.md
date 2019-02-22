Go 语言安装包
===


package遵循以下原则：
1. package是最基本的分发单位和工程管理中依赖关系的体现
2. 每个Go语言源代码文件开头都必须要有一个package声明，表示源代码文件所属包
3. 要生成Go语言可执行程序，必须要有名为main的package包，且在该包下必须有且只有一个main函数
4. 同一个路径下只能存在一个package，一个package可以由多个源代码文件组成

#### 默认安装




```bash
go get 托管网站/用户名/包名

# 示例
go get github.com/PuerkitoBio/goquery
```

#### 手动安装

由于 golang.org 在国内已被墙，不能直接使用go get安装对应的包，但是可以通过github间接安装。

```bash

cd $GOPATH/src
mkdir -p golang.org/x/
cd  golang.org/x/
git clone https://github.com/golang/crypto.git


```


### import

跟package类似，import原理遵守以下几个原则：

如果一个main导入其他的包，包会被顺序导入
如果导入的包（pkg1）依赖其他的包（包pkg2），会首先导入pkg2，然后初始化pkg2中的常量与变量，如果pkg2中有init函数，会自动执行init
所有包导入完成后才会对main的常量和变量进行初始化，然后执行main中的init函数（如果有的话），最后执行main函数
如果一个包被导入多次实际上只会导入一次

import 别名

import xxx "github.com/dd/xxx"



1 概述
Go 语言的源码复用建立在包（package）基础之上。包通过 package, import, GOPATH 操作完成。

2 main包
Go 语言的入口 main() 函数所在的包（package）叫 main，main 包想要引用别的代码，需要import导入！

3 包定义，package
Go 语言的包与文件夹一一对应，同一个目录下的所有.go文件的第一行添加 包定义，以标记该文件归属的包，演示语法：

package 包名
包需要满足：

一个目录下的同级文件归属一个包。
包名可以与其目录不同名。
包名为 main 的包为应用程序的入口包，其他包不能使用。
包可以嵌套定义，对应的就是嵌套目录，但包名应该与所在的目录一致，例如：

// 文件：foo/bar/tool.go中
package bar
// 可以被导出的函数
func FuncPublic() {
}
// 不可以被导出的函数
func funcPrivate() {
}
包中，通过标识符首字母是否大写，来确定是否可以被导出。首字母大写才可以被导出，视为 public 公共的资源。

4 导入包，import
要引用其他包，可以使用 import 关键字，可以单个导入或者批量导入，语法演示：

// 单个导入
import "package"
// 批量导入
import (
  "package1"
  "package2"
  )
导入时，可以为包定义别名，语法演示：

import (
  p1 "package1"
  p2 "package2"
  )
// 使用时
p1.Method()
以上测试请使用系统包测试。若需要导入自定义包，需要设置GOPATH环境变量。

4 GOPATH环境变量
import导入时，会从GO的安装目录（也就是GOROOT环境变量设置的目录）和GOPATH环境变量设置的目录中，检索 src/package 来导入包。如果不存在，则导入失败。 GOROOT，就是GO内置的包所在的位置。 GOPATH，就是我们自己定义的包的位置。

通常我们在开发Go项目时，调试或者编译构建时，需要设置GOPATH指向我们的项目目录，目录中的src目录中的包就可以被导入了： 例如，我么的项目目录为： D:\projects\goProject，那么我么就需要将我们的源代码放在 D:\projects\goProject\src 下，同时设置GOPATH为 D:\projects\goProject。设置GOPATH的方案有：

windows 通过 系统->系统信息->高级系统设置->环境变量 中完成设置。
windows 中通过 CMD 或者 powershell 也可以完成设置。通常是临时有效的，CMD或者powershell关闭失效！
CMD：
set GOPATH=D:\projects\goProject
set GOPATH 可以查看
powershell：
$env:GOPATH="D:\projects\goProject"
$env:GOPATH 可以查看
linux 通过 /etc/profile 进行设置
5 init() 包初始化
可以在源码中，定义 init() 函数。此函数会在包被导入时执行，例如如果是在 main 中导入包，包中存在 init()，那么 init() 中的代码会在 main() 函数执行前执行，用于初始化包所需要的特定资料。例如： 包源码：

src/userPackage/tool.go

package userPackage
import "fmt"
func init() {
  fmt.Println("tool init")
}
主函数源码：

src/main.go

package main
import (
  "userPackage"
  )
func main() {
  fmt.Println("main run")
  // 使用userPackage
  userPackage.SomeFunc()
}
执行时，会先输出 "tool init"，再输出 "main run"。

如果仅仅需要导入包时执行初始化操作，并不需要使用包内的其他函数，常量等资源。则可以在导入包时，匿名导入：

import (
  _ "userPackage"
  )
使用下划线作为包的别名，会仅仅执行init()