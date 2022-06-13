
### 项目结构说明

- Controller：服务的入口，负责处理路由、参数校验、请求转发
- Service：服务层（或者叫逻辑层），负责处理业务逻辑
- Dao：负责数据与存储相关功能
  - mysql：使用 `github.com/go-sql-driver/mysql` 连接mysql数据库
  - redis：使用 `github.com/go-redis/redis` 连接redis
- Logger：日志服务
  - 使用 `go.uber.org/zap` 日志库
  - 使用 `github.com/natefinch/lumberjack` 对日志文件做切割
  - 自定义中间件替换gin框架默认的两个中间件服务
- Settings：整个项目的配置信息
  - 使用 `github.com/spf13/viper`读取配置文件，并反序列化至结构体
- Routers：路由层
- Models：存放数据相关的结构体，如接收前端的参数和返回结果的数据格式
- pkg
  - snowflake：使用`github.com/sony/sonyflake`雪花算法生成用户id
  - jwt：使用`github.com/dgrijalva/jwt-go`来实现我们生成JWT和解析JWT的功能
    - jwt简介：JWT全称JSON Web Token是一种跨域认证解决方案，属于一个开放的标准，它规定了一种Token实现方式，目前多用于前后端分离项目和OAuth2.0业务场景下。
    - JWT全称JSON Web Token是一种跨域认证解决方案，属于一个开放的标准，它规定了一种Token实现方式，目前多用于前后端分离项目和OAuth2.0业务场景下。
- config.yaml：项目的配置信息
- web_app.log：项目的日志文件
- main.go：项目主入口
- Makefile：借助Makefile我们在编译过程中不再需要每次手动输入编译的命令和编译的参数，可以极大简化项目编译过程。
  - Makefile简介：我们可以把Makefile简单理解为它定义了一个项目文件的编译规则。借助Makefile我们在编译过程中不再需要每次手动输入编译的命令和编译的参数，可以极大简化项目编译过程。同时使用Makefile也可以在项目中确定具体的编译规则和流程，很多开源项目中都会定义Makefile文件。

*注意事项:*
*1.如果生成了swagger，在swagger请求对应的接口的时候，要确保浏览器地址栏里面的地址和你的main函数中host是一致的，否则容易出现跨域问题。*
*1的错误示例：我浏览器访问swagger的host是localhost，但是main中配置的地址却是127.0.0.1.在使用swagger调用接口的时候接口就会返回跨域的问题*