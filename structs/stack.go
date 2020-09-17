package structs

// 最基本的栈，不考虑多线程等
type Stack struct {
  items []int
}

func (s *Stack) New() *Stack {
  s.items = []int{}
  return s
}

func (s *Stack) Push(item int) {
  s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
  length := len(s.items)
  popItem := s.items[length - 1]
  s.items = s.items[:length - 1]
  return popItem
}
