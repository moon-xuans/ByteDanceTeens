### 框架笔记

#### 1.Gorm

https://gorm.cn/zh_CN/docs/

##### 1.1.基本使用

Gorm的约定(默认)

Gorm使用名为ID的字段作为主键

使用结构体的蛇形负数作为表名

字段名的蛇形作为列名

使用CreateAt、UpdateAt字段作为创建、更新时间

##### 1.2.支持的数据库

Gorm目前支持MySQL、SQLServer、PostgreSQL、SQLite。

Gorm通过驱动来连接数据库，如果需要连接其他类型的数据库，可以复用/自行开发驱动

##### 1.3.创建数据

**如何使用Upsert?**

使用clause.OnConflict处理数据冲突。如果不解决的话，那么数据就无法插入成功

**如何设置默认值？**

通过使用default标签为字段定义默认值

##### 1.4.查询数据

**First的使用踩坑**

使用First时，需要注意查询不到数据会返回ErrRecordNotFound。使用Find查询多条数据，查询不到数据不会返回错误

**使用结构体作为查询条件**

当使用结构体作为条件查询时，GORM只会查询非零值字段。这意味着如果字段值为0、""、false或者其他零值，该字段不会被用于构建查询条件，使用Map来构建查询条件

##### 1.5.更新数据

使用Struct更新时，只会更新非零值，如果需要更新零值可以使用Map更新或使用Select选择字段。

##### 1.6.删除数据

Gorm有物理删除和软删除。

Gorm提供了gorm.DeleteAt用于帮助用户实现软删除。

拥有软删除能力的Model调用Delete时，记录不会被从数据库中真正删除。但Gorm会将DeleteAt置为当前时间，并且不能通过正常的查询方法找到该记录。

使用Unscoped可以查询到被软删除的数据

##### 1.7.事务

Gorm提供了Begin、Commit、Rollback方法用于使用事务

提供了Transaction方法用于自动提交事务，避免拥护漏写Commit、Rollback。

##### 1.8.Hook

Gorm在提供了CURD的Hook能力。

Hook是在创建、查询、更新、删除等操作之前、之后自动调用的函数。

如果任何Hook返回错误，GORM将停止后续的操作并回滚事务。

##### 1.9.性能提高

对于写操作(创建、更新、删除)，为了确保数据的完整性，GORM会将他们封装在事务内运行。但这会降低性能，可以使用SkipDefaultTransaction关闭事务。

使用PrepareStmt缓存预编译语句可以提高后续调用的速度，本机测试提高大约35%左右。

##### 1.10.生态

http://github.com/go-gorm

#### 2.Kitex

https://www.cloudwego.io/zh/docs/kitex

##### 2.1.安装代码生成工具

Kitex目前对Windows支持不完善。

生成工具：

```shell
go install github.com/cloudwego/kitex/tool/kitex@latest
go install github.com/cloudwego/thiftgo@latest
```

##### 2.2.定义IDL

使用IDL定义服务与接口

如果需要进行RPC，就需要知道对方的接口是什么，需要传什么参数，同时也需要知道返回值是什么样的。这时候，就需要通过IDL来约定对方的协议，就像写代码的时候需要调用某个函数，我们需要知道函数签名一样。

Thrift：https://thrift.apache.org/docs/idl

Proto3:https://developers.google.com/protocol-buffers/docs/proto3

##### 2.3.生成代码

使用kitex -module example -service example echo.thrift 命令生成代码

https://www.cloudwego.io/zh/docs/kitex/getting-started/

##### 2.4.生态

https://github.com/kitex-contrib/

#### 3.Hertz

https://www.cloudwego.io/zh/docs/hertz/getting-started/

##### 3.1.基本使用

使用Hertz实现，服务监听8080端口并注册了一个GET方法的路由函数

##### 3.2.路由

Hertz提供了GET、POST、PUT、DELETE、ANY等方法用于注册路由。

提供了路由组的能力，用于支持路由分组的功能。

提供了参数路由和通配路由，路由的优先级为：静态路由 > 命名路由 > 统配路由

##### 3.3.参数绑定

Hertz提供了Bind、Validate、BindAndValidate函数用于进行参数绑定和校验

##### 3.4.代码生成工具

提供了代码生成工具，通过定义IDL文件即可生成对应的基础服务代码

##### 3.5.生态

https://github.com/hertz-contrib

https://github.com/cloudwego/hertz#extensions