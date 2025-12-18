package main

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

func TopKFrequentElements(nums []int, k int) []int {
	freMap := make(map[int]int)
	for _, n := range nums {
		freMap[n]++
	}

	keys := make([]int, 0, len(freMap))
	for key := range freMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return freMap[keys[i]] > freMap[keys[j]]
	})

	return keys[:k]
}

type User struct {
	name string `json:"name,omitempty"`
	age  int    `json:"age,omitempty"`
}

func main() {

	//nums := []int{1, 4, 5, 6, 2, 3, 5, 5, 2, 7, 3, 3, 4, 4, 2, 6, 8, 8}
	//fmt.Println(TopKFrequentElements(nums, 5))

	//stackInt := &stack.Stack{1, 3}
	//stackInt.Push(4)
	//fmt.Println(*stackInt)

	//h := &myHeap.IntHeap{1, 3, 4, 5, 6, 7}
	//heap.Init(h)
	//heap.Push(h, 2)
	//fmt.Println(h)
	//
	//fmt.Println("return:", f())
	//myDB := &database.MySQL{
	//	ConnString: "root:@tcp(127.0.0.1:3306)/test?charset=utf8",
	//}
	//
	//srv := service.OrderService{
	//	Saver: myDB,
	//}
	//
	//srv.CreateOrder("")
	count := 10
	sum := 100
	wg := sync.WaitGroup{}
	c := make(chan struct{}, count)
	defer close(c)
	for i := 0; i < sum; i++ {
		wg.Add(1)
		c <- struct{}{}
		go func(j int) {
			defer wg.Done()
			fmt.Println(j)
			<-c
		}(i)
	}
	wg.Wait()
}

func getStructTag(f reflect.StructField) string {
	return string(f.Tag)
}

func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
