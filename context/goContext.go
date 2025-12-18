package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	//
	//doBusiness(ctx)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("done")
				return
			default:
				fmt.Println("正在监控...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("主程序决定停止监控")
	cancel()

	time.Sleep(1 * time.Second)
}

func doBusiness(ctx context.Context) {
	done := make(chan struct{})

	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	select {
	case <-done:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("任务失败", ctx.Err())
	}
}
