/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
 *
 * https://leetcode.com/problems/path-sum-iii/description/
 *
 * algorithms
 * Medium (48.10%)
 * Likes:    4873
 * Dislikes: 319
 * Total Accepted:    258.8K
 * Total Submissions: 537.2K
 * Testcase Example:  '[10,5,-3,3,2,null,11,3,-2,null,1]\n8'
 *
 * You are given a binary tree in which each node contains an integer value.
 *
 * Find the number of paths that sum to a given value.
 *
 * The path does not need to start or end at the root or a leaf, but it must go
 * downwards
 * (traveling only from parent nodes to child nodes).
 *
 * The tree has no more than 1,000 nodes and the values are in the range
 * -1,000,000 to 1,000,000.
 *
 * Example:
 *
 * root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
 *
 * ⁠     10
 * ⁠    /  \
 * ⁠   5   -3
 * ⁠  / \    \
 * ⁠ 3   2   11
 * ⁠/ \   \
 * 3  -2   1
 *
 * Return 3. The paths that sum to 8 are:
 *
 * 1.  5 -> 3
 * 2.  5 -> 2 -> 1
 * 3. -3 -> 11
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
/*
 暴力解法：
 计算从root出发，满足条件的 path 数量
 再把每一个节点当作 root，挨个计算

优化:
对树做一个前序遍历:
- 遍历拿到的值添加到 path 末尾
  - 从 path 尾部开始往前求和，碰到结果==sum就把 count+1
  - 不从前面开始加，是因为前面的 path 已经被计算过了
- 每次到达叶子节点后，都把他从 path 里闪电，也就是在一个递归调用栈结束时删除 path 里的最后一个元素去掉

以下面的输入为例：

[5,4,8,11,null,13,4,7,2,null,null,5,1]
22

[5] 0
[5 4] 0
[5 4 11] 0
[5 4 11 7] 0 <-- 到达叶子节点, 去掉，计算另一个方向
[5 4 11 2] 1 <-- 叶子节点
[5 8] 2 <-- 计算右子树
[5 8 13] 2 <-- 叶子节点
[5 8 4] 2  <-- 叶子节点
[5 8 4 5] 2  <-- 叶子节点
[5 8 4 1] 3  <-- 叶子节点
*/
var count int

func pathSum(root *TreeNode, sum int) int {
	count = 0
	rootPathSum(root, sum, []int{})
	return count
}

func rootPathSum(root *TreeNode, sum int, path []int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	var s int
	for i := len(path) - 1; i > 0; i-- {
		s += path[i]
		if s == sum {
			count++
		}
	}
	rootPathSum(root.Left, sum, path)
	rootPathSum(root.Right, sum, path)
	path = path[:len(path)-1]
	return
}

// @lc code=end

