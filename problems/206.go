package problems

// 206. Reverse Linked List
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import (
  "fmt"
  . "github.com/atbeta/go_leetcode/structs"
)
func ReverseList(head *ListNode) *ListNode {
  // 每次将后一个变成新的head
  if head == nil || head.Next == nil {
    return head
  }
  currentHead := head
  current := head.Next
  originHead := head
  for current.Next != nil {
    temp := current.Next
    current.Next = currentHead
    originHead.Next = temp
    currentHead = current
    current = temp
    fmt.Println("-------------------")
    currentHead.Print()
  }
  originHead.Next = nil
  current.Next = currentHead
  currentHead = current
  return currentHead
}
