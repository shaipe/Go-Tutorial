golang本身没有提供连接mysql的驱动，但是定义了标准接口供第三方开发驱动。这里连接mysql可以使用第三方库，第三方库推荐使用https://github.com/Go-SQL-Driver/MySQL这个驱动，更新维护都比较好。下面演示下具体的使用，完整代码示例可以参考最后。

下载驱动
```bash
sudo go get github.com/go-sql-driver/mysql
```

如果提示这样的失败信息：cannot download, $GOPATH not set. For more details see: go help gopath，可以使用如下命令解决
```bash

sudo env GOPATH=/Users/chenjiebin/golang go get github.com/go-sql-driver/mysql

```
GOPATH的值根据自行环境进行替换。

创建测试表
在mysql test库中创建测试表
```sql

CREATE TABLE IF NOT EXISTS `test`.`user` (
 `user_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户编号',
 `user_name` VARCHAR(45) NOT NULL COMMENT '用户名称',
 `user_age` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户年龄',
 `user_sex` TINYINT(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户性别',
 PRIMARY KEY (`user_id`))
 ENGINE = InnoDB
 AUTO_INCREMENT = 1
 DEFAULT CHARACTER SET = utf8
 COLLATE = utf8_general_ci
 COMMENT = '用户表'
 
```
数据库连接
数据库连接使用datebase/sql Open函数进行连接

```go

db, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?charset=utf8")

```
其中连接参数可以有如下几种形式：

user@unix(/path/to/socket)/dbname?charset=utf8
user:password@tcp(localhost:5555)/dbname?charset=utf8
user:password@/dbname
user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname

通常我们都用第二种。

插入操作
```go


stmt, err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
checkErr(err)
res, err := stmt.Exec("tony", 20, 1)
checkErr(err)
id, err := res.LastInsertId()
checkErr(err)
fmt.Println(id)
```
这里使用结构化操作，不推荐使用直接拼接sql语句的方法。

查询操作
```go


rows, err := db.Query("SELECT * FROM user")
checkErr(err)
 
for rows.Next() {
    var userId int
    var userName string
    var userAge int
    var userSex int
    rows.Columns()
    err = rows.Scan(&userId, &userName, &userAge, &userSex)
    checkErr(err)
    fmt.Println(userId)
    fmt.Println(userName)
    fmt.Println(userAge)
    fmt.Println(userSex)
}
```
这里查询的方式使用声明4个独立变量userId、userName、userAge、userSex来保存查询出来的每一行的值。在实际开发中通常会封装数据库的操作，对这样的查询通常会考虑返回字典类型。
```go


//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
columns, _ := rows.Columns()
scanArgs := make([]interface{}, len(columns))
values := make([]interface{}, len(columns))
for i := range values {
    scanArgs[i] = &values[i]
}
 
for rows.Next() {
    //将行数据保存到record字典
    err = rows.Scan(scanArgs...)
    record := make(map[string]string)
    for i, col := range values {
        if col != nil {
            record[columns[i]] = string(col.([]byte))
        }
    }
    fmt.Println(record)
}
```

修改操作
```go
stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
checkErr(err)
res, err := stmt.Exec(21, 2, 1)
checkErr(err)
num, err := res.RowsAffected()
checkErr(err)
fmt.Println(num)
```

删除操作
```go

stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
checkErr(err)
res, err := stmt.Exec(1)
checkErr(err)
num, err := res.RowsAffected()
checkErr(err)
fmt.Println(num)
````
修改和删除操作都比较简单，同插入数据类似，只是使用RowsAffected来获取影响的数据行数。

完整代码
```go


package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)
 
func main() {
    insert()
}
 
//插入demo
func insert() {
    db, err := sql.Open("mysql", "root:@/test?charset=utf8")
    checkErr(err)
 
    stmt, err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
    checkErr(err)
    res, err := stmt.Exec("tony", 20, 1)
    checkErr(err)
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(id)
}
 
//查询demo
func query() {
    db, err := sql.Open("mysql", "root:@/test?charset=utf8")
    checkErr(err)
 
    rows, err := db.Query("SELECT * FROM user")
    checkErr(err)
 
    //普通demo
    //for rows.Next() {
    //    var userId int
    //    var userName string
    //    var userAge int
    //    var userSex int
 
    //    rows.Columns()
    //    err = rows.Scan(&userId, &userName, &userAge, &userSex)
    //    checkErr(err)
 
    //    fmt.Println(userId)
    //    fmt.Println(userName)
    //    fmt.Println(userAge)
    //    fmt.Println(userSex)
    //}
 
    //字典类型
    //构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
    columns, _ := rows.Columns()
    scanArgs := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
    for i := range values {
        scanArgs[i] = &values[i]
    }
 
    for rows.Next() {
        //将行数据保存到record字典
        err = rows.Scan(scanArgs...)
        record := make(map[string]string)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
        fmt.Println(record)
    }
}
 
//更新数据
func update() {
    db, err := sql.Open("mysql", "root:@/test?charset=utf8")
    checkErr(err)
 
    stmt, err := db.Prepare(`UPDATE user SET user_age=?,user_sex=? WHERE user_id=?`)
    checkErr(err)
    res, err := stmt.Exec(21, 2, 1)
    checkErr(err)
    num, err := res.RowsAffected()
    checkErr(err)
    fmt.Println(num)
}
 
//删除数据
func remove() {
    db, err := sql.Open("mysql", "root:@/test?charset=utf8")
    checkErr(err)
 
    stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
    checkErr(err)
    res, err := stmt.Exec(1)
    checkErr(err)
    num, err := res.RowsAffected()
    checkErr(err)
    fmt.Println(num)
}
 
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
```
小结
整体上来说都比较简单，就是查询那边使用字典来存储返回数据比较复杂一些。既然说到数据库连接，通常应用中都会使用连接池来减少连接开销，关于连接池下次整理一下再放上来。