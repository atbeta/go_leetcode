package problems

/**
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
Note that the row index starts from 0.

In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:
Input: 3
Output: [1,3,3,1]

Follow up:
Could you optimize your algorithm to use only O(k) extra space?
 */

// O(k) 额外空间，只开辟一个slice
func GetPascalRow(rowIndex int) []int {
  row := make([]int, rowIndex + 1, rowIndex + 1)
  for i:=0;i<=rowIndex;i++ {
    row[0] = 1
    row[i] = 1
    if i > 1 {
      // 改变从第2个值到倒数第2个值
      // 当前位置值 = 当前位置原值 + 最后一个的值
      // 最后一个的值替换为当前位置原值 = 当前位置值 - 最后一个的值
      for j:=1;j<i;j++ {
        row[j] = row[j] + row[i]
        row[i] = row[j] - row[i]
        //fmt.Println(row)
      }
    }
  }
  return row
}
