package problems

import "fmt"

/**

367. Valid Perfect Square
Easy
Given a positive integer num, write a function which returns True if num is a perfect square else False.
Follow up: Do not use any built-in library function such as sqrt.

Example 1:
Input: num = 16
Output: true

Example 2:
Input: num = 14
Output: false

Constraints:
    1 <= num <= 2^31 - 1
*/

//func IsPerfectSquare(num int) bool {
//  result := false
//  for i:=1;i<=num;i++ {
//    if i*i == num {
//      result = true
//    }
//  }
//  return result
//}

// 使用二分搜索
func IsPerfectSquare(num int) bool {
  if num == 1 {
    return true
  }
  start,end := 1, num
  result := false
  for start < end - 1 { // 两个整数之间不可能有整数
    mid := (start + end)/2
    fmt.Println(mid)
    if mid * mid == num {
      result = true
      break
    } else if mid * mid > num {
      end = mid
    } else {
      start = mid
    }
  }
  return result
}
