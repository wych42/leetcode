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
	vals []*int

	segs []string
	idx  int
}

func Constructor() Codec {
	c := Codec{
		vals: make([]*int, 0),
	}
	return c
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	this.bfs(root)
	var s strings.Builder
	for _, val := range this.vals {
		if val != nil {
			s.WriteString(strconv.Itoa(*val))
		} else {
			s.WriteString("null")
		}
		s.WriteString(",")
	}
	return s.String()
}

func (this *Codec) bfs(root *TreeNode) {
	if root == nil {
		this.vals = append(this.vals, nil)
		return
	}
	this.vals = append(this.vals, &root.Val)
	this.bfs(root.Left)
	this.bfs(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	this.segs = strings.Split(data, ",")
	return this.debfs()
}

func (this *Codec) debfs() *TreeNode {
	if this.idx >= len(this.segs) {
		return nil
	}
	seg := this.segs[this.idx]
	if seg == "" || seg == "null" {
		return nil
	}
	val, _ := strconv.Atoi(seg)
	root := &TreeNode{Val: val}
	this.idx++
	root.Left = this.debfs()
	this.idx++
	root.Right = this.debfs()
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

