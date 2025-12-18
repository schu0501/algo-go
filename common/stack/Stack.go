package stack

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) Len() int {
	return len(s.arr)
}

// Push 入栈
func (s *Stack[T]) Push(v T) {
	s.arr = append(s.arr, v)
}

// Pop 出栈
func (s *Stack[T]) Pop() T {
	if len(s.arr) == 0 {
		var zero T
		return zero
	}

	e := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return e
}

// Peek 查看栈顶元素
func (s *Stack[T]) Peek() T {
	if len(s.arr) == 0 {
		var zero T
		return zero
	}

	return s.arr[len(s.arr)-1]
}
