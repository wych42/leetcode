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
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	k := l / 2 // 中间位置

	i, j := 0, 0
	pair := [2]int{}
	cnt := 0
	for i < len(nums1) && j < len(nums2) {
		var min int
		if nums1[i] < nums2[j] { // 小的放前面
			min = nums1[i]
			i++
		} else {
			min = nums2[j]
			j++
		}
		if cnt == k-1 {
			pair[0] = min
		}
		if cnt == k {
			pair[1] = min
		}
		cnt++
	}
	for ; i < len(nums1) && cnt <= k+1; i++ {
		if cnt == k-1 {
			pair[0] = nums1[i]
		}
		if cnt == k {
			pair[1] = nums1[i]
		}
		cnt++
	}
	for ; j < len(nums2) && cnt <= k+1; j++ {
		if cnt == k-1 {
			pair[0] = nums2[j]
		}
		if cnt == k {
			pair[1] = nums2[j]
		}
		cnt++
	}
	if l%2 == 0 {
		return float64(pair[0]+pair[1]) / float64(2)
	}
	return float64(pair[1])
}

// @lc code=end

