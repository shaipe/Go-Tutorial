Go语言编码规范
===

#### 项目目录结构规范

```
PROJECT_NAME
├── README.md 介绍软件及文档入口
├── bin 编译好的二进制文件,执行./build.sh自动生成，该目录也用于程序打包
├── build.sh 自动编译的脚本
├── doc 该项目的文档
├── pack 打包后的程序放在此处
├── pack.sh 自动打包的脚本，生成类似xxxx.20170713_14:45:35.tar.gz的文件，放在pack文件下
└── src 该项目的源代码
    ├── main 项目主函数
    ├── model 项目代码
    ├── research 在实现该项目中探究的一些程序
    └── vendor 存放go的库
        ├── github.com/xxx 第三方库
        └── xxx.com/obc 公司内部的公共库
```

项目的目录结构尽量做到简明、层次清楚

#### 文件名命名规范
用小写，尽量见名思义，看见文件名就可以知道这个文件下的大概内容，对于源代码里的文件，文件名要很好的代表了一个模块实现的功能。

### 命名规范
#### 包名
包名用小写,使用短命名,尽量和标准库不要冲突

#### 接口名
单个函数的接口名以”er”作为后缀，如Reader,Writer

接口的实现则去掉“er”
```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```
两个函数的接口名综合两个函数名
```go
type WriteFlusher interface {
    Write([]byte) (int, error)
    Flush() error
}
```
三个以上函数的接口名，类似于结构体名

```go
type Car interface {
    Start([]byte) 
    Stop() error
    Recover()
}
```
#### 变量
全局变量：采用驼峰命名法，仅限在包内的全局变量，包外引用需要写接口，提供调用 局部变量：驼峰式，小写字母开头

#### 常量
常量：大写，采用下划线

#### import 规范
import在多行的情况下，goimports会自动帮你格式化，在一个文件里面引入了一个package，建议采用如下格式：
```go
import (
    "fmt"
)
```
如果你的包引入了三种类型的包，标准库包，程序内部包，第三方包，建议采用如下方式进行组织你的包：
```go
import (
    "encoding/json"
    "strings"

    "myproject/models"
    "myproject/controller"
    "git.obc.im/obc/utils"

    "git.obc.im/dep/beego"
    "git.obc.im/dep/mysql"
)  
```
在项目中不要使用相对路径引入包：
```go
// 这是不好的导入

import “../net”
// 这是正确的做法

import “xxxx.com/proj/net”
```
#### 函数名
函数名采用驼峰命名法，尽量不要使用下划线

#### 错误处理
error作为函数的值返回,必须尽快对error进行处理
采用独立的错误流进行处理
不要采用这种方式
```go
    if err != nil {
        // error handling
    } else {
        // normal code
    }
```
而要采用下面的方式
```go
    if err != nil {
        // error handling
        return // or continue, etc.
    }
    // normal code
```
如果返回值需要初始化，则采用下面的方式
```go
x, err := f()
if err != nil {
    // error handling
    return
}
// use x
```
#### Panic
在逻辑处理中禁用panic
在main包中只有当实在不可运行的情况采用panic，例如文件无法打开，数据库无法连接导致程序无法 正常运行，但是对于其他的package对外的接口不能有panic，只能在包内采用。 建议在main包中使用log.Fatal来记录错误，这样就可以由log来结束程序。

#### Recover
recover用于捕获runtime的异常，禁止滥用recover，在开发测试阶段尽量不要用recover，recover一般放在你认为会有不可预期的异常的地方。
```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    // do 函数可能会有不可预期的异常
    do(work)
}
```
#### Defer
defer在函数return之前执行，对于一些资源的回收用defer是好的，但也禁止滥用defer，defer是需要消耗性能的,所以频繁调用的函数尽量不要使用defer。
```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```
### 控制结构
#### if
if接受初始化语句，约定如下方式建立局部变量
```go
if err := file.Chmod(0664); err != nil {
    return err
}
```
#### for
采用短声明建立局部变量
```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```
#### range
如果只需要第一项（key），就丢弃第二个：
```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```
如果只需要第二项，则把第一项置为下划线
```go
sum := 0
for _, value := range array {
    sum += value
}
```
#### return
尽早return：一旦有错误发生，马上返回
```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```
#### 方法的接收器
名称 一般采用strcut的第一个字母且为小写，而不是this，me或者self
```go
    type T struct{} 
    func (p *T)Get(){}
```
如果接收者是map,slice或者chan，不要用指针传递
```go
//Map
package main

import (
    "fmt"
)

type mp map[string]string

func (m mp) Set(k, v string) {
    m[k] = v
}

func main() {
    m := make(mp)
    m.Set("k", "v")
    fmt.Println(m)
}
//Channel
package main

import (
    "fmt"
)

type ch chan interface{}

func (c ch) Push(i interface{}) {
    c <- i
}

func (c ch) Pop() interface{} {
    return <-c
}

func main() {
    c := make(ch, 1)
    c.Push("i")
    fmt.Println(c.Pop())
}
```
如果需要对slice进行修改，通过返回值的方式重新赋值
```go
//Slice
package main

import (
    "fmt"
)

type slice []byte

func main() {
    s := make(slice, 0)
    s = s.addOne(42)
    fmt.Println(s)
}

func (s slice) addOne(b byte) []byte {
    return append(s, b)
}
```
如果接收者是含有sync.Mutex或者类似同步字段的结构体，必须使用指针传递避免复制
```go
package main

import (
    "sync"
)

type T struct {
    m sync.Mutex
}

func (t *T) lock() {
    t.m.Lock()
}

/*
Wrong !!!
func (t T) lock() {
    t.m.Lock()
}
*/

func main() {
    t := new(T)
    t.lock()
}
```
如果接收者是大的结构体或者数组，使用指针传递会更有效率。
```go
package main

import (
    "fmt"
)

type T struct {
    data [1024]byte
}

func (t *T) Get() byte {
    return t.data[0]
}

func main() {
    t := new(T)
    fmt.Println(t.Get())
}
```


优秀的命名
优秀的命名应当是一贯的、短小的、精确的。
所谓一贯，就是说同一个意义在不同的环境下的命名应当一致，譬如依赖关系，不要在一个方法中命名为depend，另一个方法中命名为rely。
所谓短小，不必多言，当命名过长的时候，读者可能更关注命名本身，而忽视真正的逻辑内容。
所谓精确，就是命名达意、易于理解

首条经验
声明位置与使用位置越远，则命名应当越长。

骆驼命名法
Go语言应该使用 MixedCase
(不要使用 names_with_underscores)
首字母缩写词都应该用大写,譬如ServeHTTP、sceneID、CIDRProcessor。

局部变量
局部变量应当尽可能短小，譬如使用buf指代buffer，使用i指代index
在很长的函数中可能会有很多的变量，这个时候可以适当使用一些长名字。
但是写出这么长的函数，通常意味着代码需要重构了！🙅🏻‍

参数
函数的参数和局部变量类似，但是它们默认还具有文档的功能
当参数类型具有描述性的时候，参数名就应该尽可能短小：

func AfterFunc(d Duration, f func()) *Timer
func Escape(w io.Writer, s []byte)
当参数类型比较模糊的时候，参数名就应当具有文档的功能：

func Unix(sec, nsec int64) Time
func HasPrefix(s, prefix []byte) bool
返回值
在Go语言中，返回值可以定义名称的，它可以当做一种特殊的参数。
尤其重要的是，在外部可见的函数中，返回值的名称应当可以作为文档参考。

func Copy(dst Writer, src Reader) (written int64, err error)
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte,
 err error)
方法接收者（Receiver）
方法接收者也是一种特殊的参数。Go语言中没有明显的面向对象的概念，可以对方法定义方法接收者来实现类似于对象的方法的概念。

按照惯例，由于方法接收者在函数内部经常出现，因此它经常采用一两个字母来标识方法接收者的类型。

func (b *Buffer) Read(p []byte) (n int, err error)
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request)
func (r Rectangle) Size() Point
需要注意的是，方法接收者的名字在同一类型的不同方法中应该保持统一，这也是前文所述的一贯性的需求。

导出包级别命名
导出名被使用的时候通常是放在包名后
所以，在导出变量、常数、函数和类型的时候，
不要把包名的意义再写一遍

比较好的名字
bytes.Buffer strings.Reader

比较蠢的名字
bytes.ByteBuffer strings.StringReader

接口类型
只含有一个方法的接口类型通常以函数名加上er后缀作为名字

type Reader interface {
    Read(p []byte) (n int, err error)
}
有时候可能导致蹩脚的英文，但别管他，能看懂就好

type Execer interface {
    Exec(p []byte) (n int, err error)
}
有时候可以适当调整一下英文单词的顺序，增加可读性：

type ByteReader interface {
    ReadByte(p []byte) (n int, err error)
}
当接口含有多个方法的时候，还是要选取一个能够精准描述接口目的的名字，譬如net.Conn、http/ResponseWriter

Error的命名
Error类型应该写成FooError的形式

type ExitError struct {
	....
}
Error变量协程ErrFoo的形式

var ErrFormat = errors.New("unknown format")
包的命名
应当与它导出代码的内容相关，避免util、common这种宽泛的命名

引入路径
包路径的最后一个单词应该和包名一致

包路径应该尽可能简洁

记得把库的主要代码直接放在代码库的根目录

避免在包路径中使用任何大写字母（并非所有文件系统都区分大小写）

标准库
上述很多例子都是从标准库中来的

标准库的很多内容都可以作为参考
多看看标准库来寻求灵感吧

但是要记住：

当作者写标准库的时候，他们自己也在学习过程中。
多数情况下作者是对的，但是偶尔还是会犯一些错误


