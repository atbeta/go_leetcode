package problems

/**
136. Single Number
Easy
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:
Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:
Input: [2,2,1]
Output: 1

Example 2:
Input: [4,1,2,1,2]
Output: 4
*/

// 之前提交过，这是位运算经典应用，相同数字按位异或为0,任意数字与0按位异或保持不变
func SingleNumber(nums []int) int {
  result := 0
  for _, num := range nums {
    result = result ^ num
  }
  return result
}
