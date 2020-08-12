package problems

/**
342. Power of Four
Easy
Given an integer (signed 32 bits), write a function to check whether it is a power of 4.

Example 1:
Input: 16
Output: true

Example 2:
Input: 5
Output: false

Follow up: Could you solve it without loops/recursion?
 */
func IsPowerOfFour(num int) bool {
  result := false
  current := 1
  for i:=current;i<=num;i=i*4 {
    if i == num {
      result = true
    }
  }
  return result
}
