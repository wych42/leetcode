/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 *
 * https://leetcode.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (40.92%)
 * Likes:    3703
 * Dislikes: 5273
 * Total Accepted:    860.7K
 * Total Submissions: 2.1M
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as
 * one sorted array.
 *
 * The number of elements initialized in nums1 and nums2 are m and n
 * respectively. You may assume that nums1 has a size equal to m + n such that
 * it has enough space to hold additional elements from nums2.
 *
 *
 * Example 1:
 * Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
 * Output: [1,2,2,3,5,6]
 * Example 2:
 * Input: nums1 = [1], m = 1, nums2 = [], n = 0
 * Output: [1]
 *
 *
 * Constraints:
 *
 *
 * nums1.length == m + n
 * nums2.length == n
 * 0 <= m, n <= 200
 * 1 <= m + n <= 200
 * -10^9 <= nums1[i], nums2[i] <= 10^9
 *
 *
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	/*
		从后往前数，用大的数填充空格
	*/
	idx := m + n - 1
	// 定位最后一个元素
	m--
	n--
	for n >= 0 {
		for m >= 0 && nums1[m] > nums2[n] { // 把 nums1 里大的数先放进去
			nums1[idx], nums1[m] = nums1[m], nums1[idx]
			idx--
			m--
		}
		// 否则就把 nums2 的数放到 idx
		nums1[idx], nums2[n] = nums2[n], nums1[idx]
		idx--
		n--
	}
}

// @lc code=end

