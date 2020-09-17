package problems

import
(
  . "github.com/atbeta/go_leetcode/structs"
)

/**
98. Validate Binary Search Tree
Given a binary tree, determine if it is a valid binary search tree (BST).
Assume a BST is defined as follows:
    The left subtree of a node contains only nodes with keys less than the node's key.
    The right subtree of a node contains only nodes with keys greater than the node's key.
    Both the left and right subtrees must also be binary search trees.

Example 1:
    2
   / \
  1   3

Input: [2,1,3]
Output: true

Example 2:
    5
   / \
  1   4
     / \
    3   6

Input: [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历，检查是否是升序排列，如果是则有效
func IsValidBST(root *TreeNode) bool {
  return isAscending(inorderTraverse(root))
}

func inorderTraverse(root *TreeNode) []int {
  var result []int
  if root == nil {
    return result
  }
  result = append(result, inorderTraverse(root.Left)...)
  result = append(result, root.Val)
  result = append(result, inorderTraverse(root.Right)...)
  return result
}

func isAscending(nums []int) bool {
  result := true
  length := len(nums)
  for i := 0; i < length - 1; i++ {
    if nums[i+1] <= nums[i] {
      result = false
    }
  }
  return result
}
