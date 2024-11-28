package logger

import (
	"bytes"
	"fmt"
	"github.com/577683719/common/common/util/common_util/date_util"
	"github.com/duke-git/lancet/v2/datetime"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger
var hook *LevelFileHook

func init() {
	ZeroPointExecution(func() {
		LogrusObj.Infof("每天凌晨执行:%s", date_util.DateToDateTimeStr(time.Now()))
		initialize()
		LogrusObj.Infof("日志切割成功")
	})

	// 实例化
	initialize()
}

// 零点进行日志切割
func ZeroPointExecution(invoke func()) {
	ticker := time.NewTicker(CalculateDurationUntilTomorrow())
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("凌晨日志切割-调用开始")
				invoke()
				fmt.Println("凌晨日志切割-调用结束")
				//重新计时
				tomorrow := CalculateDurationUntilTomorrow()
				ticker.Reset(tomorrow)
				fmt.Println("凌晨日志切割-重新计时")
			}
		}
	}()

}

// 获取明天0点到现在的时间Duration
func CalculateDurationUntilTomorrow() time.Duration {
	// 获取当前时间
	now := time.Now()
	return datetime.BeginOfDay(datetime.AddDay(time.Now(), 1)).Sub(now)
}

func initialize() {
	if LogrusObj != nil {
		hook.Close()
	}
	logger := logrus.New()
	logger.SetReportCaller(true)
	//src, _ := setOutputFile()
	// 	 设置输出到控制台
	logger.SetLevel(logrus.DebugLevel)
	//logger.Out = io.MultiWriter(os.Stdout, src)
	logger.Out = os.Stdout
	// 设置日志级别
	hook = NewLevelFileHook()
	logger.AddHook(hook)

	logger.SetFormatter(&StyleFormatter{})
	LogrusObj = logger
	LogrusObj.Infof("日志初始化成功")
}
func reLogger() {
	if LogrusObj != nil {
		hook.Close()
	}

	// 实例化
	logger := logrus.New()
	logger.SetReportCaller(true)
	//src, _ := setOutputFile()
	// 	 设置输出到控制台
	logger.SetLevel(logrus.DebugLevel)

	//logger.Out = io.MultiWriter(os.Stdout, src)
	// 设置日志级别
	hook = NewLevelFileHook()
	logger.AddHook(hook)

	logger.SetFormatter(&StyleFormatter{})
	LogrusObj = logger
}

// Close 方法用于关闭所有文件
func (hook *LevelFileHook) Close() {
	hook.closeFile(hook.InfoFile)
	hook.closeFile(hook.ErrorFile)
	hook.closeFile(hook.DebugFile)
	hook.closeFile(hook.WarnFile)
}

// closeFile 方法用于关闭单个文件
func (hook *LevelFileHook) closeFile(file *os.File) {
	if file != nil {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}
}

// LevelFileHook 自定义钩子结构
type LevelFileHook struct {
	InfoFile  *os.File
	ErrorFile *os.File
	DebugFile *os.File
	WarnFile  *os.File
}

// NewLevelFileHook 创建并返回一个 LevelFileHook 实例
func NewLevelFileHook() *LevelFileHook {
	now := time.Now()
	logFilePath := "logs/"
	os.MkdirAll(logFilePath, 0755)

	infoFile, _ := os.OpenFile(path.Join(logFilePath, now.Format("2006-01-02")+"-info.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	errorFile, _ := os.OpenFile(path.Join(logFilePath, now.Format("2006-01-02")+"-error.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	debugFile, _ := os.OpenFile(path.Join(logFilePath, now.Format("2006-01-02")+"-debug.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	warnFile, _ := os.OpenFile(path.Join(logFilePath, now.Format("2006-01-02")+"-warn.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)

	return &LevelFileHook{
		InfoFile:  infoFile,
		ErrorFile: errorFile,
		DebugFile: debugFile,
		WarnFile:  warnFile,
	}
}

// Fire 实现钩子的 Fire 方法
func (hook *LevelFileHook) Fire(entry *logrus.Entry) error {
	switch entry.Level {
	case logrus.InfoLevel:
		formatter := StyleFormatter{}
		format, _ := formatter.FormatOS(entry)
		fmt.Fprint(os.Stdout, string(format))
		entry.Logger.Out = hook.InfoFile
	case logrus.ErrorLevel:
		formatter := StyleFormatter{}
		format, _ := formatter.FormatOS(entry)
		fmt.Fprint(os.Stdout, string(format))
		entry.Logger.Out = hook.ErrorFile
	case logrus.DebugLevel:
		formatter := StyleFormatter{}
		format, _ := formatter.FormatOS(entry)
		fmt.Fprint(os.Stdout, string(format))
		entry.Logger.Out = hook.DebugFile
	case logrus.WarnLevel:
		formatter := StyleFormatter{}
		format, _ := formatter.FormatOS(entry)
		fmt.Fprint(os.Stdout, string(format))
		entry.Logger.Out = hook.WarnFile
	default:
		entry.Logger.Out = hook.InfoFile
	}
	return nil
}

// Levels 实现钩子的 Levels 方法
func (hook *LevelFileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// billing_name
func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}

func shortenFileName(file string) string {
	segments := strings.Split(file, "/")
	return segments[len(segments)-1]
}

type StyleFormatter struct{}

func (f *StyleFormatter) FormatOS(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 获取调用者信息，2表示日志记录器上两层的调用者
	pc, file, line, ok := runtime.Caller(9) // 调整这里的层数以获取合适的调用者信息
	if !ok {
		file = "???"
		line = 0
	}

	// 获取函数名
	fn := runtime.FuncForPC(pc).Name()
	shortFile := shortenFileName(file)

	// 构建类似Java的日志条目
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	level := strings.ToUpper(entry.Level.String())
	b.WriteString(fmt.Sprintf("%s [%s] %s(%s:%d) - %s\n",
		timestamp, level, fn, shortFile, line, entry.Message))

	return b.Bytes(), nil
}
func (f *StyleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 获取调用者信息，2表示日志记录器上两层的调用者
	pc, file, line, ok := runtime.Caller(7) // 调整这里的层数以获取合适的调用者信息
	if !ok {
		file = "???"
		line = 0
	}

	// 获取函数名
	fn := runtime.FuncForPC(pc).Name()
	shortFile := shortenFileName(file)

	// 构建类似Java的日志条目
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	level := strings.ToUpper(entry.Level.String())
	b.WriteString(fmt.Sprintf("%s [%s] %s(%s:%d) - %s\n",
		timestamp, level, fn, shortFile, line, entry.Message))

	return b.Bytes(), nil
}
