# TLOG

> 对go.uber.org/zap日志库的包装


## 基本使用

在默认情况下，tlog仅会在console输出debug等级以上的日志

```go
tlog.Info("this is an info level message")

tlog.WithField("name", "tlog").Info("this is an info level message with a field named name")

tlog.WithError(io.EOF).Error("this is an error level message with error io.EOF")
```

## 通过配置使用

用户可以通过名为`config.yaml`的配置文件来达到一定程度的日志自定义输出，以下是一个配置文件示例

```yaml
log:
  - type: console
    level: debug
  - type: file
    level: debug
    prefix: tlog
  - type: file
    level: error
    prefix: tlog.error
    max-age: 24h
    rotation-time: 1h
    rotation-size: 1
```

配置支持的参数：

| 配置名称      | 类型   | 默认值  | 解释                                                                   |
| ------------- | ------ | ------- |----------------------------------------------------------------------|
| type          | string | /       | 支持console或file<br />console类型的日志即使配置了多个也只有第一个生效                      |
| level         | string | info    | 日志等级：debug、info、warn、error、fatal                                     |
| prefix        | string | default | 文件名前缀，**必须全局唯一**<br />若prefix为tmp，则文件名为tmp.log<br />（仅在type为file时生效） |
| max-age       | string | 168h    | 文件最大保存时间，使用time.ParseDuration函数进行计算<br />（仅在type为file时生效）            |
| rotation-time | string | 24h     | 文件切割时间间隔，使用time.ParseDuration函数进行计算<br />（仅在type为file时生效）            |
| rotation-size | Int    | /       | 文件最大大小，单位MB，不填写则不限制文件大小<br />（仅在type为file时生效）                        |

