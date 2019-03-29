go语言操作Redis
===

在go语言中常使用的是go-redis


## 接口：github.com/go-redis/redis 中的 redis.client
#### 1、AllKeys 
入参: 无
出参：1、[]string 实例对应库中所有的Key值
     2、redis.Error
返回对应库中所有的key（等于keys *）


#### 2、Bgsave
入参：无
出参：redis.Error
后台保存rdb快照，执行后dump.rdb文件立即更新


#### 3、Blpop
入参：keys string, timeout int
出参：1、[][]byte 弹出的结果 
      2、redis.Error 
#### 链表左边等待弹出key的尾/头元素，timeout为等待超时时间，如果timeout为0则一直等待下去


#### 4、Brpop
入参：keys string, timeout int
出参：1、[][]byte 弹出的结果 
      2、redis.Error 
链表右边等待弹出key的尾/头元素，timeout为等待超时时间，如果timeout为0则一直等待下去


#### 5、Brpoplpush
入参：src string, desc string, timeout int
出参：1、[][]byte 弹出插入的结果
     2、redis.Error
功能：从src链表右边弹出数据，在desc左边插入数据。在windows下测试发现功能ok，已经使redis完成对应功能，但是返回出错：SYSTEM_ERROR - ServiceRequest [cause: connHdl.ServiceRequest(BRPOPLPUSH) - failed to get response]测试的时候timeout已经设置为0了


#### 6、Dbsize
入参：无
出参：result int64, redis.Error
功能：返回redis实例对应库中key的个数


#### 7、Decr
入参：key string 对应string类型的key
出参：1、result int64 /*key值对应的value值减1后的结果*/
      2、redis.Error
功能：实现对key对应的value值的原子减一操作。


#### 8、Decrby
入参: key string 对应string类型的key， arg1 int64
出参：1、result int64 /*key值对应的value值减去arg1后的结果*/
      2、redis.Error
#### 功能：实现对key对应的value值的原子减对应值操作。


#### 9、Del
入参：key string 想要删除的key
出参：1、rusult bool 删除是否成功
      2、redis.Error
功能描述：删除的key


#### 10、Exists
入参：key string 想要检查的key
出参：1、rusult bool 删除是否成功（功能：true，失败：false)
      2、redis.Error
功能描述：检查key值是否存在。


#### 11、Expire
入参: 1、key string 要进行设置的key值
      2、arg1 int64 设置的生命周期 单位为秒
出参：1、rusult bool 是否设置成功（功能：true，失败：false)
      2、redis.Error
功能描述:为key设置生命周期。


#### 12、Flushall （谨用）
入参：无
出参：redis.Error
功能描述：删除所有数据库中的所有的key


#### 13、Flushdb （谨用）
入参：无
出参：redis.Error
功能描述：删除当前库环境中的所有的key


#### 14、Get
入参：key string 要获取的key值
出参: 1、result []byte 获取的结果
      2、redis.Error
功能描述:获取string类型value


#### 15、Getset
入参: 1、key string 需要查找的key值
      2、arg1 []byte 需要将key设置为对应的值
出参：1、result []byte 返回key对应的value值
      2、redis.Error
功能描述：获取key的值，同时设置key的值


#### 16、Hget
入参: 1、key string 查找的key值
      2、hashkey string 需要查找的key里边的field值
出参: 1、result []byte 查找的结果
      2、redis.Error
功能描述:获取哈希类型对应域中的值


#### 17、Hgetall
入参: 1、key string 查找的key值
出参: 1、result [][]byte 查找的结果,结果集为，field1 value1 field2 value2..
      2、redis.Error
功能描述:获取哈希类型所有域中的值


#### 18、Hset
入参：1、key string 需要设置的key值
      2、hashkey string hashkey值
      3、arg1 []byte hashkey要设置的value
出参：redis.Error
功能描述:新建或者设置一个hash类型的key


#### 19、Incr
入参：key string 要增加的key值
出参：1、result int64 增加后的结果
      2、redis.Error
功能描述：原则增加string类型key中value的值（+1）


#### 20、Incrby
入参：1、key string 要增加的key值
      2、arg1 int64 需要增加的值
出参：1、result int64 增加后的结果
      2、redis.Error
功能描述：原子增加string类型key中value的值（+arg1）


#### 21、Info
入参：无
出参：1、result map[string]string 对应的结果信息
      2、redis.Error
功能描述：获取redis实例的信息。测试发现返回一直为空


#### 22、Keys
入参: key string 要检查的key值
出参：1、result []byte 返回结果
      2、redis.Error
功能描述：检查key是否存在


#### 23、Lastsave
入参: 无
出参：1、result int64 redis实例最后保存时间
      2、redis.Error
功能描述：查找redis实例最后保存数据的时间


#### 24、Lindex
入参: 1、key string 对应链表类型的key
      2、arg1 int64 数据的位置
出参：1、result []byte 对应链表位置上的value
      2、redis.Error
功能描述：查找链表类型arg1位置上的value值


#### 25、Llen
入参: 1、key string 对应链表类型的key
     
出参：1、result int64 链表的大小
      2、redis.Error
功能描述：获取链表字段的大小


#### 26、Lpop
入参: key string 对应链表类型的key
出参：1、result []byte 链表的大小
      2、redis.Error
功能描述：返回并删除链表左边的值


#### 27、Lpush
入参: 1、key string 要插入链表类型的key
      2、arg1 []byte 要插入的value值
出参：redis.Error
功能描述：从链表左边插入值


#### 28、Lrange
入参: 1、key string 要查找链表类型的key
      2、arg1 int64 初始位置
      3、arg2 int64 结束位置
出参：1、result [][]byte 结果集
      2、redis.Error
功能描述：查找链表中初始位置到结束位置之间的值，当结束位置为-1时，返回从初始位置以后的值


#### 29、Lrem
入参: 1、key string 要删除链表类型的key
      2、arg1 []byte 要删除的值
      3、arg2 int64 要删除的个数
出参：1、result int64 删除的个数
      2、redis.Error
功能描述：在链表中删除arg2个值为arg1的成员，当arg2个数为0是全部删除，大于0从表头开始找成员，小于0从表尾找成员。


#### 30、Lset
入参: 1、key string 要修改链表类型的key
      2、arg1 []byte 修改后的值
      3、arg2 int64 修改的成员偏移
出参：1、result int64 是否修改成功
      2、redis.Error
功能描述：将链表中arg2位置上成员的值修改为arg1


#### 31、Ltrim
入参: 1、key string 要修改链表类型的key
      2、arg1 int64 开始位置
      3、arg2 int64 结束位置
出参：redis.Error
功能描述：将链表key裁剪为只有从arg1到arg2位置间的值


#### 32、Mget
入参: 1、key string 要获取的字符串类型的key
      2、arg1 []string 要获取的字符串类型的key的集合
出参：1、result [][]byte 返回的结果集合
      2、redis.Error
功能描述：将链表key裁剪为只有从arg1到arg2位置间的值


#### 33、Move
入参: 1、key string 要转移的key
      2、arg1 int64 要转移到的库
出参：1、result bool 是否转移成功
      2、redis.Error
功能描述：将key转移到对应库位置。


#### 34、Ping
入参: 无入参
出参：redis.Error
功能描述：测试客户端和服务端是否连接成功


#### 35、Publish
入参: 1、channel string 发送到对应的频道
      2、message []byte 发送的消息
出参：1、receviercout int64 接受者的数量
      2、redis.Error
功能描述：往频道发送消息，是个阻塞的接口


#### 36、Quit
入参: 无入参
出参：redis.Error
功能描述：退出redis客户端


#### 37、Randomkey
入参: 无入参
出参：1、result string 返回的key
      2、redis.Error
功能描述：随机返回一个key


#### 38、Rename
入参: 1、key string 要重命名的key
      2、arg1 string 重命名后的key
出参：redis.Error
功能描述：给一个key重命名


#### 39、renamenx 
入参: 1、key string 要重命名的key
      2、arg1 string 重命名后的key
出参：1、result bool 是否修改成功
      2、redis.Error
功能描述：给一个key重命名


#### 40、Rpop
入参: key string 对应链表类型的key
出参：1、result []byte 链表的大小
      2、redis.Error
功能描述：返回并删除链尾部的值


#### 41、Rpoplpush
入参: 1、key string pop操作对应链表的Key
      2、arg1 string push操作对应链表的key
出参：1、result []byte 操作的值
      2、redis.Error
功能描述：将key中尾部的值插入到arg1的头部


#### 42、Rpush
入参: 1、key string 要插入链表类型的key
      2、arg1 []byte 要插入的value值
出参：redis.Error
功能描述：从链表尾部插入值


#### 43、Sadd
入参: 1、key string 要插入集合类型的key
      2、arg1 []byte 要插入的value值
出参：redis.Error
功能描述：插入数据或者新建集合


#### 44、Save
入参: 无入参
出参：redis.Error
功能描述：保存rdb快照


#### 45、Setnx
入参: 1、key string 要新建集合类型的key
      2、arg1 []byte 要插入的value值
出参：1、result bool 是否成功
      2、redis.Error
功能描述：新建集合，只有key不存在的时候才成功


#### 46、Sinter
入参: 1、key string 集合类型的key
      2、arg1 []string 集合类型key2、key3
出参：1、result [][]byte 多个集合的交集
      2、redis.Error
功能描述：取多个集合的交集


#### 47、Scard
入参: 1、key string 集合类型的key
出参：1、result int64 集合成员的个数
      2、redis.Error
功能描述：取两个集合的交集


#### 48、Sdiff
入参: 1、key string 集合类型的key
      2、arg1 []string 集合类型key2、key3...
出参：1、result [][]byte key中存在，key2、key3...不存在的值
      2、redis.Error
功能描述：取key中存在，key2、key3...不存在的值


#### 49、Sdiffstore
入参: 1、key string 集合类型的key
      2、arg1 []string 集合类型key2、key3...
出参：redis.Error
功能描述：取key2中存在，key3、key4...不存在的值，并保存在key中


#### 50、Sinterstore
入参: 1、key string 集合类型的key
      2、arg1 []string 集合类型key2、key3...
出参：redis.Error
功能描述：取多个集合的交集，并将结果存储到key中


#### 51、Sismember
入参: 1、key string 集合类型的key
      2、arg1 string 集合内的成员的值
出参：1、result int64 值为arg1的个数
      2、redis.Error
功能描述：获取集合中对应值的个数


#### 52、Smembers
入参: 1、key string 集合类型的key
出参：1、result [][]byte 集合中所有的value
      2、redis.Error
功能描述：获取集合所有值


#### 53、Smove
入参: 1、key string 集合类型的key
      2、arg1 string 集合类型key2
      3、arg2 []byte 需要从key2移动到key的值
出参：1、result bool 移动后是否成功
      2、redis.Error
功能描述：将集合arg1中的值arg2移动到集合key中


#### 54、Srandmember
入参: 1、key string 集合类型的key
出参：1、result []byte 随机返回的结果
      2、redis.Error
功能描述：从集合Key中随机返回一个结果


#### 55、Srem
入参: 1、key string 集合类型的key
      2、arg1 []byte 要删除的值
出参：1、result bool 删除是否成功
      2、redis.Error
功能描述：从集合Key删除一个值


#### 56、Sunion
入参: 1、key string 集合类型的key
      2、arg1 []string 要和key取并集的keys
出参：1、result [][]byte 取并集后的结果
      2、redis.Error
功能描述：去集合的并集


#### 57、Sunionstore
入参: 1、key string 集合类型的key
      2、arg1 []string 要并集的keys
出参：redis.Error
功能描述：arg1取并集后存储在key中


#### 58、Ttl
入参: 1、key string 集合类型的key
出参：1、result int64 集合的生命周期
      2、redis.Error
功能描述：获取key的生命周期（-1表示不删除）


#### 59、Type
入参: 1、key string 集合类型的key
出参：1、result KeyType key的类型 
/*
const (
    RT_NONE KeyType = iota
    RT_STRING
    RT_SET
    RT_LIST
    RT_ZSET
)
*/
      2、redis.Error
功能描述：获取key的类型


#### 60、Zadd
入参: 1、key string 有序集合类型的key
      2、arg1 float64 位置中的位置
      3、arr2 []byte 插入的值
出参：1、result bool 是否插入成功
      2、redis.Error
功能描述：插入或者插件有序集合


#### 61、Zcard
入参: 1、key string 有序集合类型的key
出参：1、result int64 集合的中成员的个数
      2、redis.Error
功能描述：获取有序集合成员的个数


#### 62、Zrange
入参: 1、key string 有序集合类型的key
      2、arg1 float64 开始位置
      3、arg2 float64 结束位置
出参：1、result [][]byte 取出来的集合
      2、redis.Error
功能描述：获取有序集合中对应位置范围的值，-1表示到尾部


#### 63、Zrangebyscore
入参: 1、key string 有序集合类型的key
      2、arg1 float64 开始位置
      3、arg2 float64 结束位置
出参：1、result [][]byte 取出来的集合
      2、redis.Error
功能描述：获取有序集合中对应位置范围的值，-1表示到尾部。 没看出和Zrange差别


#### 64、Zrem
入参: 1、key string 有序集合类型的key
      2、arg1 []byte 要删除的value
出参：1、result bool 是否操作成功
      2、redis.Error
功能描述：获取有序集合中的值


#### 65、Zrevrange
入参: 1、key string 有序集合类型的key
      2、arg1 float64 开始位置
      3、arg2 float64 结束位置
出参：1、result [][]byte 取出来的集合
      2、redis.Error
功能描述：获取有序集合中对应位置范围的值，-1表示到尾部。 没看出和Zrange差别


#### 66、Zscore
入参: 1、key string 有序集合类型的key
      2、arg1 []byte 成员的值
出参：1、result float64 成员的位置
      2、redis.Error
功能描述：获取有序集合成员的位置。

