package problems

/**

118. Pascal's Triangle
Easy
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:
Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

*/
func GeneratePascalTriangle(numRows int) [][]int {
  var result [][]int
  for i:=0;i<numRows;i++ {
    if i == 0 {
      result = append(result, []int{1})
    }
    if i == 1 {
      result = append(result, []int{1,1})
    }
    if i > 1 {
      nthRow := make([]int, i + 1, i + 1)
      nthRow[0] = 1
      nthRow[i] = 1
      for j:=1;j<i;j++ {
        nthRow[j] = result[i-1][j-1] + result[i-1][j]
      }
      result = append(result, nthRow)
    }
  }
  return result
}
