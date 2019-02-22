go数据库操作

===

### Mongo

[官方资料文件](https://godoc.org/github.com/mongodb/mongo-go-driver/mongo#pkg-examples)

#### 1. 安装包
```bash
go get github.com/mongodb/mongo-go-driver
```

#### 2. Create the wireframe

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/mongodb/mongo-go-driver/bson"
    "github.com/mongodb/mongo-go-driver/mongo"
    "github.com/mongodb/mongo-go-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
    Name string
    Age  int
    City string
}

func main() {
    // Rest of the code will go here
}
```

#### 3. Connect to MongoDB using the Go Driver

```go
client, err := mongo.Connect(context.TODO(), "mongodb://localhost:27017")

if err != nil {
    log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
    log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")

# 获取一个文档集
collection := client.Database(demo).Collection("trainers")

# 断开数据库连接
err = client.Disconnect(context.TODO())

if err != nil {
    log.Fatal(err)
}
fmt.Println("Connection to MongoDB closed.")
```