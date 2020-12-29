# logger
convenient log package

# 1. 使用说明
```go
    import  "github.com/gkyh/logger"

    logger.Init(3, `{
        "filename": "logs/out.log",
        "level": "TRAC",
        "maxlines": 1000000,
        "maxsize": 1,
        "daily": true,
        "maxdays": 1,
        "append": true,
        "permit": "0660"
    }`)

    logger.Debug("this is Debug")
    logger.Info("this is Info")
    logger.Warn("this is Warn")
    logger.Error("this is Error")
    logger.Print("this is Critical")
    logger.Println("this is Alert")
```
输出结果：



# 2. 内部实现通用接口 Printf、Println


type Log struct{}

func (g *Log) Printf(format string, args ...interface{}) {
    log.Warn(formatLog(args...), args...)
}
func (g *Log) Println(args ...interface{}) {
    log.Debug(formatLog(args...), args...)
}