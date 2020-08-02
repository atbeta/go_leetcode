package problems

import (
  . "github.com/atbeta/go_leetcode/structs"
)

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
  // 计算长度
  lenOfList := head.Len()
  toDeleteIndex := lenOfList - n
  index := 0
  current := head
  for index != toDeleteIndex - 1 {
    index ++
    current = current.Next
  }
  current.Next = current.Next.Next
  return head
}
