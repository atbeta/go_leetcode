package structs

import "fmt"

type ListNode struct {
  Val int
  Next *ListNode
}

func (m *ListNode) Print() {
  current := m
  for current != nil {
    fmt.Println(current.Val)
    current = current.Next
  }
}

func (m *ListNode) Len() int {
  count := 0
  current := m
  for current != nil {
    count ++
    current = current.Next
  }
  return count
}
