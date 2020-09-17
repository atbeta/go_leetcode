package problems

import (
  . "github.com/atbeta/go_leetcode/structs"
)

/**
112. Path Sum
Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.
Note: A leaf is a node with no children.

Example:
Given the below binary tree and sum = 22,
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1

return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归容易解决,逐层减下去，如果到孩子节点能减到0，则为true
func HasPathSum(root *TreeNode, sum int) bool {
  if root == nil {
    return false
  }
  if root.Left == nil && root.Right == nil {
    return root.Val == sum
  }
  return HasPathSum(root.Left, sum - root.Val) || HasPathSum(root.Right, sum - root.Val)
}

