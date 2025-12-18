package linkedList

import (
	"errors"
	"fmt"
)

// ListNode 单链表
type ListNode[E any] struct {
	Val  E
	Next *ListNode[E]
}
type LinkedList[E any] struct {
	head *ListNode[E]
	tail *ListNode[E]
	size int
}

// NewLinkedList 创建一个新的链表
func NewLinkedList[E any]() *LinkedList[E] {
	head := &ListNode[E]{}
	return &LinkedList[E]{head: head, tail: head, size: 0}
}

func NewLinkedListFromArray[E any](arr []E) *LinkedList[E] {
	head := &ListNode[E]{}
	tail := head
	for _, val := range arr {
		newNode := &ListNode[E]{Val: val}
		tail.Next = newNode
		tail = newNode
	}
	return &LinkedList[E]{head: head, tail: tail, size: len(arr)}
}

// AddFirst 在头部添加元素
func (list *LinkedList[E]) AddFirst(e E) {
	newListNode := &ListNode[E]{Val: e}
	newListNode.Next = list.head.Next
	list.head.Next = newListNode
	if list.size == 0 {
		list.tail = newListNode
	}
	list.size++
}

// AddLast 在尾部添加元素
func (list *LinkedList[E]) AddLast(e E) {
	newListNode := &ListNode[E]{Val: e}
	list.tail.Next = newListNode
	list.tail = newListNode
	list.size++
}

// Add 在指定索引处添加元素
func (list *LinkedList[E]) Add(index int, element E) error {
	if index < 0 || index > list.size {
		return errors.New("index out of range")
	}
	// 优化：利用已有的 AddLast 处理尾部插入
	// 这很重要，因为如果 index == size，下面的遍历逻辑需要特殊处理
	if index == list.size {
		list.AddLast(element)
		return nil
	}

	// 关键修正：从 dummy head 开始遍历，找到 index 的前驱节点 (pred)
	// 我们要找 index-1，所以从 head 开始走 index 步即可
	// i=0: head -> pred (index 0 的前驱)
	pred := list.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}

	newNode := &ListNode[E]{Val: element}
	newNode.Next = pred.Next
	pred.Next = newNode

	list.size++
	return nil
}

// RemoveFirst 移除头部元素
func (list *LinkedList[E]) RemoveFirst() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("list is empty")
	}
	x := list.head.Next.Val
	list.head.Next = list.head.Next.Next
	if list.size == 1 {
		list.tail = list.head
	}
	list.size--
	return x, nil
}

// RemoveLast 移除尾部元素
func (list *LinkedList[E]) RemoveLast() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("list is empty")
	}
	tmp := list.head
	for tmp.Next != list.tail {
		tmp = tmp.Next
	}
	x := list.tail
	tmp.Next = nil
	list.tail = tmp
	list.size--
	return x.Val, nil
}

// Remove 在指定索引处移除元素
func (list *LinkedList[E]) Remove(index int) (E, error) {
	if index < 0 || index >= list.size {
		return *new(E), fmt.Errorf("index out of range")
	}
	// 关键修正：找到待删除节点的前驱 (pred)
	pred := list.head
	for i := 0; i < index; i++ {
		pred = pred.Next
	}

	// 待删除的节点
	toRemove := pred.Next
	val := toRemove.Val

	// 删除操作
	pred.Next = toRemove.Next

	// 关键修正：如果你删除了最后一个节点，需要更新 tail 指针！
	// 此时 pred 变成了新的 tail
	if index == list.size-1 {
		list.tail = pred
	}

	// 帮助 GC (可选)
	toRemove.Next = nil

	list.size--
	return val, nil
}

// GetFirst 获取头部元素
func (list *LinkedList[E]) GetFirst() (E, error) {
	if list.IsEmpty() {
		return *new(E), fmt.Errorf("list is empty")
	}
	return list.head.Next.Val, nil
}

// GetLast 获取尾部元素
func (list *LinkedList[E]) GetLast() (E, error) {
	if list.IsEmpty() {
		return *new(E), errors.New("list is empty")
	}
	return list.tail.Val, nil
}

// Get 获取指定索引的元素
func (list *LinkedList[E]) Get(index int) (E, error) {
	if index < 0 || index >= list.size {
		return *new(E), fmt.Errorf("index out of range")
	}
	return list.getNode(index).Val, nil
}

// Set 更新指定索引的元素
func (list *LinkedList[E]) Set(index int, element E) (E, error) {
	if index < 0 || index >= list.size {
		return *new(E), fmt.Errorf("index out of range")
	}

	p := list.getNode(index)
	old := p.Val
	p.Val = element
	return old, nil
}

// Size 获取链表大小
func (list *LinkedList[E]) Size() int {
	return list.size
}

// IsEmpty 检查链表是否为空
func (list *LinkedList[E]) IsEmpty() bool {
	return list.size <= 0
}

// getNode 返回指定索引的节点
func (list *LinkedList[E]) getNode(index int) *ListNode[E] {
	p := list.head.Next
	for i := 0; i < index; i++ {
		if p == nil {
			return nil
		}
		p = p.Next
	}
	return p
}

func (list *LinkedList[T]) Display() {
	fmt.Printf("size = %d\n", list.size)
	p := list.head.Next
	for p != list.tail {
		fmt.Printf("%v -> ", p.Val)
		p = p.Next
	}
	fmt.Println("null\n")
}
