## Day 1
* 项目创建
* 相关依赖安装
* 实现项目脚手架搭建
* 实现CreateArticle方法
* api测试成功

### steps:
1、新建目录，创建api文件article.api and main.api

2、在article.api中定义文章服务的请求和响应结构体

3、在main.api中定义文章服务的API

4、在main.api层级利用goctl生成代码
```abpublidot
goctl api go -api main.api -dir ../ --style=goZero
```

5、新建rpc/pb/article.proto,利用sql2pb生成文章服务的rpc定义
```abpublidot
sql2pb -go_package ./pb -host localhost -package pb -password YOURPASSWORD -port 3306 -schema YOURDB -service_name article -user root > article.proto
```

6、利用protocol文件生成rpc层其他代码
```abpublidot
goctl rpc protoc article.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero
```
7、修改api层和rpc层的配置文件，将数据库连接信息填写到配置文件中

8、完善逻辑代码

9、测试


## Day 2
### 实现一个简单的 bookstore 服务，实现添加书籍和查询书籍的功能
steps:
1. 新建目录，创建api文件bookstore.api
2. 在bookstore.api中定义书籍服务的请求和响应结构体
3. 生成apiGetWay代码
```abpublidot
goctl api go -api bookstore.api -dir ../ --style=goZero
```
4. 启动apiGetWay服务并测试
5. 创建add.proto文件，定义添加书籍的rpc服务
6. 生成rpc层代码
```abpublidot
goctl rpc protoc add.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero
```
7. 修改完善serviceContext.go 文件，添加Checker、adder依赖,通过ServiceContext在不同业务逻辑之间传递依赖
8. api中实现调用rpc方法
9. 连接数据库
10. 完善rpc逻辑代码
11. 测试rpc服务

## Day 3
高并发测试
### 测试 bookstore 服务的高并发性能
合理设置后，测试 bookstore 服务的高并发性能，观察服务的响应时间和吞吐量。
设置并发线程数为5000，时间为120s, 测试 bookstore 服务的高并发性能

结果：
* qps:25000左右