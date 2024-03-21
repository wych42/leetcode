/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.cn/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (63.60%)
 * Likes:    1802
 * Dislikes: 0
 * Total Accepted:    524.9K
 * Total Submissions: 825.3K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given an integer array nums and an integer k, return the k most frequent
 * elements. You may return the answer in any order.
 *
 *
 * Example 1:
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 * Example 2:
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 * k is in the range [1, the number of unique elements in the array].
 * It is guaranteed that the answer is unique.
 *
 *
 *
 * Follow up: Your algorithm's time complexity must be better than O(n log n),
 * where n is the array's size.

 worst case: O(nlogn), k=1, and only one element is duplicated
 use a hash map track frequency of each elements.
 sort values of the hash map.
 return top K keys.

 --
 use a map track freq: number
 track the max freq while counting
 go from max frequency, pick corresponding numbers, append to result, also need check if result count > K

 *
*/

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	// num : freq
	// O(n)
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}
	// freq : num
	maxFreq := 0
	freqMap := make(map[int][]int)
	// O(n)
	for k, v := range numMap {
		freqMap[v] = append(freqMap[v], k)
		if v > maxFreq {
			maxFreq = v
		}
	}
	// O(K)
	result := make([]int, 0)
	for i := maxFreq; i > 0; i-- {
		if len(result) >= k {
			break
		}
		elements, ok := freqMap[i]
		if !ok {
			continue
		}
		for _, element := range elements {
			if len(result) >= k {
				break
			}
			result = append(result, element)
		}
	}
	return result

}

// @lc code=end

