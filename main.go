package main

import (
  "fmt"
  "github.com/atbeta/go_leetcode/problems"
)

func main() {
  //s := []int{1,2,2,1}
  //head := NewListNode(s)
  //problems.SwapPairs(head)
  //reversed := problems.ReverseList(head)
  //fmt.Println("------------------")
  //reversed.Print()
  // p203
  //removed := problems.RemoveElements(head, 2)
  //removed.Print()
  //steps := problems.ClimbStairs(45)
  //fmt.Println(steps)
  //result := problems.MySqrt(0)
  //fmt.Println(result)
  //result := problems.IsPerfectSquare(17)
  //fmt.Println(result)
  //result := problems.JudgeSquareSum(16)
  //fmt.Println(result)
  //result := problems.IsPowerOfTwo(26)
  //fmt.Println(result)
  //result :=problems.GeneratePascalTriangle(10)
  //fmt.Println(result)
  result := problems.GetPascalRow(5)
  fmt.Println(result)
}
