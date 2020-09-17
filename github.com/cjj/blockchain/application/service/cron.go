package service

import (
	"log"

	"github.com/robfig/cron"
)

const spec = "0 0 0 * * ?" // 每天0点执行
//const spec = "*/10 * * * * ?" //10秒执行一次，用于测试

func Init() {
	c := cron.New(cron.WithSeconds()) //支持到秒级别
	//_, err := c.AddFunc(spec, GoRun)
	// if err != nil {
	// 	log.Printf("定时任务开启失败 %s", err)
	// }
	c.Start()
	log.Printf("定时任务已开启")
	select {}
}

// func GoRun() {
// 	log.Printf("定时任务已启动")
// 	//先把所有销售查询出来

// }
