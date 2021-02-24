package logger

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	PanicLevel int = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

type logFile struct {
	level    int
	logTime  int64
	name string
	fd   *os.File
}

var defaultLog logFile

func InitLevel(name string,l int) {
	defaultLog.name = name
	defaultLog.level = l

	//log.SetOutput(defaultLog)
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)

}

// Debug 调试
func Debug(args ...interface{}) {
	if defaultLog.level >= DebugLevel {
		log.SetPrefix("debug ")
		log.Output(2, fmt.Sprintln(args...))
	}
}

func Debugf(format string, args ...interface{}) {
	if defaultLog.level >= DebugLevel {
		log.SetPrefix("debug ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

// Info 普通
func Info(args ...interface{}) {
	if defaultLog.level >= InfoLevel {
		log.SetPrefix("info ")
		log.Output(2, fmt.Sprintln(args...))
	}
}

func Infof(format string, args ...interface{}) {
	if defaultLog.level >= InfoLevel {
		log.SetPrefix("info ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

// Warn 警告
func Warn(args ...interface{}) {
	if defaultLog.level >= WarnLevel {
		log.SetPrefix("warn ")
		log.Output(2, fmt.Sprintln(args...))
	}
}

func Warnf(format string, args ...interface{}) {
	if defaultLog.level >= WarnLevel {
		log.SetPrefix("warn ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

// Error 错误
func Error(args ...interface{}) {
	if defaultLog.level >= ErrorLevel {
		log.SetPrefix("error ")
		log.Output(2, fmt.Sprintln(args...))
	}
}

func Errorf(format string, args ...interface{}) {
	if defaultLog.level >= ErrorLevel {
		log.SetPrefix("error ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func Fatalf(format string, args ...interface{}) {
	if defaultLog.level >= FatalLevel {
		log.SetPrefix("fatal ")
		log.Output(2, fmt.Sprintf(format, args...))
	}
}

func (file *logFile) Write(buf []byte) (n int, err error) {
	if file.name == "" {
		fmt.Printf("consol: %s", buf)
		return len(buf), nil
	}

	if file.logTime+3600 < time.Now().Unix() {
		file.createLogFile()
		file.logTime = time.Now().Unix()
	}

	if file.fd == nil {
		return len(buf), nil
	}

	return file.fd.Write(buf)
}

func (file *logFile) createLogFile() {
	logdir := "./"
	if index := strings.LastIndex(file.name, "/"); index != -1 {
		logdir = file.name[0:index] + "/"
		os.MkdirAll(file.name[0:index], os.ModePerm)
	}

	now := time.Now()
	filename := fmt.Sprintf("%s_%04d%02d%02d_%02d%02d", file.name, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	if err := os.Rename(file.name, filename); err == nil {
		go func() {
			tarCmd := exec.Command("tar", "-zcf", filename+".tar.gz", filename, "--remove-files")
			tarCmd.Run()

			rmCmd := exec.Command("/bin/sh", "-c", "find "+logdir+` -type f -mtime +2 -exec rm {} \;`)
			rmCmd.Run()
		}()
	}

	for index := 0; index < 10; index++ {
		if fd, err := os.OpenFile(file.name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeExclusive); nil == err {
			file.fd.Sync()
			file.fd.Close()
			file.fd = fd
			break
		}

		file.fd = nil
	}
}