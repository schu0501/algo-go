package main

import (
	"algo-go/common/linkedList"
	"fmt"
)

func main() {
	arr := make([]int, 10)
	for i := 0; i < 4; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
	//arr = append(arr[:2], append([]int{666}, arr[2:]...)...)
	//更高性能的写法（零额外内存分配）：
	//如果你在写 Web3 的底层共识逻辑或高频交易撮合（对 GC 敏感），
	//推荐用下面这种“原地挪动”的写法，虽然代码长了点，但没有临时内存分配：
	arr = append(arr, 0)
	fmt.Println(arr)
	copy(arr[2:], arr[1:])
	fmt.Println(arr)
	arr[1] = 666
	fmt.Println(arr)
	//fmt.Println(arr[2:4])
	listNode := linkedList.NewLinkedListFromArray(arr)
	listNode.Display()
	_, err := listNode.RemoveFirst()
	if err != nil {
		return
	}
	listNode.Display()
	doubleLinkedList := linkedList.NewDoubleLinkedListFromArr(arr)
	doubleLinkedList.Display()
}
