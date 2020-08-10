package problems

/**
You are climbing a stair case. It takes n steps to reach to the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
Example 1:
Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 */
//func ClimbStairs(n int) int {
//  if n == 1 {
//    return n
//  }
//  if n == 2 {
//    return 2
//  }
//  return ClimbStairs(n - 1) + ClimbStairs(n - 2)
//}

func ClimbStairs(n int) int {
  step := make([]int, 50)
  step[1] = 1
  step[2] = 2
  for i:=3;i<=n;i++ {
    step[i] = step[i-1] + step[i-2]
  }
  return step[n]
}
