package logger

import (
	"strings"
)

var log *LocalLogger

type Log struct{}

func (g *Log) Printf(format string, args ...interface{}) {

	log.Warn(formatLog(args...), args...)
}
func (g *Log) Println(args ...interface{}) {

	log.Debug(formatLog(args...), args...)
}

/*
 `{
        "filename": "logs/zjt.log",
        "level": "TRAC",
        "maxlines": 1000000,
         "maxsize": 1,
        "daily": true,
        "maxdays": 1,
        "append": true,
        "permit": "0660"
    }
*/
func Init(depth int, config string) {

	log = NewLogger(depth)
	log.SetLogger(`file`, config)
}

func formatLog(v ...interface{}) string {
	var msg string = ""

	if len(v) == 0 {
		return msg
	}

	msg += strings.Repeat(" %v", len(v))

	return msg
}

func Info(args ...interface{}) {

	log.Info(formatLog(args...), args...)
}

func Warn(args ...interface{}) {

	log.Warn(formatLog(args...), args...)
}

func Error(args ...interface{}) {

	log.Error(formatLog(args...), args...)

}

func Fatal(args ...interface{}) {

	log.Emer(formatLog(args...), args...)

}

func Debug(args ...interface{}) {

	log.Debug(formatLog(args...), args...)
	//go PushMsg(args...)
}

func Print(args ...interface{}) {

	log.Info(formatLog(args...), args...)

}

func Println(args ...interface{}) {

	log.Info(formatLog(args...), args...)
}
