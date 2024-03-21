/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays/description/
 *
 * algorithms
 * Hard (31.00%)
 * Likes:    9645
 * Dislikes: 1482
 * Total Accepted:    902.6K
 * Total Submissions: 2.9M
 * Testcase Example:  '[1,3]\n[2]'
 *
 * Given two sorted arrays nums1 and nums2 of size m and n respectively, return
 * the median of the two sorted arrays.
 *
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,3], nums2 = [2]
 * Output: 2.00000
 * Explanation: merged array = [1,2,3] and median is 2.
 *
 *
 * Example 2:
 *
 *
 * Input: nums1 = [1,2], nums2 = [3,4]
 * Output: 2.50000
 * Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 *
 *
 * Example 3:
 *
 *
 * Input: nums1 = [0,0], nums2 = [0,0]
 * Output: 0.00000
 *
 *
 * Example 4:
 *
 *
 * Input: nums1 = [], nums2 = [1]
 * Output: 1.00000
 *
 *
 * Example 5:
 *
 *
 * Input: nums1 = [2], nums2 = []
 * Output: 2.00000
 *
 *
 *
 * Constraints:
 *
 *
 * nums1.length == m
 * nums2.length == n
 * 0 <= m <= 1000
 * 0 <= n <= 1000
 * 1 <= m + n <= 2000
 * -10^6 <= nums1[i], nums2[i] <= 10^6
 *
 *
 *
 * Follow up: The overall run time complexity should be O(log (m+n)).
 */

// @lc code=start
func findMedianSortedArrays(a []int, b []int) float64 {
	/*
		参考: https://leetcode.com/problems/median-of-two-sorted-arrays/discuss/2481/Share-my-O(log(min(mn)))-solution-with-explanation
		中位数的作用是把列表分为数量相等的两份，并且左边都小于右边

		左       |		右
		a[0:i-1] |  a[i:m-1]
		b[0:j-1] |  b[j:n-1]

		问题就变成了，找到 i,j,使得
		1. i+j = n-i + m-j +1(偶数)
		2. a[i-i] < b[j] && b[j-1] < a[i]

		假设 a 的元素比较少，那么:
		1. i 一定在 [0,m] 之间, 即 a 全都比 b 大/小，也不足以影响到一半
		2. j = (m+n)/2-i
	*/
	m, n := len(a), len(b)
	if m > n {
		m, n = n, m
		a, b = b, a
	}
	imin, imax, half := 0, m, (m+n+1)/2
	for imin < imax {
		i := (imin + imax) / 2
		j := half - i
		if i < m && b[j-1] > a[i] {
			imin = i + 1
		} else if i > 0 && a[i-1] > b[j] {
			imax = i - 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = b[j-1]
			} else if j == 0 {
				maxLeft = a[i-1]
			} else {
				maxLeft = max(a[i-1], b[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			var minRight int
			if i == m {
				minRight = b[j]
			} else if j == n {
				minRight = a[i]
			} else {
				minRight = min(a[i], b[j])
			}
			return float64(maxLeft+minRight) / 2.0

		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

