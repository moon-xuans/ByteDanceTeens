### 深入理解RDBMS

#### 01.经典案例

##### 1.1.从一场红包雨说起

![image-20230212101014140](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101014140.png)

##### 1.2.RDBMS事务ACID

![image-20230212101223568](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101223568.png)

##### 1.3.红包雨与ACID

![image-20230212101246539](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101246539.png)

![image-20230212101302761](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101302761.png)

![image-20230212101316754](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101316754.png)

![image-20230212101329855](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101329855.png)

##### 1.4.红包雨与高并发

![image-20230212101455708](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101455708.png)

##### 1.5.红包雨与高可靠

![image-20230212101553616](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101553616.png)

#### 02.发展历史

##### 2.1.前DBMS时代 - 人工管理

![image-20230212101708305](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101708305.png)

##### 2.2.前DBMS时代 - 文件系统

![image-20230212101843208](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212101843208.png)

##### 2.3.DBMS时代

1960s，传统的文件系统已经不能满足人们的需求，数据库管理系统(DBMS)应运而生。

DBMS：按照某种数据模型来组织、存储和管理数据的仓库。

所以通常按照数据模型的特点将传统数据库系统分成`网状数据库、层次数据库和关系数据库`三类。

![image-20230212102202760](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212102202760.png)

###### 2.3.1.DBMS数据模型 - 网状模型

网状数据库所基于的网状数据模型建立的数据之间的联系，能反映现实世界中信息的关联，是许多空间对象的自然表达形式。1964年，世界上第一个数据库系统——集成数据存储诞生于通用电气公司。IDS是世界上第一个网状数据库，奠定了数据库发展的基础，在当时得到了广泛的应用。在1970s网状数据库系统十分流行，在数据库系统产品中占据主导地位。

![image-20230212102454300](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212102454300.png)

###### 2.3.2.DBMS数据模型 - 层次模型

1968年，世界上第一个层次数据库——信息管理系统诞生于IBM公司，这也是世界上第一个大型商用的数据库系统。层次数据模型，即使用树形结构来描述实体及其之间关系的数据模型。

![image-20230212113206882](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212113206882.png)

###### 2.3.3.DBMS数据模型 - 关系模型

1970年，IBM的研究员发表了一篇名为“模型的概念”，奠定了关系模型的理论基础。1979年Oracle首次将关系型数据库商业化，后续DB2，SAP Sysbase ASE,and Informix等知名数据库产品也纷纷面世。

![image-20230212113536797](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212113536797.png)

##### 2.4.DBMS数据模型

![image-20230212113610406](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212113610406.png)

##### 2.5.SQL语言

1974年IBM的Ray Boyce和Don Chamberlin将Codd关系数据库的12条准则的数学定义以简单的关键字语法表现出来，里程碑式地提出了SQL语言。

- 语法风格接近自然语言
- 高度非过程化
- 面向集合的操作方式
- 语言简洁，易学易用

![image-20230212113820806](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212113820806.png)

##### 2.6.历史回顾

![image-20230212113858943](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212113858943.png)

#### 03.关键技术

##### 3.1.一条SQL的一生

![image-20230212113931519](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212113931519.png)

##### 3.2.SQL引擎 

###### 3.2.1.Parser

解析器一般分为词法分析、语法分析、语义分析等步骤

![image-20230212114027790](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114027790.png)

###### 3.2.2.Optimizer

为什么需要一个优化器？

![image-20230212114118472](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114118472.png)

基于规则的优化

![image-20230212114152208](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114152208.png)

基于代价的优化

一种查询有多种执行方案，CBO会选择其中代价最低的方案取真正的执行。

什么是代价？

![image-20230212114242736](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114242736.png)

###### 3.3.3.Executor

火山模型：

![image-20230212114411154](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114411154.png)

向量化：

![image-20230212114432689](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114432689.png)

编译执行：

![image-20230212114504351](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114504351.png)

##### 3.3.存储引擎 

###### 3.3.1.InnoDB

![image-20230212114531497](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114531497.png)

###### 3.3.2.Buffer Pool

![image-20230212114603802](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114603802.png)

![image-20230212114623033](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114623033.png)

###### 3.3.3.Page

![image-20230212114653246](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114653246.png)

###### 3.3.4.B+Tree

![image-20230212114744501](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114744501.png)

##### 3.4.事务引擎

![image-20230212114850809](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212114850809.png)

###### 3.4.1.Atomicity与Undo Log

![image-20230212115234044](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212115234044.png)

###### 3.4.2.Isolation与锁

![image-20230212115301943](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212115301943.png)

###### 3.4.3.Isolation与MVCC

![image-20230212115330940](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212115330940.png)

###### 3.4.4.Durability与Redo Log

![image-20230212115404426](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212115404426.png)

#### 04.企业实践

##### 4.1.春节红包雨挑战

![image-20230212115509697](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212115509697.png)

##### 4.2.大流量 - Sharding

问题背景：

- 单节点写容易成为瓶颈
- 单机数据容量上限

解决方案：

- 业务数据进行水平拆分
- 代理层进行分片路由

实施效果：

- 数据库写入性能线性扩展
- 数据库容量线性扩展

![image-20230212115741223](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212115741223.png)

##### 4.3.流量突增 - 扩容

问题背景

- 活动流量上涨
- 集群性能不满足要求

解决方案

- 扩容DB物理节点数量
- 利用影子表进行压测

实施效果

- 数据库集群提供更高的吞吐
- 保证集群可以承担预期流量

![image-20230212115915843](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212115915843.png)

##### 4.4.流量突增 - 代理连接池

问题背景

- 突增流量导致大量建联
- 大量建联导致负载变大，延时上升

解决方案

- 业务侧预热连接池
- 代理侧预热连接池
- 代理侧支持连接池

实施效果

- 避免DB被突增流量打死
- 避免代理和DB被大量建联打死

##### 4.5.稳定性&可靠性

###### 4.5.1.3AZ高可用

![image-20230212120224219](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212120224219.png)

###### 4.5.2.HA管理

问题背景

- db所在机器异常宕机
- db节点异常宕机

解决方案

- ha服务监管、切换宕机节点
- 代理支持配置热加载
- 代理自动屏蔽宕机读节点

实施效果

- 读节点宕机秒级恢复
- 写节点宕机30s内恢复服务

![image-20230212120401214](https://raw.githubusercontent.com/moon-xuans/mediaImage/main/2023/image-20230212120401214.png)
