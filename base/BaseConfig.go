package base

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mysql"
	"github.com/astaxie/beego/logs"
)

// 数据库配置
var DBSource = beego.AppConfig.String("dbsource")
var DBType = beego.AppConfig.String("dbtype")
var DBConnectstr = beego.AppConfig.String("dbconnectstr")
var MaxIdleConns, _ = beego.AppConfig.Int("maxidleconns")
var MaxOpenConns, _ = beego.AppConfig.Int("maxopenconns")

func init() {
	// 数据库配置
	orm.RegisterDriver(DBType, orm.DRMySQL)
	orm.RegisterDataBase(DBSource, DBType, DBConnectstr)
	orm.SetMaxIdleConns(DBSource, MaxIdleConns)
	orm.SetMaxOpenConns(DBSource, MaxOpenConns)

	// 日志配置
	logs.SetLogger(logs.AdapterConsole, `{"level":1}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"`+beego.AppConfig.String("appname")+`.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Async()
}
