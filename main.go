package main

import (
  "github.com/atbeta/go_leetcode/problems"
  . "github.com/atbeta/go_leetcode/structs"
)

func main() {
  node4 := ListNode{Val: 5, Next: nil}
  node3 := ListNode{Val: 4, Next: &node4}
  node2 := ListNode{Val: 3, Next: &node3}
  node1 := ListNode{Val: 2, Next: &node2}
  head := ListNode{Val: 1, Next: &node1}
  problems.SwapPairs(&head)
  //merged := problems.MergeTwoLists(&head, &node1)
  //merged.Print()
}
