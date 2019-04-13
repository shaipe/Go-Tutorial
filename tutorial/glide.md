# GLIDE go语言依赖管理


### 安装

```bash
go get -v github.com/Masterminds/glide
go install github.com/Masterminds/glide
```

初始化

如果是未使用 glide 的工程直接

glide init

在初始化过程中， glide 会询问一些问题
一般建议 y,在升级策略是建议 只下载补丁(patch)不下载升级(minor)
完成后会生成文件 glide.yaml
glide.yaml记载了依赖包的列表及其更新规则

注意这个文件必须 utf-8 编码

glide 版本规则
=: equal (aliased to no operator)
!=: not equal
>: greater than
<: less than
>=: greater than or equal to
<=: less than or equal to

1.2 - 1.4.5 which is equivalent to >= 1.2, <= 1.4.5
2.3.4 - 4.5 which is equivalent to >= 2.3.4, <= 4.5
1.2.x is equivalent to >= 1.2.0, < 1.3.0
>= 1.2.x is equivalent to >= 1.2.0
<= 2.x is equivalent to < 3
* is equivalent to >= 0.0.0
~1.2.3 is equivalent to >= 1.2.3, < 1.3.0
~1 is equivalent to >= 1, < 2
~2.3 is equivalent to >= 2.3, < 2.4
~1.2.x is equivalent to >= 1.2.0, < 1.3.0
~1.x is equivalent to >= 1, < 2
^1.2.3 is equivalent to >= 1.2.3, < 2.0.0
^1.2.x is equivalent to >= 1.2.0, < 2.0.0
^2.3 is equivalent to >= 2.3, < 3
^2.x is equivalent to >= 2.0.0, < 3

安装依赖
如果是已经使用 gilde 管理的工程，本身包含 glide.yaml 文件
glide install

安装后会生产文件 glide.lock
这个文件记载了依赖包确定的revision, 下次再执行 glide install 时，会直接读这个文件下载确定的版本
升级依赖
glide up 会按照语义化版本规则更新依赖包代码
编辑文件 glide.yaml 然后执行
glide up

获取依赖
glide get github.com/orcaman/concurrent-ma

使用镜像
可以通过配置将墙了的版本库 URL 映射到没被墙的 URL，甚至也可以映射到本地版本库
glide mirror set golang.org/x/crypto github.com/golang/crypto
glide mirror set golang.org/x/sys github.com/golang/sys

这样就将这个库从当前工程的
- package: golang.org/x/crypto

修改结果 $HOME/.glide/mirrors.yaml
repos:
- original: golang.org/x/crypto
  repo: github.com/golang/crypto

类似的
glide mirror set golang.org/x/text /home/users/qiangmzsx/var/golang/golang.org/x/text

移除镜像
glide mirror remove golang.org/x/crypto

全局配置
GLIDE_HOME变量，
该变量就是$HOME/.glide
这个目录之前有提到过，除了包含有mirrors.yaml还有一个很重要的目录cache本地 cache
每次更新代码时，glide 都会在本地保存 cache，以备下次 glide install 使用
GLIDE_HOME可以通过如下命令修改
glide --home [path]

问题
镜像功能有时候不好用，可以用 gopm 安装到全局来解决
Error scanning golang.org/x/sys/unix: cannot find package
[WARN]    Unable to checkout golang.org/x/sys/unix
[WARN]    Unable to set version on golang.org/x/sys/unix to . Err: Cannot detect VCS
[ERROR]    Error scanning golang.org/x/sys/unix: cannot find package "." in:
    /Users/sinlov/.glide/cache/src/https-golang.org-x-sys-unix
[ERROR]    Failed to retrieve a list of dependencies: Error resolving imports


fix

cat $HOME/.glide/mirrors.yaml
repos:
- original: https://golang.org/x/crypto
  repo: https://github.com/golang/crypto
- original: https://golang.org/x/crypto/acme/autocert
  repo: https://github.com/golang/crypto
  base: golang.org/x/crypto
- original: https://golang.org/x/sys/unix
  repo: https://github.com/golang/sys
  base: golang.org/x/sys
- original: https://golang.org/x/net
  repo: https://github.com/golang/net
- original: https://golang.org/x/sync
  repo: https://github.com/golang/sync
- original: https://golang.org/x/tools
  repo: https://github.com/golang/tools
- original: https://golang.org/x/grpc
  repo: https://github.com/golang/grpc
- original: https://golang.org/x/time
  repo: https://github.com/golang/time

glide mirror set https://golang.org/x/sys/unix https://github.com/golang/sys --base golang.org/x/sys

see https://github.com/xkeyideal/glide#bug-2



glide install golang.org 失败
背景
因为golang.org被墙的原因，所以国内安装其中的库会失败（Cannot detect VCS）。


解决方案一
按照如下方式配置镜像：

$ rm -rf ~/.glide
$ mkdir -p ~/.glide
$ glide mirror set https://golang.org/x/mobile https://github.com/golang/mobile --vcs git
$ glide mirror set https://golang.org/x/crypto https://github.com/golang/crypto --vcs git
$ glide mirror set https://golang.org/x/net https://github.com/golang/net --vcs git
$ glide mirror set https://golang.org/x/tools https://github.com/golang/tools --vcs git
$ glide mirror set https://golang.org/x/text https://github.com/golang/text --vcs git
$ glide mirror set https://golang.org/x/image https://github.com/golang/image --vcs git
$ glide mirror set https://golang.org/x/sys https://github.com/golang/sys --vcs git

然后在项目中执行
$ glide init
$ glide install

解决方案二
不使用glide,直接go get下载github中对应的库，拷贝到gopath中去，然后修改gopath/src中的路经，使其与代码中的一致。


