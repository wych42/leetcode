/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 *
 * https://leetcode-cn.com/problems/single-number-ii/description/
 *
 * algorithms
 * Medium (68.73%)
 * Likes:    613
 * Dislikes: 0
 * Total Accepted:    72.7K
 * Total Submissions: 102.6K
 * Testcase Example:  '[2,2,3,2]'
 *
 * 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [2,2,3,2]
 * 输出：3
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0,1,0,1,0,1,99]
 * 输出：99
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -2^31
 * nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次
 *
 *
 *
 *
 * 进阶：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
 *
 */

// @lc code=start
func singleNumber(nums []int) int {
	/*
		暴力解法: 用两个 map, 一个记录元素出现的次数，另一个记录到目前为止出现一次的元素
	*/
	countMap := make(map[int]int, len(nums))
	onceMap := make(map[int]bool)
	for _, num := range nums {
		countMap[num]++
		c := countMap[num]
		if c == 1 { // 第一次出现
			onceMap[num] = true
		} else if c == 2 {
			delete(onceMap, num) // 出现重复，删掉
		}
	}
	for k, v := range onceMap {
		if v {
			return k
		}
	}
	return 0
}

// @lc code=end

