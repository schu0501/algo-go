package queue

type Queue[T any] struct {
	arr []T
}

func (s *Queue[T]) Len() int {
	return len(s.arr)
}

// Push 入队列
func (s *Queue[T]) Push(v T) {
	s.arr = append(s.arr, v)
}

// Pop 出队列
func (s *Queue[T]) Pop() T {
	var zero T
	if len(s.arr) == 0 {
		return zero
	}

	e := s.arr[0]
	s.arr[0] = zero
	s.arr = s.arr[1:]
	return e
}

// Peek 查看队列第一个元素
func (s *Queue[T]) Peek() T {
	if len(s.arr) == 0 {
		var zero T
		return zero
	}

	return s.arr[0]
}
