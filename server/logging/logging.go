package logging

import (
	"fmt"
	"os"

	opLogging "github.com/op/go-logging"
)

var log = opLogging.MustGetLogger("logs")

const INTERNAL = "internal"

func MustGetLogger(name string) *Logger {
	host, err := os.Hostname()
	if err != nil {
		log.Error("", err.Error())
		host = "unknown"
	}
	log := &Logger{opLogging.MustGetLogger(name), host}
	fmt.Println(log)
	log.ExtraCalldepth = 1
	return log
}

type Logger struct {
	*opLogging.Logger
	Hostname string
}

func (log *Logger) Debug(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + log.Hostname + "] [" + userid + "]"}, args...)
	log.Logger.Debug("log location "+userid, args...)
}

func (log *Logger) Debugf(userid string, string_format string, args ...interface{}) {
	log.Logger.Debugf("["+log.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (log *Logger) Info(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + log.Hostname + "] [" + userid + "]"}, args...)
	log.Logger.Info("", args)
}

func (log *Logger) Infof(userid string, string_format string, args ...interface{}) {
	log.Logger.Infof("["+log.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (log *Logger) Error(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + log.Hostname + "] [" + userid + "]"}, args...)
	log.Logger.Error("", args)
}

func (log *Logger) Errorf(userid string, string_format string, args ...interface{}) {
	log.Logger.Errorf("["+log.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (log *Logger) Critical(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + log.Hostname + "] [" + userid + "]"}, args...)
	log.Logger.Critical("", args)
}

// func (log *Logger) Criticalf(userid string, string_format string, args ...interface{}) {
// 	log.Logger.Criticalf("["+log.Hostname+"] ["+userid+"] "+string_format, args...)
// }

func (log *Logger) Fatal(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + log.Hostname + "] [" + userid + "]"}, args...)
	log.Logger.Fatal(args)
}

func (log *Logger) Fatalf(userid string, string_format string, args ...interface{}) {
	log.Logger.Fatalf("["+log.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (log *Logger) Panic(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + log.Hostname + "] [" + userid + "]"}, args...)
	log.Logger.Panic(args)
}

func (log *Logger) Panicf(userid string, string_format string, args ...interface{}) {
	log.Logger.Panicf("["+log.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (log *Logger) Warning(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + log.Hostname + "] [" + userid + "]"}, args...)
	log.Logger.Warning("", args)
}

func (log *Logger) Warningf(userid string, string_format string, args ...interface{}) {
	log.Logger.Warningf("["+log.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (log *Logger) Notice(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + log.Hostname + "] [" + userid + "]"}, args...)
	log.Logger.Notice("", args)
}

func (log *Logger) Noticef(userid string, string_format string, args ...interface{}) {
	log.Logger.Noticef("["+log.Hostname+"] ["+userid+"] "+string_format, args...)
}
