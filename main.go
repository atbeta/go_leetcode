package main

import (
  "fmt"
  "github.com/atbeta/go_leetcode/problems"
  . "github.com/atbeta/go_leetcode/structs"
)

func main() {
  s := []int{1,2,3,100,21,5}
  head := NewListNode(s)
  //problems.SwapPairs(head)
  reversed := problems.ReverseList(head)
  fmt.Println("------------------")
  reversed.Print()
}
