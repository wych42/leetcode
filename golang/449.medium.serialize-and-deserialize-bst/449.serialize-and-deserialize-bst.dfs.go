/*
 * @lc app=leetcode id=449 lang=golang
 *
 * [449] Serialize and Deserialize BST
 *
 * https://leetcode.com/problems/serialize-and-deserialize-bst/description/
 *
 * algorithms
 * Medium (54.00%)
 * Likes:    1857
 * Dislikes: 94
 * Total Accepted:    146.9K
 * Total Submissions: 271.4K
 * Testcase Example:  '[2,1,3]'
 *
 * Serialization is converting a data structure or object into a sequence of
 * bits so that it can be stored in a file or memory buffer, or transmitted
 * across a network connection link to be reconstructed later in the same or
 * another computer environment.
 *
 * Design an algorithm to serialize and deserialize a binary search tree. There
 * is no restriction on how your serialization/deserialization algorithm should
 * work. You need to ensure that a binary search tree can be serialized to a
 * string, and this string can be deserialized to the original tree structure.
 *
 * The encoded string should be as compact as possible.
 *
 *
 * Example 1:
 * Input: root = [2,1,3]
 * Output: [2,1,3]
 * Example 2:
 * Input: root = []
 * Output: []
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 10^4].
 * 0 <= Node.val <= 10^4
 * The input tree is guaranteed to be a binary search tree.
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "strings"
import "strconv"

type Codec struct {
	sb strings.Builder

	segs []string
	idx  int
}

func Constructor() Codec {
	c := Codec{}
	return c
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	this.sb = strings.Builder{}
	this.preorder(root)
	str := this.sb.String()
	return str[:len(str)-1]
}

func (this *Codec) preorder(root *TreeNode) {
	if root == nil {
		return
	}
	this.sb.WriteString(strconv.Itoa(root.Val))
	this.sb.WriteString(",")
	this.preorder(root.Left)
	this.preorder(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	this.segs = strings.Split(data, ",")
	return this.depreorder(-1<<31, 1<<31)
}

/*
   第一个元素是 root
   往后每一个元素如果比 root 小，就是左节点；比root大，就是右节点
*/
func (this *Codec) depreorder(min, max int) *TreeNode {
	if this.idx >= len(this.segs) {
		return nil
	}
	val, _ := strconv.Atoi(this.segs[this.idx])
	if val < min || val > max {
		return nil
	}
	this.idx++
	root := &TreeNode{Val: val}
	root.Left = this.depreorder(min, val)  // 左子树的范围是 min, val(root)
	root.Right = this.depreorder(val, max) // 右子树的范围是 val(root), max
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// @lc code=end

