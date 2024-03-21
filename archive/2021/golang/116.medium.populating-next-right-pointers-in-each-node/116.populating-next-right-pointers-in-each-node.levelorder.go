/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
 *
 * https://leetcode.com/problems/populating-next-right-pointers-in-each-node/description/
 *
 * algorithms
 * Medium (48.98%)
 * Likes:    3120
 * Dislikes: 161
 * Total Accepted:    463.9K
 * Total Submissions: 944.3K
 * Testcase Example:  '[1,2,3,4,5,6,7]'
 *
 * You are given a perfect binary tree where all leaves are on the same level,
 * and every parent has two children. The binary tree has the following
 * definition:
 *
 *
 * struct Node {
 * ⁠ int val;
 * ⁠ Node *left;
 * ⁠ Node *right;
 * ⁠ Node *next;
 * }
 *
 *
 * Populate each next pointer to point to its next right node. If there is no
 * next right node, the next pointer should be set to NULL.
 *
 * Initially, all next pointers are set to NULL.
 *
 *
 *
 * Follow up:
 *
 *
 * You may only use constant extra space.
 * Recursive approach is fine, you may assume implicit stack space does not
 * count as extra space for this problem.
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: root = [1,2,3,4,5,6,7]
 * Output: [1,#,2,3,#,4,5,6,7,#]
 * Explanation: Given the above perfect binary tree (Figure A), your function
 * should populate each next pointer to point to its next right node, just like
 * in Figure B. The serialized output is in level order as connected by the
 * next pointers, with '#' signifying the end of each level.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the given tree is less than 4096.
 * -1000 <= node.val <= 1000
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// 就是一个前序遍历
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)

	layer := make([]*Node, 0)
	idx := 0
	size := len(queue)
	for idx < len(queue) {
		if idx >= size {
			size = len(queue)
			for i := 0; i < len(layer); i++ {
				if layer[i] == nil {
					continue
				}
				if i+1 == len(layer) {
					layer[i].Next = nil
					continue
				}
				if i+1 < len(layer) && layer[i+1] != nil {
					layer[i].Next = layer[i+1]
				}
			}
			layer = make([]*Node, 0) // reset layer
		}
		node := queue[idx]
		idx++
		if node == nil {
			continue
		}

		queue = append(queue, node.Left, node.Right)
		layer = append(layer, node)
	}
	return root
}

// @lc code=end

