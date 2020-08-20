package problems

import (
  "regexp"
  "strings"
)

/**
125. Valid Palindrome
Easy

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:
Input: "A man, a plan, a canal: Panama"
Output: true

Example 2:
Input: "race a car"
Output: false

Constraints:
    s consists only of printable ASCII characters.

*/
// 题目简单，熟悉一下 Go 正则
func IsPalindrome(s string) bool {
  result := true
  regex, _ := regexp.Compile("[^a-zA-Z0-9]+")
  newS := regex.ReplaceAllString(s,"")
  newS = strings.ToLower(newS)
  for i:=0; i<len(newS)/2; i++ {
    if newS[i] != newS[len(newS)-1-i]{
      result = false
    }
  }
  return result
}

