package main

import (
	"fmt"
	"sync"
)

// 定义最大打印到哪个数字
const maxNumber = 100

// worker 函数
//func worker(name string, start int, in chan struct{}, out chan struct{}, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	// 循环变量 i 每次 +3
//	for i := start; ; i += 3 {
//		// 1. 等待接收信号
//		// ok 为 false 说明 channel 已经被上游关闭了，我们也应该退出
//		_, ok := <-in
//		if !ok {
//			close(out) // 传递关闭信号给下游
//			return
//		}
//
//		// 3. 判断是否结束
//		// 如果下一个数字超过了 maxNumber，说明我是最后一个打印的人
//		// 我需要负责发起“关闭链条”的动作
//		if i > maxNumber {
//			close(out) // 关闭下游通道，通知下一个协程该结束了
//			return
//		}
//		// 2. 打印数字
//		fmt.Printf("Goroutine %s: %d\n", name, i)
//
//		// 4. 正常传递信号给下一个协程
//		out <- struct{}{}
//	}
//}

func main() {
	//var wg sync.WaitGroup
	//wg.Add(3)
	//
	//// 创建三个无缓冲通道
	//// 拓扑结构: ch1 -> G1 -> ch2 -> G2 -> ch3 -> G3 -> ch1
	//ch1 := make(chan struct{})
	//ch2 := make(chan struct{})
	//ch3 := make(chan struct{})
	//
	//// 启动 G1: 1, 4, 7...
	//go worker("G1", 1, ch1, ch2, &wg)
	//// 启动 G2: 2, 5, 8...
	//go worker("G2", 2, ch2, ch3, &wg)
	//// 启动 G3: 3, 6, 9...
	//go worker("G3", 3, ch3, ch1, &wg)
	//
	//// 启动发令枪：给 ch1 发送第一个信号，让 G1 开始运行
	//ch1 <- struct{}{}
	//
	//// 等待所有协程结束
	//wg.Wait()
	//fmt.Println("所有打印任务完成")

	var wg sync.WaitGroup
	wg.Add(3)

	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})

	go worker("G1", 1, ch1, ch2, &wg)
	go worker("G2", 2, ch2, ch3, &wg)
	go worker("G3", 3, ch3, ch1, &wg)

	ch1 <- struct{}{}
	wg.Wait()
}

func worker(name string, start int, from chan struct{}, to chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; ; i += 3 {
		_, ok := <-from
		if !ok {
			close(to)
			return
		}

		if i > maxNumber {
			close(to)
			return
		}
		fmt.Printf("Goroutine %s: %d\n", name, i)
		to <- struct{}{}
	}
}
