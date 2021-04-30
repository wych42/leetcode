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
		暴力解法没有用上重复元素都恰好出现三次这个条件

		2,2,3,2 为例:
		010
		010
		011
		010
		可以看出，除了3，其余元素的每一位，要么都是0，要么都是1
		把每一位都加起来 % 3, 的结果就是结果在这一位的取值

		怎么取一个二进制数的第 i 位呢？
		010 >> 1 -> 01
		01 & 1 -> 1

		知道一个数的第i位是什么，也知道前面 i-1 位的结果，怎么得到这个数呢？
		用 OR 运算
		0 | 1 = 1
		0 | 0 = 0
		1<<i -> 1000
		1000 | 0010 -> 1010
	*/
	ans := int32(0)
	for i := 0; i < 32; i++ {
		sum := int32(0)
		for _, num := range nums {
			sum += int32(num) >> i & 1
		}
		if sum%3 > 0 { // 结果的这一位是 1
			ans |= 1 << i
		}
	}
	return int(ans)
}

// @lc code=end

