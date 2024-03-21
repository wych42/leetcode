/*
 * @lc app=leetcode id=1122 lang=golang
 *
 * [1122] Relative Sort Array
 *
 * https://leetcode.com/problems/relative-sort-array/description/
 *
 * algorithms
 * Easy (68.07%)
 * Likes:    1040
 * Dislikes: 73
 * Total Accepted:    90.7K
 * Total Submissions: 133K
 * Testcase Example:  '[2,3,1,3,2,4,6,7,9,2,19]\n[2,1,4,3,9,6]'
 *
 * Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all
 * elements in arr2 are also in arr1.
 *
 * Sort the elements of arr1 such that the relative ordering of items in arr1
 * are the same as in arr2.  Elements that don't appear in arr2 should be
 * placed at the end of arr1 in ascending order.
 *
 *
 * Example 1:
 * Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
 * Output: [2,2,2,1,4,3,3,9,6,7,19]
 *
 *
 * Constraints:
 *
 *
 * 1 <= arr1.length, arr2.length <= 1000
 * 0 <= arr1[i], arr2[i] <= 1000
 * All the elements of arr2 are distinct.
 * Each arr2[i] is in arr1.
 *
 *
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	res := make([]int, 0, len(arr1))
	// 记录 arr1 里每个元素的次数
	m := make(map[int]int, len(arr1))
	for _, n := range arr1 {
		m[n]++
	}
	for _, n := range arr2 {
		for i := 0; i < m[n]; i++ {
			res = append(res, n)
		}
		delete(m, n)
	}
	temp := make([]int, 0)
	for k := range m {
		for i := 0; i < m[k]; i++ {
			temp = append(temp, k)
		}
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	res = append(res, temp...)
	return res
}

// @lc code=end

