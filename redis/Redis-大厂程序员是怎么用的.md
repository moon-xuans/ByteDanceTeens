### Redis-大厂程序员是怎么用的

#### 01.Redis是什么

**为什么需要Redis**

- 数据从单表，演进出了分库分表

- MySQL从单机演进出了集群

  - 数据量增长
  - 读写数据压力的不断增加
  

  

  ![image-20230214080811463](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214080811463.png)

- 数据分冷热

  - 热数据：经常被访问到的数据
  
- 将热数据存储到内存中
  

![image-20230214080954857](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214080954857.png)

**Redis基本工作原理**

- 数据从内存中读写
- 数据保存到硬盘上防止重启数据丢失
  - 数据增量保存到AOF文件
  - 全量数据RDB文件
- 单线程处理所有操作命令

![image-20230214081139338](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214081139338.png)

#### 02.Redis应用案例

##### 2.1.连续签到

![image-20230214081220774](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214081220774.png)

**String数据结构**

数据结构-sds

- 可以存储字符串、数字、二进制数据
- 通常和expire配合使用
- 场景：存储计数、Session

![image-20230214081318776](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214081318776.png)

##### 2.2.消息通知

用list作为消息队列

- 使用场景：消息通知。例如当文章更新时，将更新后的文章推送到ES，用户就能搜索到最新的文章数据

![image-20230214081431847](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214081431847.png)

**List数据结构QuickList**

QuickList由一个双向链表和listpack实现

![image-20230214081528342](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214081528342.png)

**Listpack数据结构**

![image-20230214085043189](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214085043189.png)

##### 2.3.计数

![image-20230214085116075](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214085116075.png)

**Hash数据结构dict**

- rehash：rehash操作时将ht[0]中的数据，全部迁移到ht[1]中。数据量小的场景下，直接将数据从ht[0]拷贝到ht[1]速度是较快的。数据量大的场景，例如存有上百万的kv时，迁移过程将会明显阻塞用户请求。
- 渐进式rehash：为避免出现这种情况，使用了rehash方案。基本原理就是，每次用户访问时都会迁移少量数据。将整个迁移过程，平摊到所有的访问用户请求过程中。

![image-20230214085428576](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214085428576.png)

##### 2.4.排行榜

![image-20230214085507906](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214085507906.png)

**zset数据结构 zskiplist**

![image-20230214085601141](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214085601141.png)

##### 2.5.限流

- 要求1秒内放行的请求为N，超过N则禁止访问

Key:comment_freq_limit_1671356046,对整个key调用Incr，超过限制N则禁止访问，1671356046是当前时间戳

![image-20230214085800340](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214085800340.png)

##### 2.6.分布式锁

并发场景，要求一次只能有一个协程执行。执行完成后，其他等待中的协程才能执行。

可以使用redis的setnx实现，利用了两个特性

- Redis是单线程执行命令
- setnx只有未设置过才能执行成功

![image-20230214085919610](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214085919610.png)

#### 03.Redis使用注意事项

##### 3.1.大Key、热Key

**大Key的定义**

![image-20230214090016252](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214090016252.png)

**大Key的危害**

- 读取成本高
- 容易导致慢查询(过期、删除)
- 主从复制异常，服务阻塞，无法正常响应请求

**业务侧使用大Key的表现**

- 请求Redis超时报错

![image-20230214090137490](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214090137490.png)

**消除大Key的方法**

1.拆分

将大key拆分为小key。例如一个String拆分成多个string

![image-20230214090342751](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214090342751.png)

2.压缩

将value压缩后写入redis，读取时解压后再使用。压缩算法可以是gzip、snappy、lz4等。通常情况下，一个压缩算法压缩率高、则解压耗时就长。需要对实际数据进行测试后，选择一个合适的算法。如果存储的是JSON字符串，可以考虑使用MessagePack进行序列化。

3.集合类结构hash、list、set

(1) 拆分：可以用hash取余、位掩码的方式决定放在哪个key中

(2)区分冷热:如榜单列表场景使用zset，只缓存前10页数据，后续数据走db

**热Key的定义**

用户访问一个Key的QPS特别高，导致Server实例出现CPU负载突增或者不均的情况。热Key没有明确的标准，QPS超过500就有可能被识别为热Key

![image-20230214090835470](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214090835470.png)

**解决热Key的方法**

1.设置Localcache

在访问Redis前，在业务服务侧设置Localcache，降低访问Redis的QPS。Localcache中缓存过期或未命中，则从Redis中将数据更新到LocalCache。Java的Guava、Golang的Bigcache就是这类Localcache。

![image-20230214091021747](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214091021747.png)

2.拆分

将key:value这一个热key复制写入多份，例如key1:value,key2:value,访问的时候访问多个key，但value是同一个，以此将qps分散到不同实例上，降低负载。代价是，更新时需要重新多个key，存在数据短暂不一致的风险

![image-20230214091225442](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214091225442.png)

3.使用Redis代理的热key承载能力

字节的Redis访问代理就具备热Key承载能力。本质上是结合了"热key发现"、“LocalCache”两个功能

![image-20230214091355896](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230214091355896.png)

##### 3.2.慢查询场景

容易导致redis慢查询的操作

(1)批量操作一次性传入过多的key/value，如mset/hmset/sadd/zadd等O(n)操作。建议单批次不要超过100，超过100之后性能下降明显。

(2)zset大部分命令都是O(log(n))，当大小超过5k以上时，简单的zadd/zrem也可能导致慢查询

(3)操作的单个value过大，超过10KB。也即，避免使用大Key

(4)对大key的delete/expire操作也可能导致慢查询，Redis4.0之前不支持异步删除unlink，大key删除会阻塞Redis

##### 3.3.缓存穿透、缓存雪崩

缓存穿透：热点数据查询绕过缓存，直接查询数据库

缓存雪崩：大量缓存同时过期

**缓存穿透的危害**：

(1)查询一个一定不存在的数据

通常不会缓存不存在的数据，这类查询请求都会直接打到db，如果有系统bug或人为攻击，那么容易导致db响应慢甚至宕机

(2)缓存过期时

在高并发场景下，一个热key如果过期，会有大量请求同时击穿至db，容易影响db性能和稳定。同一时间有大量key集中过期时，也会导致大量请求落到db上，导致查询变慢，甚至出现db无法响应新的查询

**如何减少缓存穿透**

(1)缓存空值

如一个不存在的userID。这个id在缓存和数据库中都不存在。则可以缓存一个空值，下次再查缓存直接返回空值。

(2)布隆过滤器

通过bloom filter算法来存储合法key，得益于该算法超高的压缩率，只需占用极小的空间就能存储大量key值

**如何避免缓存雪崩**

(1)缓存空值

将缓存失效时间分散开，比如在原有的失效时间基础上增加一个随机值，例如不同key过期时间，可以设置为10分1秒过期，10分23秒过期，10分8秒过期。单位秒部分就是随机时间，这样过期时间就分散了。对于热点数据，过期时间尽量设置得长一些，冷门的数据可以相对设置过期时间短一些。

(2)使用缓存集群，避免单机宕机造成的缓存雪崩。

