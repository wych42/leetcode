/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 *
 * https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/description/
 *
 * algorithms
 * Medium (42.01%)
 * Likes:    2287
 * Dislikes: 196
 * Total Accepted:    320.2K
 * Total Submissions: 760.3K
 * Testcase Example:  '[1,2,3,4,5,null,7]'
 *
 * Given a binary tree
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
 * Input: root = [1,2,3,4,5,null,7]
 * Output: [1,#,2,3,#,4,5,7,#]
 * Explanation: Given the above binary tree (Figure A), your function should
 * populate each next pointer to point to its next right node, just like in
 * Figure B. The serialized output is in level order as connected by the next
 * pointers, with '#' signifying the end of each level.
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the given tree is less than 6000.
 * -100 <= node.val <= 100
 *
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

func connect(root *Node) *Node {
	m := make(map[int]*Node)
	traversal(root, 0, m)
	return root
}

// 用map暂存同一层中，在最右边且没有被指向的节点
// 一个反向的前序遍历: 根-右-左
func traversal(root *Node, level int, m map[int]*Node) {
	if root == nil { // 走到叶子节点了
		return
	}
	// 如果同级，有存在节点，说明是当前节点的右节点, 指过去，然后当前节点变成同级最右并且没有链接的节点
	if p, ok := m[level]; ok {
		root.Next = p
	}
	// 先访问右边的节点
	// 用 map 把右边的节点暂存起来
	m[level] = root
	traversal(root.Right, level+1, m)
	traversal(root.Left, level+1, m)
}

// @lc code=end

