package main

import (
	"time"

	"github.com/angus/unit1/mylog/mylogger"
)

//声明全局接口变量
var log mylogger.Logger

//demo
//测试日志库
func main() {
	log = mylogger.NewConsoleLogger("Info")                              //终端日志实例
	log = mylogger.NewFileLogger("Info", "./", "test.log", 10*1024*1024) //文件日志实例

	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条Warning日志")
		id := 2021
		name := "China NO.1"
		log.Error("这是一条Error日志,id：%d,name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		time.Sleep(2 * time.Second)
	}

}
