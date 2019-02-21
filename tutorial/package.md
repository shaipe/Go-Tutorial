Go 语言安装包
===

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