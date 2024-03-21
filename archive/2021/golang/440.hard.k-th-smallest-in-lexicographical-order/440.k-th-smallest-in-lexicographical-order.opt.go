/*
 * @lc app=leetcode id=440 lang=golang
 *
 * [440] K-th Smallest in Lexicographical Order
 *
 * https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/description/
 *
 * algorithms
 * Hard (29.79%)
 * Likes:    415
 * Dislikes: 60
 * Total Accepted:    15K
 * Total Submissions: 50.2K
 * Testcase Example:  '13\n2'
 *
 * Given integers n and k, find the lexicographically k-th smallest integer in
 * the range from 1 to n.
 *
 * Note: 1 ≤ k ≤ n ≤ 10^9.
 *
 * Example:
 *
 * Input:
 * n: 13   k: 2
 *
 * Output:
 * 10
 *
 * Explanation:
 * The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so
 * the second smallest number is 10.
 *
 *
 *
 */

// @lc code=start
/*
规律:
1 10 11 12 13 ...
2 20 21 22 23 ...
3 30 31 32 33 ...

从 i=1 开始，找到 i 开头，并且小于 n 的数个数(count); 把 count 与 比较，如果没到，就找下一个
*/

func findKthNumber(n int, k int) int {
	prefix := 1
	idx := 1
	for idx < k {
		c := count(prefix, n) // 找到以 prefix 开头的数字，比n小的有多少个
		if idx+c > k {        // 说明结果落在 prefix 开头的数字里; 从 10 开始找
			prefix *= 10
			idx++ // 步数+1
			continue
		}
		if idx+c <= k { // 说明结果落在其他 prefix 开头的数字里
			prefix++
			idx += c // 部署都算上
			continue
		}
	}
	return prefix
}

func count(prefix int, n int) int {
	// prefix = 1;n = 13
	// a=1;b=2
	// min(13,2) - 1 // 自己
	// min(13,20)-10 // 两位数
	//                  三位数

	// 以 i 开头的数有 1(i自己)+10*i(十位数)+10*i(百位数)...个
	// 要找 i 开头，并且小于 n 的数量
	// 比较 [i,i+1] -> 1,2
	// 比较 [10*i, 10*(i+1)] -> 10,20
	// 如果 n 在 [i*10^j,(i+1)*10^j] 之间，那么个数就是 n-i*10^j+1
	// 否则就是 (i+1)*10^j-i*10^j
	c := 0
	next := prefix + 1
	for prefix <= n {
		c += min(n+1, next) - prefix
		next *= 10
		prefix *= 10
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

