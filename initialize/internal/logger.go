package internal

import (
	"fmt"

	"github.com/oldweipro/gin-admin/global"
	"gorm.io/gorm/logger"
)

type Writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *Writer {
	return &Writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *Writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.ConfigServer.System.DbType {
	case "mysql":
		logZap = global.ConfigServer.Mysql.LogZap
	case "pgsql":
		logZap = global.ConfigServer.Pgsql.LogZap
	}
	if logZap {
		global.Logger.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
