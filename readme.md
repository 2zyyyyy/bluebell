
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
- Models：
- pkg
  - snowflake：使用`github.com/sony/sonyflake`雪花算法生成用户id
- config.yaml：项目的配置信息
- web_app.log：项目的日志文件
- main.go：项目主入口