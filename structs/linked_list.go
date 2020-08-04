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

func NewListNode(s []int) *ListNode {
  head := new(ListNode)
  length := len(s)
  if length == 0 {
    return nil
  }
  current := head
  for index, value := range s {
    if index < length - 1 {
      current.Val = value
      current.Next = new(ListNode)
      current = current.Next
    } else {
      current.Val = value
      current.Next = nil
    }
  }
  return head
}
