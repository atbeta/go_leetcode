package problems

/**
Given an integer, write a function to determine if it is a power of three.

Example 1:
Input: 27
Output: true

Example 2:
Input: 0
Output: false

Example 3:
Input: 9
Output: true

Example 4:
Input: 45
Output: false

Follow up:
Could you do it without using any loop / recursion?
 */

// 使用循环
func IsPowerOfThree(n int) bool {
  result := false
  current := 1
  for i:=current;i<=n;i=i*3 {
    if i == n {
      result = true
    }
  }
  return result
}

// 不使用循环的话，一种是找到范围内最大的3的指数然后检查整除，一种是利用以3为底的对数是否为整数

