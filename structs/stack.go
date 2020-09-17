package structs

// 最基本的栈，不考虑多线程等
type Stack struct {
  items []Item
}

func NewStack() *Stack {
  return &Stack{items: []Item{}}
}

func (s *Stack) Push(item Item) {
  s.items = append(s.items, item)
}

func (s *Stack) Pop() Item {
  length := len(s.items)
  popItem := s.items[length - 1]
  s.items = s.items[:length - 1]
  return popItem
}
