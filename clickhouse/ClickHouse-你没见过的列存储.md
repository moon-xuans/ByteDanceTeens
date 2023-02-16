### ClickHouse-你没见过的列存储

#### 01.数据库基本概念

数据库是结构化信息或数据的有序集合，一般以电子形式存储在计算机系统中。通常由数据库管理系统(DBMS)来控制。在现实中，数据、DBMS及关联应用一起被称为数据库系统，通常简称为数据库。

**数据解析整理成有序集合**

![image-20230215105756002](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215105756002.png)

**可以通过查询语言获取想要的信息**

![image-20230215105916066](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215105916066.png)

**数据库类型**

- 关系数据库：关系型数据库是把数据以表的形式进行储存，然后再各个表之间建立关系，通过这些表之间的关系来操作不同表之间的数据。
- 非关系型数据库：NoSQL或非关系型数据库，支持存储和操作非结构化及半结构化数据。相比于关系型数据库，NoSQL没有固定的表结构，且数据之间不存在表与表之间的关系，数据之间可以是独立的。



- 单机数据库：在一台计算机上完成数据的存储和查询的数据库系统。
- 分布式数据库：分布式数据库由位于不同站点的两个或多个文件组成。数据库可以存储在多台计算机上，位于同一个物理位置，或分散在不同的网络上。



- OLTP数据库：OLTP(Online transactional processing)数据库是一种高速分析数据库，专为多个用户执行大量事务而设计。
- OLAP数据库：OLAP(Online analytical processing)数据库旨在同时分析多个数据维度，帮助团队更好地理解其数据中的复杂关系。

**OLAP数据库**

- 大量数据的读写，PB级别的存储
- 多维分析，复杂的聚合函数
- 窗口函数，自定义UDF
- 离线/实时分析

![image-20230215111103721](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215111103721.png)

![image-20230215111343714](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215111343714.png)

**SQL**

一种编程语言，目前几乎所有的关系数据库都是用SQL编程语言来查询、操作和定义数据，进行数据访问控制。

![image-20230215111912191](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215111912191.png)

一个简单的SQL查询包含SELECT关键词。星号("*")也可以用来指定查询应当返回查询表所有字段，可选的关键词和子句。

FROM子句指定了选择的数据表。FROM子句也可以包含JOIN二层子句来为数据表的连接设置规则

WHERE子句后接一个比较谓词以限制返回的行。WHERE子句仅保留返回结果里使得比较谓词的值为True的行。

GROUP BY子句用于将若干含有相同值的行合并。GROUP BY通常与SQL聚合函数连用，或者用于清除数据重复的行。GROUP BY子句要用在WHERE子句之后。

- 定义数据模型
- 读写数据库数据

![image-20230215112547970](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215112547970.png)

![image-20230215112600450](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215112600450.png)

**SQL的优点**

1. 标准化，ISO的ANSI是长期建立使用的SQL数据库标准

2. 高度非过程化，用SQL进行数据操作，用户只需提出“做什么”，而不必指明“怎么做”，因此用户无需了解存取路径，存取路径的选择以及SQL语句的操作过程由系统自动完成。这不但大大减轻了用户负担，而且有利于提高数据独立性。

3. 以同一种语法结构提供两种使用方式，用户可以在终端上直接输入SQL命令对数据库进行操作。作为嵌入式语言，SQL语句能够嵌入高级语言(如C、C#、JAVA)程序中，供程序员设计程序时使用。而在两种不同的使用方式下，SQL的语法结构基本上是一致的。

4. 语言简洁，易学易用：SQL功能极强，但由于设计巧妙，语言十分简洁，完成数据定义、数据操作、数据控制的核心功能只用了9个动词：CREATE、ALTER、DROP、SELECT、INSERT、UPDATE、DELETE、GRANT、REVOKE。且SQL语言语法简单，接近英语口语，因此容易学习，也容易使用。

 **数据库架构**

![image-20230215113219248](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215113219248.png)

**SQL的执行**

![image-20230215113308061](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215113308061.png)

Parser：词法分析，语法分析，生成AST树

Analyzer:变量绑定、类型推导、语义检查、安全、权限检查、完整性检查等，为生成计划做准备

例如：

- 判断a,b是不是类型正确
- a，b是不是来自表t
- group by字段是否合法，是否存在聚合函数

![image-20230215113854156](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215113854156.png)

Optimizer:为查询生成性能最优的执行计划，进行代价评估

![image-20230215114024686](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215114024686.png)

Executor:将执行计划翻译成可执行的物理计划并驱动其执行

**存储引擎**

1. 管理内存数据结构
   1. 索引
   2. 内存数据
   3. 缓存
2. 管理磁盘数据
   1. 磁盘数据的文件格式
   2. 磁盘数据的增删改查
3. 读写算子
   1. 数据写入逻辑
   2. 数据读取逻辑

![image-20230215114449628](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215114449628.png)

如何存储数据？

- 是否可以并发处理
- 是否可以构建索引
- 行存，列存或者行列混合存储

如何读写数据？

- 读多写少
- 读少写多
- 点查场景
- 分析型场景

#### 02.列式存储

**行式存储**

![image-20230215153753192](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215153753192.png)

**列式存储**

![image-20230215153834308](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215153834308.png)

**列式存储的优点**

数据压缩

- 数据压缩可以使读的数据量更少，在IO密集型计算中获得更大的性能优势
- 相同类型压缩效率更高
- 排序之后压缩效率更高
- 可以针对不同类型使用不同的压缩算法

![image-20230215154032590](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215154032590.png)

LZ4

- (5,4)代表向前5个byte，匹配到的内容长度有4，即"bcde"是一个重复
- 重复项越多或者越长，压缩率就会越高

![image-20230215154254251](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215154254251.png)

Run-length encoding

- 压缩重复的数据
- 可以再压缩数据上直接计算

![image-20230215154819624](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215154819624.png)

Delta encoding：

- 将数据存储为连续数据之间的差异，而不是直接存储数据本身
- 特定算子也能直接在压缩数据上计算

![image-20230215155115845](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215155115845.png)

数据选择：

- 可以选择特定的列做计算而不是读所有列
- 对聚合计算友好

![image-20230215155247525](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215155247525.png)

延迟物化：

物化：将列数据转换为可以被计算或者输出的行数据或者内存数据结果的过程，物化后的数据通常可以用来做数据过滤，聚合计算，Join

![image-20230215155518833](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215155518833.png)

延迟物化：尽可能推迟物化操作的发生

![image-20230215155727923](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215155727923.png)

延迟物化：

- 缓存友好
- CPU/内存带宽友好
- 可以利用到执行计划和算子的优化，例如filter
- 保留直接在压缩列做计算的机会



向量化：

- SIMD
- 数据格式
- 执行模型



![image-20230215160009122](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215160009122.png)

SIMD(single instruction multiple data),对于现代多核CPU，其都有能力用一条指令执行多条数据

SIMD程序使用的指令集有SSE和AVE系列，AVX有AVX-256和AVX-512,SSE提供128-bits的寄存器，AVX-256提供256-bits，AVX-512提供512bits的寄存器

![image-20230215160300365](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215160300365.png)

如果这时候CPU也可以并行的计算我们写的代码，那么理论上我们的处理速度就会是之前代码的100倍，SIMD指令就可以完成这样的工作，用SIMD指令完成的代码设计和执行的逻辑就叫做向量化。

![image-20230215180342762](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215180342762.png)

数据格式要求：

- 需要处理多个数据，因此数据需要是连续内存
- 需要明确数据类型

执行模型要求：

- 数据需要按批处理
- 函数的调用需要明确数据类型

列存数据库适合设计出这样的执行模型，从而使用向量化技术：

- 按列读取
- 每种列类型定义数据读写逻辑
- 函数按列类型处理

#### 03.ClickHouse存储设计

**表定义和结构**

![image-20230215180837248](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215180837248.png)

**集群架构**

![image-20230215180959770](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215180959770.png)

**引擎架构**

![image-20230215181021833](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215181021833.png)

**存储架构**

![image-20230215181044543](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215181044543.png)

![image-20230215181113763](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215181113763.png)

part和partition

- part是物理文件夹的名字
- partition是逻辑结构

![image-20230215181205919](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215181205919.png)

part和column

- 每个column都是一个文件
- 所有的column文件都在自己的part文件夹下

column和index

- 一个part有一个主键索引
- 每个column都有列索引

**索引设计**

Hash Index

1. 将输入的key通过一个HashFunction映射到一组bucket上
2. 每个bucket都包含一个指向一条记录的地址
3. 哈希索引在查找的时候只适应于等值比较

![image-20230215181801360](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215181801360.png)

B-Tree

1. 数据写入是有序的，支持增删改查
2. 每个节点有多个孩子节点
3. 每个节点都按照升序排列key值
4. 每个key有两个指向左右孩子节点的引用
   1. 左孩子节点保存的key都小于当前key
   2. 右孩子节点的保存的key都大于当前key

![image-20230215182122086](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215182122086.png)

B+Tree

1. 所有的数据都存储在叶子节点，非叶子节点只保存key值
2. 叶子节点维护到相邻叶子节点的引用
3. 可以通过key值做二分查找，也可以通过叶子节点做顺序访问

![image-20230215182512931](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215182512931.png)

**索引设计**

- 对于大数据量，B(B+)-Tree深度太高
- 索引数据量太大，多个列如何平衡查询和存储 —— LSM-Tree
- OLAP场景写入量非常大，如何优化写入

Log-structured merge-tree(LSM tree)是一种为大吞吐写入场景而设计的数据结构

- 着重优化顺序写入
- 主要数据结构
  - SSTables
  - Memtable

SSTables

1. Key按顺序存储到文件中，称为segment
2. 包含多个segment
3. 每个segment写入磁盘后都是不可更改的，新加的数据只能生成新的segment

![image-20230215183906646](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215183906646.png)

Memtable

- 在内存中的数据保存在memtable中，大多数实现都是一颗Binary search tree
- 当memtable存储的数据到达一定的阈值的时候，就会按顺序写入到磁盘

数据查询

- 需要从最新的segment开始遍历每个key
- 也可以为每个segment建一个索引，例如

![image-20230215184240052](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215184240052.png)

Compaction(合并)

- Compaction指将多个segments合并成一个segments的过程
- 一般是有一个后台线程完成
- 不同的segments写入新的

segment的时候也是需要排序，形成新的segment之后，旧的segment文件就会被删除

![image-20230215184611251](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215184611251.png)

**索引实现**

主键索引

![image-20230215184641462](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215184641462.png)

数据按照主键顺序依次做排序

- 首先按照UserID做排序
- 再按照URL排序
- 最后是EventTime

![image-20230215184758038](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215184758038.png)

数据被划分为granules

1. granules是最小的数据读取单元
2. 不同的granulas可以并行读取

![image-20230215190042099](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215190042099.png)

每个granule都对应primary.idx里面的一行

![image-20230215190530820](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215190530820.png)

默认每8192行记录主键的一行值，primary.idx需要被全部加载到内存里面

![image-20230215190736771](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215190736771.png)

里面保存的每一行数据称为一个index mark

![image-20230215190946304](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215190946304.png)

每个列都有这样一个mark文件

1. mark文件保存的是每个granules的物理地址
2. 每一列都有一个自己的mark文件

![image-20230215191118837](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215191118837.png)

mark文件里面的每一行保存两个地址：

1. block_offset:用于定位一个granule的压缩数据在物理文件中的位置，压缩数据会以一个block为单位解压到内存中
2. granule_offset,用于定位一个granule在解压之后的block中的位置

缺陷：数据按照key的顺序做排序，因此只有第一个key的过滤效果好，后面的key过滤效果依赖第一个key的基数大小

![image-20230215191433322](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215191433322.png)

**查询优化**

secondary index:在URL列上构建二级索引

![image-20230215191515947](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215191515947.png)

构建多个主键索引

- 再建一个表，使用需要优化的字段做主键第一位
- 建一个物化视图
- 使用Projection

再建一个表

![image-20230215191636689](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215191636689.png)

再建一个表

1. 数据需要同步两份
2. 查询需要用户判断查哪张表

![image-20230215191814972](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215191814972.png)

建一个物化视图

物化视图：可以通过select查询将一个表的数据写入一张隐式表

![image-20230215191955908](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215191955908.png)

![image-20230215192109180](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215192109180.png)

使用Projection

Projection：类似于物化试图，但是不是将数据写入新的表，而是存储在原始表，以一个列文件的形式存在。

![image-20230215192208153](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215192208153.png)

1. 数据自动同步到隐式表
2. 查询自动路由到最优的表

![image-20230215192253363](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215192253363.png)

**数据合并**

![image-20230215192323559](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215192323559.png)

数据的可见性

1. 数据合并过程中，未被合并的数据对查询可见
2. 数据合并完成后，新part可见，被合并的part被标记删除

![image-20230215192439725](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215192439725.png)

1.通过主键找到需要读的mark

![image-20230215192504932](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215192504932.png)

2.切分marks，然后并发的调度reader

![image-20230215192535377](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215192535377.png)

3.Reader通过mark,block_offset得到需要读的数据文件的偏移量

4.Reader通过mark,granule_offset得到解压之后数据的偏移量

![image-20230215192642719](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215192642719.png)

5.构建列式filter做数据过滤

![image-20230215192717171](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230215192717171.png)

#### 04.ClickHouse典型应用场景

大宽表存储和查询

1.大宽表查询

- 可以建非常多的列
- 可以增加，删除，清空每一列的数据
- 查询的时候引擎可以快速选择需要的列
- 可以将列涉及的过滤条件下推到存储层从而加速查询

![image-20230216084842811](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216084842811.png)

2.动态表结构

- map中的每个key都是一列
- map中的每一列都可以单独的查询
- 使用方式同普通列，可以做任何计算

![image-20230216084909879](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216084909879.png)

离线数据分析

1.数据导入

- 数据可以通过spark生成clickhouse格式的文件
- 导入到hdfs上由hive2ch导入工具完成数据导入
- 数据直接导入到各个物理节点

![image-20230216085044464](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085044464.png)

2.数据按列导入

- 保证查询可以及时访问已有数据
- 可以按需加载需要的列

![image-20230216085059431](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085059431.png)

![image-20230216085125634](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085125634.png)

使用memory table减少parts的数量

1. 数据先缓存在内存中
2. 到达一定阈值再写到磁盘

![image-20230216085150063](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085150063.png)

复杂类型查询

1.bitmap索引

![image-20230216085220936](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085220936.png)

2.Bitmap64类型

![image-20230216085251039](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085251039.png)

3.lowcardinality

1. 对于低基数列使用字典编码
2. 减少数据存储和读写的IO使用
3. 可以做运行时的压缩数据过滤

![image-20230216085336743](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085336743.png)

![image-20230216085354010](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085354010.png)

![image-20230216085406129](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085406129.png)

![image-20230216085417804](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230216085417804.png)
