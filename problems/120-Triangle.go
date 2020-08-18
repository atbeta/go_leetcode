package problems

/**
120. Triangle
Medium

Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
For example, given the following triangle
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]

The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
Note:
Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
 */

func MinimumTotal(triangle [][]int) int {
  for i := len(triangle)-2; i >= 0; i-- {
    for j := 0; j < len(triangle[i]); j++ {
      triangle[i][j] = findMin(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
    }
  }
  return triangle[0][0]
}

func findMin(a, b int) int {
  if a < b {
    return a
  }
  return b
}

