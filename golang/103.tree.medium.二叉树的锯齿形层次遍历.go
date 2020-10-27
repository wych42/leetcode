/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/description/
 *
 * algorithms
 * Medium (55.11%)
 * Likes:    285
 * Dislikes: 0
 * Total Accepted:    76.6K
 * Total Submissions: 138.9K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值的锯齿形层次遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回锯齿形层次遍历如下：
 *
 * [
 * ⁠ [3],
 * ⁠ [20,9],
 * ⁠ [15,7]
 * ]
 *
 *
 */

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	fmt.Println(
		zigzagLevelOrder(&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		))
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	toggleOrder := false
	for len(queue) > 0 {
		layer := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			var node *TreeNode
			node = queue[length-1-i]
			if node == nil {
				continue
			}
			layer = append(layer, node.Val)
			if toggleOrder {
				queue = append(queue, node.Right, node.Left)
			} else {
				queue = append(queue, node.Left, node.Right)
			}
		}
		queue = queue[length:]
		if len(layer) > 0 {
			result = append(result, layer)
		}
		toggleOrder = !toggleOrder
	}
	return result
}

// @lc code=end
