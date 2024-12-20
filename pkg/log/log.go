package log

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Log struct {
	init bool
	slog *zap.SugaredLogger
}

var (
	pid    string
	once   sync.Once
	logger Log
)

func init() {
	pid = strconv.Itoa(os.Getpid())
	logger.init = false
}

func InitLogger(level string, path string) {
	once.Do(func() {
		logger.slog = newCustomLogger(level, path).Sugar()
		logger.init = true
	})
}

func IncreaseLogLevel(loglevel string) {

	if !logger.init {
		return
	}

	// 设置日志级别 debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	logger.slog = logger.slog.WithOptions(zap.IncreaseLevel(level))
}

func ReleaseLogger() {
	if !logger.init {
		return
	}
	logger.slog.Sync()
}

func Debug(template string, args ...interface{}) {
	if logger.init {
		logger.debug(template, args...)
	}
}

func Info(template string, args ...interface{}) {
	if logger.init {
		logger.info(template, args...)
	}
}

func Warn(template string, args ...interface{}) {
	if logger.init {
		logger.warn(template, args...)
	}
}

func Error(template string, args ...interface{}) {
	if logger.init {
		logger.error(template, args...)
	}
}

func Panic(template string, args ...interface{}) {
	if logger.init {
		logger.panic(template, args...)
	}
}

func Fatal(template string, args ...interface{}) {
	if logger.init {
		logger.fatal(template, args...)
	}
}

func newCustomLogger(loglevel string, logpath string) *zap.Logger {

	// 设置日志级别 debug->info->warn->error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:       "msg",
		TimeKey:          "time",
		NameKey:          "logger",
		CallerKey:        "file",
		FunctionKey:      "function",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeName:       zapcore.FullNameEncoder,
		ConsoleSeparator: " ",
	}

	// 如果logpath为空，则只打印到控制台
	if logpath == "" {
		return zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), level))
	}

	// 日志分割
	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径，默认 os.TempDir()
		MaxSize:    10,      // 每个日志文件保存10M，默认 100M
		MaxBackups: 5,       // 保留30个备份，默认不限
		MaxAge:     30,      // 保留7天，默认不限
		Compress:   false,   // 是否压缩，默认不压缩
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		level,
	)

	// 构造日志
	logger := zap.New(core)
	return logger
}

/////////////////////////////////////////

func (l *Log) debug(template string, args ...interface{}) {

	file := "unknow"
	line := 0

	_, file, line, _ = runtime.Caller(2)

	file_slice := strings.Split(file, "/")

	template = "[Debug] [PID:" + pid + "] [" + file_slice[len(file_slice)-1] + ":" + strconv.Itoa(line) + "] " + template
	l.slog.Debugf(template, args...)
}

func (l *Log) info(template string, args ...interface{}) {

	file := "unknow"
	line := 0

	_, file, line, _ = runtime.Caller(2)

	file_slice := strings.Split(file, "/")

	template = "[Info] [PID:" + pid + "] [" + file_slice[len(file_slice)-1] + ":" + strconv.Itoa(line) + "] " + template
	l.slog.Infof(template, args...)
}

func (l *Log) warn(template string, args ...interface{}) {

	file := "unknow"
	line := 0

	_, file, line, _ = runtime.Caller(2)

	file_slice := strings.Split(file, "/")

	template = "[Warn] [PID:" + pid + "] [" + file_slice[len(file_slice)-1] + ":" + strconv.Itoa(line) + "] " + template
	l.slog.Warnf(template, args...)
}

func (l *Log) error(template string, args ...interface{}) {

	file := "unknow"
	line := 0

	_, file, line, _ = runtime.Caller(2)

	file_slice := strings.Split(file, "/")

	template = "[Error] [PID:" + pid + "] [" + file_slice[len(file_slice)-1] + ":" + strconv.Itoa(line) + "] " + template
	l.slog.Errorf(template, args...)
}

func (l *Log) panic(template string, args ...interface{}) {

	file := "unknow"
	line := 0

	_, file, line, _ = runtime.Caller(2)

	file_slice := strings.Split(file, "/")

	template = "[Panic] [PID:" + pid + "] [" + file_slice[len(file_slice)-1] + ":" + strconv.Itoa(line) + "] " + template
	l.slog.Panicf(template, args...)
}

func (l *Log) fatal(template string, args ...interface{}) {

	file := "unknow"
	line := 0

	_, file, line, _ = runtime.Caller(2)

	file_slice := strings.Split(file, "/")

	template = "[Fatal] [PID:" + pid + "] [" + file_slice[len(file_slice)-1] + ":" + strconv.Itoa(line) + "] " + template
	l.slog.Fatalf(template, args...)
}
