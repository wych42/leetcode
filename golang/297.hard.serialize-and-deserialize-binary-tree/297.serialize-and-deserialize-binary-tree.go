/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
 *
 * https://leetcode.com/problems/serialize-and-deserialize-binary-tree/description/
 *
 * algorithms
 * Hard (49.72%)
 * Likes:    4043
 * Dislikes: 188
 * Total Accepted:    420.2K
 * Total Submissions: 841.9K
 * Testcase Example:  '[1,2,3,null,null,4,5]'
 *
 * Serialization is the process of converting a data structure or object into a
 * sequence of bits so that it can be stored in a file or memory buffer, or
 * transmitted across a network connection link to be reconstructed later in
 * the same or another computer environment.
 *
 * Design an algorithm to serialize and deserialize a binary tree. There is no
 * restriction on how your serialization/deserialization algorithm should work.
 * You just need to ensure that a binary tree can be serialized to a string and
 * this string can be deserialized to the original tree structure.
 *
 * Clarification: The input/output format is the same as how LeetCode
 * serializes a binary tree. You do not necessarily need to follow this format,
 * so please be creative and come up with different approaches yourself.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,3,null,null,4,5]
 * Output: [1,2,3,null,null,4,5]
 *
 *
 * Example 2:
 *
 *
 * Input: root = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1]
 * Output: [1]
 *
 *
 * Example 4:
 *
 *
 * Input: root = [1,2]
 * Output: [1,2]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 10^4].
 * -1000 <= Node.val <= 1000
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
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

