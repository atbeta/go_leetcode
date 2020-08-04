package problems
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import (
  . "github.com/atbeta/go_leetcode/structs"
)
func RemoveElements(head *ListNode, val int) *ListNode {
  if head == nil {
    return nil
  }
  current := head
  newHead := head
  // 先排除第一个开始就是 val
  for current != nil && current.Val == val {
    newHead = current.Next
    current = current.Next
  }
  // 如果下一个值等于 val，就跳过
  for current != nil && current.Next != nil {
    if current.Next.Val == val {
      current.Next = current.Next.Next
    } else {
      current = current.Next
    }
  }
  return newHead
}
