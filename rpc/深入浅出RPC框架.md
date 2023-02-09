### 深入浅出RPC框架

#### 01.基本概念

##### 1.1.本地函数调用

##### 1.2.远程函数调用

RPC需要解决的问题：1.函数映射；2.数据转换成字节流；3.网络传输

##### 1.3.RPC概念模型

![image-20230209092318475](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209092318475.png)

1984年Nelson发表了论文，其中提出了RPC的过程由5个模型组成：User、User-Stub、RPC-Runtime、Server-Stub、Server

##### 1.4.一次RPC的完整过程

![image-20230209092342324](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209092342324.png)

**IDL(Interface description language)文件**

IDL通过一种中立的方式来描述接口，使得在不同平台上运行的对象和用不同语言编写的程序可以相互通信

**生成代码**

通过编译器工具把IDL文件转换成语言对应的静态库

**编解码**

从内存中表示字节序列的转换称为编码，反之为解码，也常叫做序列化和反序列化

**通信协议**

规范了数据在网络中的传输内容和格式。除必须的请求/响应数据外，通常还会包含额外的元数据

**网络传输**

通常基于成熟的网络库走TCP/UDP传输

##### 1.5.RPC的好处

1.单一职责，有利于分工协作和运维开发

2.可扩展性强，资源使用率更优

3.故障隔离，服务的整体可靠性更高

##### 1.6.RPC带来的问题

1.服务宕机，对方该如何处理？

2.在调用过程中发生网络异常，如何保证消息的可达性？

3.请求量突增导致服务无法及时处理，有哪些应对措施？

#### 02.分层设计

##### 2.1.以Apache Thrift为例

![image-20230209092751568](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209092751568.png)

##### 2.2.编解码层

![image-20230209093030938](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209093030938.png)

##### 2.3.编解码层-生成代码

![image-20230209093100834](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209093100834.png)

##### 2.4.编解码层-数据格式

- 语言特定的格式：许多编程语言都内建了将内存对象为字节序列的支持，例如Java有Java.io.Serializable
- 文本格式：JSON、XML、CSV等文本格式，具有人类可读性
- 二进制编码：具备跨语言和高性能等优点，常见有Thrift的BinaryProtocol，Protobuf等

##### 2.5.编解码层-二进制编码

TLV编码

- Tag：标签，可以理解为类型
- Length：长度
- Value：值，Value也可以是个TLV结构

![image-20230209093737864](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209093737864.png)

##### 2.6.编解码层-选型

- 兼容性：支持自动增加新的字段，而不影响老的服务，这将提高系统的灵活度
- 通用性：支持跨平台、跨语言
- 性能：从空间和时间两个维度来考虑，也就是编码后数据大小和编码耗费时长

##### 2.7.协议层

![image-20230209093944625](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209093944625.png)

##### 2.8.协议层-概念

- 特殊结束符：一个特殊字符作为每个协议单元结束的标识

![image-20230209094053638](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209094053638.png)

- 变长协议：以定长加不定长的部分组成，其中定长的部分需要描述不定长的内容长度

![image-20230209094221981](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209094221981.png)

##### 2.9.协议层-协议构造

![image-20230209094252649](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209094252649.png)

LENGTH：数据包大小，不包含自身

HEADER MAGIC：表示版本信息，协议解析时候快速校验

SEQUENCE NUMBER：表示数据包的seqID，可用于多路复用，单连接内递增

HEADER SIZE：头部长度，从第14个字节开始计算一直到PAYLOAD前

PROTOCOL ID：编解码方式，有Binary和Compact两种

TRANSFORM ID：压缩方式，如zlib和snappy

INFO ID：传递一些定制的meta信息

PAYLOAD：消息体

##### 2.10.协议层-协议解析

![image-20230209094539690](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209094539690.png)

##### 2.11.网络通信层

![image-20230209100744661](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209100744661.png)

##### 2.12.网络通信层-Sockets API

![image-20230209100834493](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209100834493.png)

##### 2.13.网络通信层-网络库

- 提供易用API
  - 封装底层Socket API
  - 连接管理和事件分发
- 功能
  - 协议支持：tcp、udp和uds等
  - 优雅退出、异常处理等
- 性能
  - 应用层buffer减少copy
  - 高性能定时器、对象池等

#### 03.关键指标

##### 3.1.稳定性-保障策略

- 熔断：保护调用方，防止被调用的服务出现问题而影响到整个链路
- 限流：保护被调用方，防止大流量把服务压垮
- 超时控制：避免浪费资源在不可用节点上

![image-20230209101209582](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209101209582.png)]

##### 3.2.稳定性-请求成功率

负载均衡、重试

![image-20230209101416317](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209101416317.png)

![image-20230209101429441](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209101429441.png)

##### 3.3.稳定性-长尾请求

![image-20230209101557636](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209101557636.png)

##### 3.4.稳定性-注册中间件

![image-20230209101623263](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209101623263.png)

##### 3.5.易用性

- 开箱即用：合理的默认参数选项、丰富的文档
- 周边工具：生成代码工具、脚手架工具

##### 3.6.扩展性

- Middleware
- Option
- 编解码层
- 协议层
- 网络传输层
- 代码生成工具插件扩展

##### 3.7.观测性

- Log、Metric、Tracing
- 内置观测性服务

![image-20230209101914185](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209101914185.png)

##### 3.8.高性能

- 场景
  - 单机多机
  - 单连接多连接
  - 单/多client 单/多server
  - 不同大小的请求包
  - 不同请求类型：例如pingpong、streaming等
- 目标
  - 高吞吐
  - 低延迟
- 手段
  - 连接池
  - 多路复用
  - 高性能编解码协议
  - 高性能网络库

#### 04.企业实践

 ##### 4.1.整体架构-Kitex

- Kitex Core 核心组件
- Kitex Byted 与公司内部基础设施集成
- Kitex Tool 代码生成工具

![image-20230209145639859](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209145639859.png)

##### 4.2.自研网络库 - 背景

- 原生库无法感知连接状态：在使用连接池时，池中存在失效连接，影响连接池的复用
- 原生库存在goroutine暴涨的风险：一个连接一个goroutine的模式，由于连接利用率低下，存在大量goroutine占用调度开销，影响性能。

##### 4.3.自研网络库-Netpoll

- 解决无法感知连接状态问题：引入epoll主动监听机制，感知连接状态
- 解决goroutine暴涨的风险：建立goroutine池，复用goroutine
- 提升性能：引入Nocopy Buffer,向上层提供NoCopy的调用接口，编解码层零拷贝

##### 4.4.扩展性设计

支持多协议，也支持灵活的自定义协议扩展

##### 4.5.性能优化-网络库优化

- 调度优化：
  - epoll_wait在调度上的控制
  - gopool重用goroutine降低同时运行协程数
- LinkBuffer：
  - 读写并行无锁，支持nocopy地流式读写
  - 高效扩缩容
  - Nocopy Buffer池化，减少GC
- Pool：
  - 引入内存池和对象池，减少GC开销

##### 4.6.性能优化-编解码优化

- Codegen
  - 预计算并预分配内存，减少内存操作次数，包括内存分配和拷贝
  - Inline减少函数调用次数和避免不必要的反射操作等
  - 自研了Go语言实现的Thrift IDL解析和代码生成器，支持完善的Thrift IDL语法和语义检查，并支持了插件机制 - Thriftgo
- JIT
  - 使用JIT编译计数改善用户体验的同时带来更强的编解码性能，减轻用户维护生成代码的负担
  - 基于JIT编译技术的高性能动态Thrift编解码器 - Frugal

##### 4.7.合并部署

微服务过微，传输和序列化开销越来越大

将亲和性强的服务实例尽可能调度到同一个物理机，远程RPC调用优化为本地IPC调用

![image-20230209151142659](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209151142659.png)



![image-20230209151158568](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230209151158568.png)

- 中心化的部署调度和流量控制
- 基于共享内存的通信协议
- 定制化的服务发现和连接池实现
- 定制化的服务启动和监听逻辑







