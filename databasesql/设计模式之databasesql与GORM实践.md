### 设计模式之databasesql与GORM实践

#### 01.理解database/sql

##### 1.1.基本用法

import driver实现，使用driver + DSN初始化DB连接

执行一条SQL，通过rows取回返回的数据，处理完毕，需要释放连接

数据、错误处理

##### 1.2.设计原理

应用程序通过操作接口操作database/sql,然后通过连接接口以及操作接口，处理数据库

DB连接的几种类型：直接连接、预编译、事务

#### 02.GORM的基本使用

 设计原则：API精简、测试优先、最小惊讶、灵活扩展、无依赖

功能完善：

- 关联：一对一、一对多、单表自关联、多态；Preload、Joins预加载、级联删除；关联模式：自定义关联表
- 事务：事务代码块、嵌套事务、Save Point
- 多数据库、读写分离、命名参数、Map、子查询、分组条件、代码共享、SQL表达式(查询、创建、更新)、自动选字段、查询优化器
- 字段权限、软删除、批量数据处理、Prepared Stmt、自定义类型、命名策略、虚拟字段、自动track时间、SQL Builder、Logger
- 代码生成、复合主键、Constraint、Prometheus、Auto Migration
- 多模式灵活自由扩展
- Developer Friendly

#### 03.GORM设计原理

在应用程序以及database/sql中加了一层GORM

##### 3.1.SQL是怎么生成的

GORM API方法添加Clauses至GORM Statement，到Finisher方法执行Statement

##### 3.2.插件是怎么工作的

Finisher Method -> 决定Statement类型 -> 执行Callbacks -> 生成SQL并执行

