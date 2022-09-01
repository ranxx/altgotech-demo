package log

import (
	"io"
	"os"
	"runtime"
	"strconv"

	"github.com/ranxx/altgotech-demo/config"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                  //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return "", f.File + ":" + strconv.Itoa(f.Line)
		},
	})
	log.SetReportCaller(true)
	log.SetOutput(NewWriter())
}

func Logrus() *logrus.Logger {
	return log
}

type Writer struct {
	File io.Writer
}

func NewWriter() *Writer {
	return &Writer{
		File: config.GetLogio(),
	}
}

func (w *Writer) WriteString(s string) (n int, err error) {
	if w.File == nil || w.File == os.Stdout { // 防止写两个日志
		return w.File.Write([]byte(s))
	}
	w.File.Write([]byte(s))
	return os.Stdout.WriteString(s)
}

func (w *Writer) Write(b []byte) (n int, err error) {
	if w.File == nil || w.File == os.Stdout { // 防止写两个日志
		return w.File.Write(b)
	}
	w.File.Write(b)
	return os.Stdout.Write(b)
}
