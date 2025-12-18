package linkedList

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	Val        T
	Prev, Next *Node[T]
}
type DoubleLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	// 1. 创建两个哨兵节点
	head := &Node[T]{}
	tail := &Node[T]{}

	// 2. 将它们互连： head <-> tail
	head.Next = tail
	tail.Prev = head

	// 3. 返回链表结构
	return &DoubleLinkedList[T]{
		head: head,
		tail: tail,
		size: 0,
	}
}
func NewDoubleLinkedListFromArr[T any](arr []T) *DoubleLinkedList[T] {
	list := NewDoubleLinkedList[T]()
	if arr == nil || len(arr) == 0 {
		return list
	}

	for _, val := range arr {
		list.AddLast(val)
	}
	return list
}

func (list *DoubleLinkedList[T]) AddLast(val T) {
	newNode := &Node[T]{Val: val}
	pre := list.tail.Prev
	newNode.Prev = pre
	newNode.Next = list.tail
	pre.Next = newNode
	list.tail.Prev = newNode
	list.size++
}

func (list *DoubleLinkedList[T]) AddFirst(val T) {
	newNode := &Node[T]{Val: val}
	tmp := list.head.Next

	newNode.Next = tmp
	newNode.Prev = tmp.Prev
	tmp.Prev = newNode
	list.head.Next = newNode
	list.size++
}

func (list *DoubleLinkedList[T]) Add(index int, val T) error {
	// 检查范围 [0, size]
	if err := list.checkPositionIndex(index); err != nil {
		return err
	}

	// 1. 找到“后继节点” (successor)
	// 此时 getNode(index) 非常好用：
	// 如果 index == size，它刚好返回 tail (哨兵)
	// 如果 index == 0，它刚好返回 head.Next (第一个数据)
	succ := list.getNode(index)

	// 2. 找到“前驱节点” (predecessor)
	pred := succ.Prev

	// 3. 在 pred 和 succ 之间连接新节点
	newNode := &Node[T]{Val: val}

	newNode.Prev = pred
	newNode.Next = succ

	pred.Next = newNode
	succ.Prev = newNode

	list.size++
	return nil
}

func (list *DoubleLinkedList[T]) GetFirst() (T, error) {
	if list.size < 1 {
		var zero T
		return zero, errors.New("list is empty")
	}
	return list.head.Next.Val, nil
}

func (list *DoubleLinkedList[T]) GetLast() (T, error) {
	if list.size < 1 {
		var zero T
		return zero, errors.New("list is empty")
	}
	return list.tail.Prev.Val, nil
}

func (list *DoubleLinkedList[T]) Get(index int) (T, error) {
	if err := list.checkElementIndex(index); err != nil {
		var zero T
		return zero, err
	}
	return list.getNode(index).Val, nil
}

func (list *DoubleLinkedList[T]) RemoveFirst() (T, error) {
	if list.size < 1 {
		var zero T
		return zero, errors.New("list is empty")
	}
	x := list.head.Next
	tmp := x.Next
	tmp.Prev = list.head
	list.head.Next = tmp
	list.size--
	return x.Val, nil
}

func (list *DoubleLinkedList[T]) RemoveLast() (T, error) {
	if list.size < 1 {
		var zero T
		return zero, errors.New("list is empty")
	}
	x := list.tail.Prev
	tmp := x.Prev
	tmp.Next = list.tail
	list.tail.Prev = tmp
	list.size--
	return x.Val, nil
}

func (list *DoubleLinkedList[T]) Remove(index int) (T, error) {
	if err := list.checkElementIndex(index); err != nil {
		var zero T
		return zero, err
	}
	x := list.getNode(index)
	next := x.Next
	prev := x.Prev

	next.Prev = prev
	prev.Next = next
	list.size--
	return x.Val, nil
}

func (list *DoubleLinkedList[T]) Set(index int, val T) (T, error) {
	if err := list.checkElementIndex(index); err != nil {
		var zero T
		return zero, err
	}
	p := list.getNode(index)
	oldVal := p.Val
	p.Val = val
	return oldVal, nil
}

func (list *DoubleLinkedList[T]) getNode(index int) *Node[T] {
	// 如果 index 小于 size 的一半，从前往后找
	if index < list.size/2 {
		p := list.head.Next
		for i := 0; i < index; i++ {
			p = p.Next
		}
		return p
	}

	// 否则，从后往前找
	// 注意：我们要找的是 index 位置。
	// tail 是哨兵。tail.Prev 是 size-1。
	p := list.tail
	// 如果 index == size (用于 Add)，直接返回 tail，不用循环
	for i := list.size; i > index; i-- {
		p = p.Prev
	}
	return p
}

func (list *DoubleLinkedList[T]) isPositionIndex(index int) bool {
	return index >= 0 && index <= list.size
}

func (list *DoubleLinkedList[T]) isElementIndex(index int) bool {
	return index >= 0 && index < list.size
}

func (list *DoubleLinkedList[T]) checkElementIndex(index int) error {
	if !list.isElementIndex(index) {
		return fmt.Errorf("index: %d, Size: %d", index, list.size)
	}
	return nil
}

func (list *DoubleLinkedList[T]) checkPositionIndex(index int) error {
	if !list.isPositionIndex(index) {
		return fmt.Errorf("index: %d, Size: %d", index, list.size)
	}
	return nil
}

func (list *DoubleLinkedList[T]) Display() {
	fmt.Printf("size = %d\n", list.size)
	p := list.head.Next
	for p != list.tail {
		fmt.Printf("%v <-> ", p.Val)
		p = p.Next
	}
	fmt.Println("null\n")
}

func (list *DoubleLinkedList[T]) Size() int {
	return list.size
}

func (list *DoubleLinkedList[T]) IsEmpty() bool {
	return list.size == 0
}
