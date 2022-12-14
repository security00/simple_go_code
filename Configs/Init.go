package Configs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/xormplus/xorm"
	"io"
	"log"
	"mygo/Lib/CommonFunc"
	"os"
)

var Eg *xorm.EngineGroup

func init() {
	initLog()
	initMysql()
}
func initMysql() {
	var mysqlConn = CommonFunc.Conf().Mysql
	var err error
	Eg, err = xorm.NewEngineGroup("mysql", mysqlConn, xorm.RandomPolicy())
	if err != nil {
		fmt.Println("init mysql err:", err.Error())
	}
	Eg.SetMaxIdleConns(20)
	Eg.SetMaxOpenConns(25)
}

func initLog() {
	fmt.Println("main init")
	gin.SetMode(gin.DebugMode)
	gin.DisableConsoleColor()
	var f *os.File
	if _, err := os.Stat(CommonFunc.LogFileName); err != nil {
		if os.IsNotExist(err) {
			f, _ = os.Create(CommonFunc.LogFileName)
		}
	} else {
		f, _ = os.OpenFile(CommonFunc.LogFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.SetOutput(f) // 将文件设置为log输出的文件
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	logrus.SetOutput(f)
}
