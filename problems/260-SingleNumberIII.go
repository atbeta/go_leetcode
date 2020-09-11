package problems

/**
260. Single Number III
Medium
Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once. You can return the answer in any order.
Follow up: Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?

Example 1:
Input: nums = [1,2,1,3,2,5]
Output: [3,5]
Explanation:  [5, 3] is also a valid answer.

Example 2:
Input: nums = [-1,0]
Output: [-1,0]

Example 3:
Input: nums = [0,1]
Output: [1,0]

Constraints:
    1 <= nums.length <= 30000
     Each integer in nums will appear twice, only two integers will appear once.
*/

// 思路：只有一个不同的可以通过异或运算快速解决，本问题可以通过将数字分为两类，每类中仅有一个单独的数字
// 先将所有数字逐个异或一次，得到的结果就是这两个数字的异或结果
// 根据异或的定义，结果二进制位中的出现的第一个1表示在这一位上这两个数字不同，因此可以以此为标准分组数字
func SingleNumberIII(nums []int) []int {
  // 逐个异或
  result1 := 0
  for _, num := range nums {
    result1 = result1 ^ num
  }
  // 找到第一个1,得到一个需要右称的值
  right := 0
  for result1 % 2 == 0 {
    result1 = result1 >> 1
    right ++
  }
  // 分类数组并逐个异或
  // 将数字右移right位后对2取余分组
  var nums1 []int
  var nums2 []int
  for _, num := range nums {
    if (num >> right) % 2 == 0 {
      nums1 = append(nums1, num)
    } else {
      nums2 = append(nums2, num)
    }
  }
  // 分别循环nums1和nums2，每个得到一个单独出现的值
  num1 := 0
  num2 := 0
  for _,num := range nums1 {
    num1 = num1 ^ num
  }
  for _,num := range nums2 {
    num2 = num2 ^ num
  }
  return []int{num1, num2}
}
