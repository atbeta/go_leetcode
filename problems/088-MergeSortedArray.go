package problems

/**
88. Merge Sorted Array
Easy
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:
    The number of elements initialized in nums1 and nums2 are m and n respectively.
    You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.

Example:
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
 */

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int)  {
  p1 := m - 1
  p2 := n - 1
  p := m + n -1

  // 每次都把两个切片中最大的值放在最后
  for p1 >= 0 && p2 >= 0 {
    if nums1[p1] < nums2[p2] {
      nums1[p] = nums2[p2]
      p2--
    } else {
      nums1[p] = nums1[p1]
      p1--
    }
    p--
  }
  // 最后把 nums2 中剩余值覆盖 nums1 对应值
  copy(nums1[:p2+1], nums2[:p2+1])
}
