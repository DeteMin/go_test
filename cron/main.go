package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main(){
	c := cron.New()

	// @every后加一个时间间隔，表示每隔多长时间触发一次
	c.AddFunc("@every 1s", func(){
		fmt.Println("tick every 1 second")
	})

	// 启动循环
	c.Start()
	time.Sleep(time.Second * 5)
}
