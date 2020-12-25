package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	timer1 := time.NewTimer(10 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05"))

	go func(t *time.Timer) {
		times := 0
		for {
			<-t.C
			fmt.Println("timer", time.Now().Format("2006-01-02 15:04:05"))

			// 从t.C中获取数据，此时time.Timer定时器结束。如果想再次调用定时器，只能通过调用 Reset() 函数来执行
			// Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。
			// 如果调用时 t 还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
			times++
			// 调用 reset 重发数据到chan C
			fmt.Println("调用 reset 重新设置一次timer定时器，并将时间修改为2秒")
			t.Reset(2 * time.Second)
			if times > 3 {
				fmt.Println("调用 stop 停止定时器")
				t.Stop()
				wg.Done()
			}
		}
	}(timer1)
	wg.Wait()


}
