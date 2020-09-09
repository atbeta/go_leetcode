package problems

/**
137. Single Number II
Medium
Given a non-empty array of integers, every element appears three times except for one, which appears exactly once. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
Input: [2,2,3,2]
Output: 3

Example 2:
Input: [0,1,0,1,0,1,99]
Output: 99
*/

// 136 的变种版本，异或不再有效
// 不使用位运算，比较容易想到的方法是使用 Map，逐个数字统计
// 位运算方法，使用一个长度为32的数组存在每个数字的二进制位和，不进位，各位对3取余，剩下的就是单独的值，再转换回10进制
func SingleNumberII(nums []int) int {
  sums := make([]int, 32)
  for _, num := range nums {
    // 各数求各二进制位
    temp := num
    index := 0
    for temp > 0 {
      sums[index] += temp % 2
      index ++
      temp = temp /2
    }
  }
  // 各二进制位之和对3取余之后转成十进制
  result := 0
  for index, num := range sums {
    if num % 3 == 1 {
      result += 1 >> index
    }
  }
  return result
}
