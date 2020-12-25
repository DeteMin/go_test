package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	ch = make(chan int)
)

func main(){

	//TestMutex()
	//TestRWMutex()
	TestWaitGroup()
	time.Sleep(time.Second * 10)
}

//互斥锁
func TestMutex(){
	var (
		TestNum = 10
	)

	lock := sync.Mutex{}
	go func() {
		lock.Lock()
		defer lock.Unlock()

		fmt.Println(fmt.Sprintf("=== goroutine-1:star TestNum:%v===", TestNum))
		TestNum += 1
		time.Sleep(5 * time.Second)
		fmt.Println(fmt.Sprintf("=== goroutine-1:end TestNum:%v", TestNum))
	}()

	//启动goroutine-1后睡 1s
	time.Sleep(time.Second * 1)
	go func() {
		lock.Lock()
		lock.Unlock()

		fmt.Println(fmt.Sprintf("=== goroutine-2:star TestNum:%v===", TestNum))
		TestNum -= 2
		fmt.Println(fmt.Sprintf("=== goroutine-2:end TestNum:%v", TestNum))
	}()
}

//读写锁
func TestRWMutex(){
	var (
		TestNum = 10
		NoLockNum = 20
	)

	lock := sync.RWMutex{}
	go func() {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println(fmt.Sprintf("=== goroutine-1-star TestNum:%v	NoLockNum:%v", TestNum, NoLockNum))
		time.Sleep(5 * time.Second)
		TestNum += 1

		fmt.Println(fmt.Sprintf("=== goroutine-1-star TestNum:%v	NoLockNum:%v", TestNum, NoLockNum))
	}()

	time.Sleep(1 * time.Second)

	go func() {
		lock.RLock()
		defer lock.RUnlock()
		//读写锁，应该能正常读
		fmt.Println(fmt.Sprintf("=== goroutine-2-star TestNum:%v	NoLockNum:%v", TestNum, NoLockNum))
		TestNum -= 2
		fmt.Println(fmt.Sprintf("=== goroutine-2-star TestNum:%v	NoLockNum:%v", TestNum, NoLockNum))
	}()
}

//channel锁用法
func TestChannel(){
	n := <-ch
	fmt.Println(fmt.Sprintf("=== TestChannel-star getNum:%v", n))

}

func TestWaitGroup(){
	var (
		arr = []int32{1,2,3}
		wg = &sync.WaitGroup{}
	)

	wg.Add(len(arr))
	for _, v := range arr {
		go func(i int32) {
			fmt.Println(i)
			time.Sleep(time.Duration(i) * time.Second)
			defer wg.Done()
		}(v)
	}
	//等待所有协程完成
	wg.Wait()
	fmt.Println("WaitGroup done")
}
