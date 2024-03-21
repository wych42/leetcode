/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 *
 * https://leetcode.com/problems/unique-binary-search-trees/description/
 *
 * algorithms
 * Medium (54.39%)
 * Likes:    4346
 * Dislikes: 162
 * Total Accepted:    350K
 * Total Submissions: 643K
 * Testcase Example:  '3'
 *
 * Given an integer n, return the number of structurally unique BST's (binary
 * search trees) which has exactly n nodes of unique values from 1 to n.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3
 * Output: 5
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 19
 *
 *
 */

// @lc code=start
var cache map[int]int

func numTrees(n int) int {
	/*
		暴力解法+缓存：
		节点取值区间是 [1:n]
		每一个节点都可以成为 root. 取 i 为当前的 root
		[1:i) 为左子树；(i:n] 为右子树
		递归取左右子树的数量，相乘
	*/
	cache := make(map[int]int, n)
	return numTreesByRange(1, n)
}

func numTreesByRange(s, e int) int {
	if s > e {
		return 0
	}
	if s == e {
		return 1
	}
	if v, ok := cache[e-s]; ok {
		return v
	}
	count := 0
	for i := s; i <= e; i++ {
		lc := numTreesByRange(s, i-1)
		rc := numTreesByRange(i+1, e)
		if lc == 0 {
			count += rc
		}
		if rc == 0 {
			count += lc
		}
		count += lc * rc
	}
	cache[e-s] = count
	return count
}

// @lc code=end

