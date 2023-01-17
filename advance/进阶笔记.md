### 进阶

#### 1.语言进阶

**Goroutine**

线程：用户态，轻量级线程，栈MB级别

协程：内核台，线程多个协程，栈KB级别

**CSP**

协程是通过通道共享内存，而不是通过共享内存实现通信

**Channel**

可通过make(chan 元素类型,[缓冲大小])，来得到通道

**Lock**

可以通过sync.Mutex来实现并发安全

**WaitGroup**

可以实现阻塞等待，类似java中的countDownLatch

#### 2.依赖管理

通过go.mod文件管理依赖包版本

通过go get/go mod 指令工具管理依赖包

**go mod**

go mod init 初始化，创建go.mod文件

go mod download 下载模块到本地缓存

go mod tidy 增加需要的依赖，删除不需要的依赖

#### 3.测试

可以使用mock来完成单元测试

go test judgment_test.go judegment.go --cover 可以得到测试的覆盖率

在并发情况下使用rand来生成随机数性能较差，可以使用fastrand来生成

