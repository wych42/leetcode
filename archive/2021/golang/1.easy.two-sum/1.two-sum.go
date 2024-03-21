/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	diffMap := make(map[int]int)
	for i, num := range nums {
		if idx, ok := diffMap[num]; ok {
			return []int{idx, i}
		} else {
			diffMap[target-num] = i
		}
	}
	return nil
}

// @lc code=end

