package adt

type Stack struct {
	size  int
	items [10]string
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(value string) {
	s.items[s.size] = value
	s.size++
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Pop() string {
	s.size--
	return s.items[s.size]
}

func NewStack() *Stack {
	return &Stack{size: 0}
}
