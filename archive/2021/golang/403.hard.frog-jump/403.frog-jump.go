/*
 * @lc app=leetcode id=403 lang=golang
 *
 * [403] Frog Jump
 *
 * https://leetcode.com/problems/frog-jump/description/
 *
 * algorithms
 * Hard (41.74%)
 * Likes:    1549
 * Dislikes: 131
 * Total Accepted:    120.7K
 * Total Submissions: 288.9K
 * Testcase Example:  '[0,1,3,5,6,8,12,17]'
 *
 * A frog is crossing a river. The river is divided into some number of units,
 * and at each unit, there may or may not exist a stone. The frog can jump on a
 * stone, but it must not jump into the water.
 *
 * Given a list of stones' positions (in units) in sorted ascending order,
 * determine if the frog can cross the river by landing on the last stone.
 * Initially, the frog is on the first stone and assumes the first jump must be
 * 1 unit.
 *
 * If the frog's last jump was k units, its next jump must be either k - 1, k,
 * or k + 1 units. The frog can only jump in the forward direction.
 *
 *
 * Example 1:
 *
 *
 * Input: stones = [0,1,3,5,6,8,12,17]
 * Output: true
 * Explanation: The frog can jump to the last stone by jumping 1 unit to the
 * 2nd stone, then 2 units to the 3rd stone, then 2 units to the 4th stone,
 * then 3 units to the 6th stone, 4 units to the 7th stone, and 5 units to the
 * 8th stone.
 *
 *
 * Example 2:
 *
 *
 * Input: stones = [0,1,2,3,4,8,9,11]
 * Output: false
 * Explanation: There is no way to jump to the last stone as the gap between
 * the 5th and 6th stone is too large.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= stones.length <= 2000
 * 0 <= stones[i] <= 2^31 - 1
 * stones[0] == 0
 *
 *
 */

// @lc code=start
func canCross(stones []int) bool {
	/*
		如果要跳到节点 i, 那么从 i-1 往前找 j，符合条件的结果:
		- 差值 steps = stones[i] - stones[j]
		- 跳到 j 的位置时，可能有[m, n...]多种路径
		- 对于每种路径 k, 如果 steps = k-1/k/k+1 那么就是一个合理路径
		- 记录下来
		- 如果找不到路径，那么说明跳不到 i
	*/
	dp := make([]map[int]bool, len(stones))
	dp[0] = map[int]bool{1: true} // 第一步只能跳一个单位
	for i := 1; i < len(stones); i++ {
		m := make(map[int]bool)
		for j := i - 1; j >= 0; j-- {
			steps := stones[i] - stones[j] // 要能跳 steps 单位, 才能从 j 到 i
			if dp[j][steps] {              // 能做到
				m[steps-1] = true
				m[steps] = true
				m[steps+1] = true
			}
		}
		dp[i] = m
	}
	return len(dp[len(stones)-1]) > 0
}

// @lc code=end

