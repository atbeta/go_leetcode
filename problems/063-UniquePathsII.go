package problems
/**
63. Unique Paths II
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
Now consider if some obstacles are added to the grids. How many unique paths would there be?

An obstacle and empty space is marked as 1 and 0 respectively in the grid.
Note: m and n will be at most 100.
Example 1:
Input:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
Output: 2
Explanation:
There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
 */

// 与62类似，如果有障碍物，则路径直接为0
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
  m := len(obstacleGrid)
  n := len(obstacleGrid[0])
  result := make([][]int, m)
  for i := 0; i < m; i++ {
    result[i] = make([]int, n)
  }
  // 第一行和第一列初始值计算
  if obstacleGrid[0][0] == 0 {
    result[0][0] = 1
  } else {
    result[0][0] = 0
  }
  // 如果前一个位置为0或者当前grid有障碍物则为0，否则为1
  for i := 1; i < m; i++ {
    if obstacleGrid[i][0] == 1 || result[i-1][0] == 0 {
      result[i][0] = 0
    } else {
      result[i][0] = 1
    }
  }
  for j := 1; j < n; j++ {
    if obstacleGrid[0][j] == 1 || result[0][j-1] == 0 {
      result[0][j] = 0
    } else {
      result[0][j] = 1
    }
  }
  // 求其他位置值，如果位置上有障碍物则为0，否则为左侧和上方路径数之和
  for i := 1; i < m; i++ {
    for j := 1; j < n; j++ {
      if obstacleGrid[i][j] == 1 {
        result[i][j] = 0
      } else {
        result[i][j] = result[i-1][j] + result[i][j-1]
      }
    }
  }
  return result[m-1][n-1]
}
