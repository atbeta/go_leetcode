package problems

/**
121. Best Time to Buy and Sell Stock
Easy

Say you have an array for which the ith element is the price of a given stock on day i.
If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
Note that you cannot sell a stock before you buy one.

Example 1:
Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.

Example 2:
Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.

*/

// 暴力方法： 将数组拆成两部分，从前半部分找最小的，后半部分找最大的，计算差值
func MaxProfit(prices []int) int {
  length := len(prices)
  profit := 0
  for splitIndex:=1;splitIndex<length;splitIndex++ {
    left := prices[:splitIndex]
    right := prices[splitIndex:]
    currentProfit := max(right) - min(left)
    if currentProfit > profit {
      profit = currentProfit
    }
  }
  return profit
}

func max(l []int) (max int) {
  max = l[0]
  for _, v := range l {
    if v > max {
      max = v
    }
  }
  return
}

func min(l []int) (min int) {
  min = l[0]
  for _, v := range l {
    if v < min {
      min = v
    }
  }
  return
}
