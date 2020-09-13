package problems

/**
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
Note: You can only move either down or right at any point in time.

Example:
Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
 */

// 与p62类似，动态规划
func MinPathSum(grid [][]int) int {
  // 初始化
  m := len(grid)
  n := len(grid[0])
  result := make([][]int, m)
  for i, _ := range result {
    result[i] = make([]int, n)
  }
  result[0][0] = grid[0][0]
  for i := 1; i < n; i++ {
    result[0][i] = result[0][i-1] + grid[0][i]
  }
  for j := 1; j < m; j++ {
    result[j][0] = result[j-1][0] + grid[j][0]
  }
  // 更新各值，每一个位置的最小值都是上方或左方位置的最小和加当前位置值
  for i := 1; i < m; i++ {
    for j:= 1; j < n; j++ {
      min := result[i-1][j]
      if result[i][j-1] < min {
        min = result[i][j-1]
      }
      result[i][j] = min + grid[i][j]
    }
  }
  return result[m-1][n-1]
}
