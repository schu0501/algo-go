package main

import (
	"fmt"
	"sync"
	"time"
)

//func main() {
//	maxWorkers := 3
//	maxCount := 10
//
//	var wg sync.WaitGroup
//	ch := make(chan struct{}, maxWorkers)
//
//	for i := 0; i < maxCount; i++ {
//		wg.Add(1)
//		ch <- struct{}{}
//		go func(j int) {
//			defer wg.Done()
//			//归还令牌
//			defer func() { <-ch }()
//
//			fmt.Printf("Worker %d processing task\n", j)
//			time.Sleep(1 * time.Second)
//		}(i)
//	}
//
//	wg.Wait()
//	fmt.Printf("All tasks processed\n")
//}

type Task func()

type Pool struct {
	workQueue  chan Task
	wg         sync.WaitGroup
	maxWorkers int
}

func NewPool(maxWorkers, queueSize int) *Pool {
	return &Pool{
		workQueue:  make(chan Task, queueSize),
		maxWorkers: maxWorkers,
	}
}

func (p *Pool) Start() {
	for i := 0; i < p.maxWorkers; i++ {
		p.wg.Add(1)
		go p.worker(i)
	}
}

func (p *Pool) worker(WorkerID int) {
	defer p.wg.Done()

	for task := range p.workQueue {
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("Worker %d panic: %v\n", WorkerID, err)
				}
			}()
			task()
		}()
	}
	fmt.Printf("Worker %d stopped\n", WorkerID)
}

func (p *Pool) Submit(task Task) {
	p.workQueue <- task
}

func (p *Pool) Shutdown() {
	close(p.workQueue)
	p.wg.Wait()
	fmt.Println("All workers stopped")
}

func main() {
	pool := NewPool(2, 10)
	pool.Start()

	for i := 0; i < 10; i++ {
		taskId := i
		pool.Submit(func() {
			fmt.Printf("Task %d is running\n", taskId)
			time.Sleep(1 * time.Second)

			if taskId == 5 {
				panic("oops")
			}
		})
	}

	pool.Shutdown()
}
