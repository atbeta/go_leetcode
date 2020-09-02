package problems

import (
  . "github.com/atbeta/go_leetcode/structs"
)

/**
83. Remove Duplicates from Sorted List
Easy

Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2

Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 使用 Map 存储，有值就删除
func deleteDuplicates(head *ListNode) *ListNode {
  keyMap := make(map[int]bool)
  if head == nil {
    return nil
  }
  keyMap[head.Val] = true
  current := head
  for current.Next != nil {
    if keyMap[current.Next.Val] {
      // 如果已经有此值就删除节点
      current.Next = current.Next.Next
    } else {
      //  如果没有就设置该key为true,移动到下一个位置
      keyMap[current.Next.Val] = true
      current = current.Next
    }
  }
  return head
}
