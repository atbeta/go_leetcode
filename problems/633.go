package problems

/**
633. Sum of Square Numbers
Easy
Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:
Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5

Example 2:
Input: 3
Output: False
*/
// 会超时
//func JudgeSquareSum(c int) bool {
// isPerfectSquare := func(num int) bool {
//   if num == 0 || num == 1 {
//   return true
// }
//   start,end := 1, num
//   result := false
//   for start < end - 1 { // 两个整数之间不可能有整数
//   mid := (start + end)/2
//   if mid * mid == num {
//   result = true
//   break
// } else if mid * mid > num {
//   end = mid
// } else {
//   start = mid
// }
// }
//   return result
// }
// result := false
// for i:=0;i<=c/2;i++ {
//   if isPerfectSquare(i) {
//     result = isPerfectSquare(c - i)
//     if result == true {
//       //fmt.Println(i, c-i)
//       break
//     }
//   }
// }
// return result
//}

// 无论怎么优化都会超时
//func JudgeSquareSum(c int) bool {
// // 生成平方不大于c的所有完全平方数的数组
// var perfectSquareList []int
// for i:=0;i*i<=c;i++ {
//   perfectSquareList = append(perfectSquareList, i*i)
// }
// // 判定是否在数组中
// listLen := len(perfectSquareList)
// isContained := func(num int, startIndex int) bool {
//   if perfectSquareList[startIndex] > c/2 {
//     return false
//   }
//   if num + perfectSquareList[startIndex] > c{
//     return false
//   }
//   result := false
//   for i:=listLen-1;i>=startIndex;i-- {
//     if perfectSquareList[i] == num {
//       result = true
//       break
//     }
//   }
//   return result
// }
// result := false
// for i:=0;i<listLen;i++ {
//   if isContained(c - perfectSquareList[i], i) {
//     result = true
//     break
//   }
// }
// return result
//}

// 使用 map
func JudgeSquareSum(c int) bool {
  var isSquare map[int]bool
  isSquare = make(map[int]bool)
  // 将平方数对应的map赋值为true
  for i:=0;i*i<=c;i++ {
    isSquare[i*i] = true
  }
  result := false
  for k, _:= range isSquare {
    if isSquare[c-k] {
      result = true
      break
    }
  }
  return result
}
