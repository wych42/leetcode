/*
 * @lc app=leetcode id=889 lang=golang
 *
 * [889] Construct Binary Tree from Preorder and Postorder Traversal
 *
 * https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (67.56%)
 * Likes:    1177
 * Dislikes: 62
 * Total Accepted:    49.7K
 * Total Submissions: 73.3K
 * Testcase Example:  '[1,2,4,5,3,6,7]\n[4,5,2,6,7,3,1]'
 *
 * Return any binary tree that matches the given preorder and postorder
 * traversals.
 *
 * Values in the traversals pre and post are distinct positive integers.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
 * Output: [1,2,3,4,5,6,7]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= pre.length == post.length <= 30
 * pre[] and post[] are both permutations of 1, 2, ..., pre.length.
 * It is guaranteed an answer exists. If there exists multiple answers, you can
 * return any of them.
 *
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
func constructFromPrePost(pre []int, post []int) *TreeNode {
	/*
		pre root-left-right
		post left-right-root
		pre = [1,2,4,5,3,6,7]
		post = [4,5,2,6,7,3,1]

		root=1;
		pre = [2,4,5,3,6,7]
		post = [4,5,2,6,7,3]

		此时 pre[0] 是 root.Left, 在 post 中 idx=2
		也就是说 [245] 和 [452] 是左子树
	*/
	return build(pre, post)
}
func build(pre []int, post []int) *TreeNode {
	if len(pre) == 0 || len(post) == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}
	if len(pre) == 1 && len(post) == 1 {
		return root
	}
	leftCount := 0
	for i, n := range post {
		if n == pre[1] {
			leftCount = i + 1
			break
		}
	}
	//    fmt.Println(pre, post, leftCount)

	root.Left = build(pre[1:leftCount+1], post[0:leftCount])
	root.Right = build(pre[leftCount+1:], post[leftCount:len(post)-1])
	return root
}

// @lc code=end

