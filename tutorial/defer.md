Defer 的理解 
===

Golang中的defer关键字实现比较特殊的功能，按照官方的解释，defer后面的表达式会被放入一个列表中，在当前方法返回的时候，列表中的表达式就会被执行。一个方法中可以在一个或者多个地方使用defer表达式，这也是前面提到的，为什么需要用一个列表来保存这些表达式。在Golang中，defer表达式通常用来处理一些清理和释放资源的操作。


### defer 的作用和执行时机

go 的 defer 语句是用来延迟执行函数的，而且延迟发生在调用函数 return 之后，比如

```go
func a() int {
  defer b()
  return 0
}
```

b 的执行是发生在 return 0 之后，注意 defer 的语法，关键字 defer 之后是函数的调用。

defer 的重要用途一：清理释放资源
由于 defer 的延迟特性，defer 常用在函数调用结束之后清理相关的资源，比如
```go

f, _ := os.Open(filename)
defer f.Close()

```

文件资源的释放会在函数调用结束之后借助 defer 自动执行，不需要时刻记住哪里的资源需要释放，打开和释放必须相对应。

用一个例子深刻诠释一下 defer 带来的便利和简洁。


示例
```go
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }

    written, err = io.Copy(dst, src)
    dst.Close()
    src.Close()
    return
}

```

CopyFile方法简单的实现了文件内容的拷贝功能，将源文件的内容拷贝到目标文件。咋一看还没什么问题，不过Golang中的资源也是需要释放的，假如os.Create方法的调用出了错误，下面的语句会直接return，导致这两个打开的文件没有机会被释放。这个时候，defer就可以派上用场了。

使用defer改进过后的例子：

```go
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```


改进的代码中两处都使用到了defer表达式，表达式的内容就是关闭文件。前面介绍过，虽然表达式的具体行为是关闭文件，但是并不会被马上执行，两个表达式都会被放入一个list，等待被调用。

### defer特性

#### defer表达式中变量的值在defer表达式被定义时就已经明确

```go
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```



上面的这段代码，defer表达式中用到了i这个变量，i在初始化之后的值为0，接着程序执行到defer表达式这一行，表达式所用到的i的值就为0了，接着，表达式被放入list，等待在return的时候被调用。所以，后面尽管有一个i++语句，仍然不能改变表达式 fmt.Println(i)的结果。

所以，程序运行结束的时候，输出的结果是0而不是1。

#### defer表达式的调用顺序是按照先进后出的方式
```go
func b() {
    defer fmt.Print(1)
    defer fmt.Print(2)
    defer fmt.Print(3)
    defer fmt.Print(4)
}
```
前面已经提到过，defer表达式会被放入一个类似于栈(stack)的结构，所以调用的顺序是后进先出的。所以，上面这段代码输出的结果是4321而不是1234。在实际的编码中应该主意，程序后面的defer表达式会被优先执行。

#### defer表达式中可以修改函数中的命名返回值
Golang中的函数返回值是可以命名的，这也是Golang带给开发人员的一个比较方便特性。
```go

func c() (i int) {
    defer func() { i++ }()
    return 1
}

```
上面的示例程序，返回值变量名为i，在defer表达式中可以修改这个变量的值。所以，虽然在return的时候给返回值赋值为1，后来defer修改了这个值，让i自增了1，所以，函数的返回值是2而不是1。


### defer 表达式的使用场景
defer 通常用于 open/close, connect/disconnect, lock/unlock 等这些成对的操作, 来保证在任何情况下资源都被正确释放. 在这个角度来说, defer 操作和 Java 中的 try ... finally 语句块有异曲同工之处.
例如:

```go
var mutex sync.Mutex
var count = 0

func increment() {
    mutex.Lock()
    defer mutex.Unlock()
    count++
}
```
在increment 函数中, 我们为了避免竞态条件的出现, 而使用了 Mutex 进行加锁. 而在进行并发编程时, 加锁了却忘记(或某种情况下 unlock 没有被执行), 往往会造成灾难性的后果. 为了在任意情况下, 都要保证在加锁操作后, 都进行对应的解锁操作, 我们可以使用 defer 调用解锁操作.


### defer 的作用域

- defer 只对当前协程有效（main 可以看作是主协程）；
- 当任意一条（主）协程发生 panic 时，会执行当前协程中 panic 之前已声明的 defer；
- 在发生 panic 的（主）协程中，如果没有一个 defer 调用 recover()进行恢复，则会在执行完最后一个已声明的 defer 后，引发整个进程崩溃；
- 主动调用 os.Exit(int) 退出进程时，defer 将不再被执行。

```go
package mainimport (
        "errors"
        "fmt"
        "time"
        // "os")func main() {
        e := errors.New("error")
        fmt.Println(e)
        // （3）panic(e) // defer 不会执行
        // （4）os.Exit(1) // defer 不会执行
        defer fmt.Println("defer")
        // （1）go func() { panic(e) }() // 会导致 defer 不会执行
        // （2）panic(e) // defer 会执行
        time.Sleep(1e9)
        fmt.Println("over.")
        // （5）os.Exit(1) // defer 不会执行}
```

### 参考资料

[defer-panic-and-recover](https://blog.golang.org/defer-panic-and-recover)
