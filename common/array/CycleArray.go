package array

import (
	"errors"
	"fmt"
)

type CycleArray[T any] struct {
	arr   []T
	start int
	end   int
	count int
	size  int
}

func NewCycleArray[T any]() *CycleArray[T] {
	return NewCycleArrayWithSize[T](1)
}

func NewCycleArrayWithSize[T any](size int) *CycleArray[T] {
	return &CycleArray[T]{
		arr:   make([]T, size),
		start: 0,
		end:   0,
		count: 0,
		size:  size,
	}
}

func (s *CycleArray[T]) Len() int {
	return s.count
}

func (s *CycleArray[T]) resize(newSize int) {
	newArr := make([]T, newSize)
	for i := 0; i < s.count; i++ {
		newArr[i] = s.arr[(s.start+i)%s.size]
	}
	s.arr = newArr
	s.start = 0
	s.end = s.count
	s.size = newSize
}

func (s *CycleArray[T]) AddFirst(v T) {
	if s.IsFull() {
		s.resize(s.size * 1)
	}
	s.start = (s.start - 1 + s.size) % s.size
	s.arr[s.start] = v
	s.count++
}

func (s *CycleArray[T]) RemoveFirst() error {
	if s.isEmpty() {
		return errors.New("the array is empty")
	}
	s.arr[s.start] = *new(T)
	s.start = (s.start + 1) % s.size
	s.count--
	if s.count > 0 && s.count == s.size/4 {
		s.resize(s.size / 2)
	}
	return nil
}

func (s *CycleArray[T]) AddLast(v T) {
	if s.IsFull() {
		s.resize(s.size * 2)
	}
	s.arr[s.end] = v
	s.end = (s.end + 1) % s.size
	s.count++
}

func (s *CycleArray[T]) RemoveLast() error {
	if s.isEmpty() {
		return errors.New("the array is empty")
	}
	s.end = (s.end - 1 + s.size) % s.size
	s.arr[s.end] = *new(T)
	s.count--
	if s.count > 0 && s.count == s.size/4 {
		s.resize(s.size / 2)
	}
	return nil
}

func (s *CycleArray[T]) getFirst() (T, error) {
	if s.isEmpty() {
		return *new(T), errors.New("the array is empty")
	}
	return s.arr[s.start], nil
}

func (s *CycleArray[T]) getLast() (T, error) {
	if s.isEmpty() {
		return *new(T), errors.New("the array is empty")
	}
	return s.arr[(s.end-1+s.size)%s.size], nil
}

func (s *CycleArray[T]) IsFull() bool {
	return s.count == s.size
}

func (s *CycleArray[T]) isEmpty() bool {
	return s.size == 0
}

func (s *CycleArray[T]) Print() {
	if s.isEmpty() {
		fmt.Println("[]")
	}
	for i := 0; i < s.count; i++ {
		fmt.Printf("%d\t", s.arr[(s.start+i)%s.size])
	}
}
