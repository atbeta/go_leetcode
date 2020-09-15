package problems

import (
  . "github.com/atbeta/go_leetcode/structs"
)

/**
Given a binary tree, return the inorder traversal of its nodes' values.

Example:
Input: [1,null,2,3]
   1
    \
     2
    /
   3
Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?

 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 中序遍历二叉树，递归和非递归
func InorderTraversal(root *TreeNode) []int {
  var result []int
  if root == nil {
    return result
  }
  result = append(result, InorderTraversal(root.Left)...)
  result = append(result, root.Val)
  result = append(result, InorderTraversal(root.Right)...)
  return result
}
