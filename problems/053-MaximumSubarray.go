package problems

/**
53. Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

Example 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Example 2:
Input: nums = [1]
Output: 1
Example 3:
Input: nums = [0]
Output: 0
Example 4:
Input: nums = [-1]
Output: -1
Example 5:
Input: nums = [-2147483647]
Output: -2147483647

Constraints:
1 <= nums.length <= 2 * 104
-231 <= nums[i] <= 231 - 1
 */

// 动态规划与分治，穷举会超出时间限制
// 动态规划，result[i][j]表示从i到j的子数组的和, j>i
// result[i][j] = result[i][j-1]+nums[j]
//func MaxSubArray(nums []int) int {
//  length := len(nums)
//  result := make([][]int, length)
//  for i := 0; i < length; i++ {
//    result[i] = make([]int, length)
//  }
//  // 初始值，如果 j=i，那么 result[i][j]=nums[i]
//  for i := 0; i < length; i++ {
//    result[i][i] = nums[i]
//  }
//  // 循环计算
//  for i := 0; i < length; i++ {
//    for j := i+1; j < length; j++ {
//      result[i][j] = result[i][j-1] + nums[j]
//    }
//  }
//  // 找到最大值
//  maxSum := result[0][0]
//  for i := 0; i < length; i++ {
//    for j := i; j < length; j++ {
//      if result[i][j] > maxSum {
//        maxSum = result[i][j]
//      }
//    }
//  }
//  return maxSum
//}

// 未经优化的动态规划，会超出空间限制，不再记录所有值，只保留以i位置开始的和，并记录当前最大值
func MaxSubArray(nums []int) int {
  length := len(nums)
  result := make([]int, length)
  maxSum := nums[0]
  for i := 0; i < length; i++ {
    for j := i; j < length; j++ {
      // 以i开头，以j结尾的和
      if j == i {
        result[j] = nums[j]
      } else {
        result[j] = result[j-1] + nums[j]
      }
      // 找到该行最大的值,与 maxSum 比较
      if result[j] > maxSum {
        maxSum = result[j]
      }
    }
  }
  return maxSum
}
