package main

import (
	"fmt"
	"time"
)

// 不同情况下，Timer.Reset()的返回值
func test1() {
	fmt.Println("第1个测试：Reset返回值和什么有关？")
	tm := time.NewTimer(time.Second)
	defer tm.Stop()

	quit := make(chan bool)

	// 退出事件
	go func() {
		time.Sleep(3 * time.Second)
		quit <- true
	}()

	// Timer未超时，看Reset的返回值
	if !tm.Reset(time.Second) {
		fmt.Println("未超时，Reset返回false")
	} else {
		fmt.Println("未超时，Reset返回true")
	}

	// 停止timer
	tm.Stop()
	if !tm.Reset(time.Second) {
		fmt.Println("停止Timer，Reset返回false")
	} else {
		fmt.Println("停止Timer，Reset返回true")
	}

	// Timer超时
	for {
		select {
		case <-quit:
			return

		case <-tm.C:
			if !tm.Reset(time.Second) {
				fmt.Println("超时，Reset返回false")
			} else {
				fmt.Println("超时，Reset返回true")
			}
		}
	}
}

func test2() {
	fmt.Println("\n第2个测试:超时后，不读通道中的事件，可以Reset成功吗？")
	sm2Start := time.Now()
	tm2 := time.NewTimer(time.Second)
	time.Sleep(2 * time.Second)
	fmt.Printf("Reset前通道中事件的数量:%d\n", len(tm2.C))
	if !tm2.Reset(time.Second) {
		fmt.Println("不读通道数据，Reset返回false")
	} else {
		fmt.Println("不读通道数据，Reset返回true")
	}
	fmt.Printf("Reset后通道中事件的数量:%d\n", len(tm2.C))

	select {
	case t := <-tm2.C:
		fmt.Printf("tm2开始的时间: %v\n", sm2Start.Unix())
		fmt.Printf("通道中事件的时间：%v\n", t.Unix())
		if t.Sub(sm2Start) <= time.Second+time.Millisecond {
			fmt.Println("通道中的时间是重新设置sm2前的时间，即第一次超时的时间，所以第二次Reset失败了")
		}
	}

	fmt.Printf("读通道后，其中事件的数量:%d\n", len(tm2.C))
	tm2.Reset(time.Second)
	fmt.Printf("再次Reset后，通道中事件的数量:%d\n", len(tm2.C))
	time.Sleep(2 * time.Second)
	fmt.Printf("超时后通道中事件的数量:%d\n", len(tm2.C))
}

func test3() {
	fmt.Println("\n第3个测试：Reset前清空通道，尽可能通畅")
	smStart := time.Now()
	tm := time.NewTimer(time.Second)
	time.Sleep(2 * time.Second)

	// 停掉定时器再清空
	if !tm.Stop() && len(tm.C) > 0 {
		<-tm.C
	}
	tm.Reset(time.Second)

	// 超时
	t := <-tm.C
	fmt.Printf("tm开始的时间: %v\n", smStart.Unix())
	fmt.Printf("通道中事件的时间：%v\n", t.Unix())
	if t.Sub(smStart) <= time.Second+time.Millisecond {
		fmt.Println("通道中的时间是重新设置sm前的时间，即第一次超时的时间，所以第二次Reset失败了")
	} else {
		fmt.Println("通道中的时间是重新设置sm后的时间，Reset成功了")
	}
}


func startTimer(f func()) {
	go func() {
		fmt.Println("etw")
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

func main() {

	startTimer(func() {
		fmt.Println("test")
	})

	return
	n := time.Now()
	next := n.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())

	fmt.Println(n, next)
	fmt.Println(next.Sub(n))
	// test1()
	// test2()
	// test3()
}