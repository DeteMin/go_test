package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

func main(){
	//timer1 := time.NewTimer(10 * time.Second)
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05"))
	//
	//go func(t *time.Timer) {
	//	times := 0
	//	for {
	//		<-t.C
	//		fmt.Println("timer", time.Now().Format("2006-01-02 15:04:05"))
	//
	//		// 从t.C中获取数据，此时time.Timer定时器结束。如果想再次调用定时器，只能通过调用 Reset() 函数来执行
	//		// Reset 使 t 重新开始计时，（本方法返回后再）等待时间段 d 过去后到期。
	//		// 如果调用时 t 还在等待中会返回真；如果 t已经到期或者被停止了会返回假。
	//		times++
	//		// 调用 reset 重发数据到chan C
	//		fmt.Println("调用 reset 重新设置一次timer定时器，并将时间修改为2秒")
	//		t.Reset(2 * time.Second)
	//		if times > 3 {
	//			fmt.Println("调用 stop 停止定时器")
	//			t.Stop()
	//			wg.Done()
	//		}
	//	}
	//}(timer1)
	//wg.Wait()
	//
	//l1 := []int32{2,5}
	//l2 := []int32{2,5}
	//b := StringSliceEqualBCE(l1, l2)
	//if b {
	//	fmt.Println("相同")
	//} else {
	//	fmt.Println("不想等")
	//}

	//var a float32 = 2263.99999
	//var b int64
	//b = int64(a)
	//fmt.Println(fmt.Sprintf("%.1f", a))

	//var t1 int64 = 1615877920
	//_,_, d,_,_,_ := TimeDiff(time.Unix(t1, 0), time.Now())
	//fmt.Println("相差%d", d)

	Decimal("14.8997")

}

func StringSliceEqualBCE(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TimeDiff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func Decimal(data string) {
	yh, _ := decimal.NewFromString(data)
	y := yh.Round(3).StringFixed(3)
	fmt.Printf("parse:%v", y)
}