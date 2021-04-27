/*
 * @lc app=leetcode id=1011 lang=golang
 *
 * [1011] Capacity To Ship Packages Within D Days
 *
 * https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/description/
 *
 * algorithms
 * Medium (59.99%)
 * Likes:    2018
 * Dislikes: 59
 * Total Accepted:    71.8K
 * Total Submissions: 119.7K
 * Testcase Example:  '[1,2,3,4,5,6,7,8,9,10]\n5'
 *
 * A conveyor belt has packages that must be shipped from one port to another
 * within D days.
 *
 * The i^th package on the conveyor belt has a weight of weights[i]. Each day,
 * we load the ship with packages on the conveyor belt (in the order given by
 * weights). We may not load more weight than the maximum weight capacity of
 * the ship.
 *
 * Return the least weight capacity of the ship that will result in all the
 * packages on the conveyor belt being shipped within D days.
 *
 *
 * Example 1:
 *
 *
 * Input: weights = [1,2,3,4,5,6,7,8,9,10], D = 5
 * Output: 15
 * Explanation: A ship capacity of 15 is the minimum to ship all the packages
 * in 5 days like this:
 * 1st day: 1, 2, 3, 4, 5
 * 2nd day: 6, 7
 * 3rd day: 8
 * 4th day: 9
 * 5th day: 10
 *
 * Note that the cargo must be shipped in the order given, so using a ship of
 * capacity 14 and splitting the packages into parts like (2, 3, 4, 5), (1, 6,
 * 7), (8), (9), (10) is not allowed.
 *
 *
 * Example 2:
 *
 *
 * Input: weights = [3,2,2,4,1,4], D = 3
 * Output: 6
 * Explanation: A ship capacity of 6 is the minimum to ship all the packages in
 * 3 days like this:
 * 1st day: 3, 2
 * 2nd day: 2, 4
 * 3rd day: 1, 4
 *
 *
 * Example 3:
 *
 *
 * Input: weights = [1,2,3,1,1], D = 4
 * Output: 3
 * Explanation:
 * 1st day: 1
 * 2nd day: 2
 * 3rd day: 3
 * 4th day: 1, 1
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= D <= weights.length <= 5 * 10^4
 * 1 <= weights[i] <= 500
 *
 *
 */

// @lc code=start
func shipWithinDays(weights []int, d int) int {
	/*
		要找的数是载重量 w，最大 sum(weights), days=1; 最小 max(weights), days=len(weights)
		现在要找一个 w，使得 days=d
		并且包裹不能拆分
	*/
	sum := 0
	max := 0
	for _, weight := range weights {
		if weight > max {
			max = weight
		}
		sum += weight
	}
	if d == 1 {
		return sum
	}
	if d == len(weights) {
		return max
	}
	res := sum
	start, end := max, sum
	for start+1 < end {
		mid := (end-start)/2 + start
		days := getDays(weights, mid)
		if days == d { // 找到了一个解答，但可能有更小的，试试看
			if mid < res {
				res = mid
			}
			end = mid
		}
		if days > d { // 要运的天数比目标多，说明载重量小了，应该加大
			start = mid
		} else {
			end = mid
		}
	}
	//fmt.Println(start, end, res, getDays(weights, start))
	if start < res { // 检查边界情况
		if getDays(weights, start) <= d {
			res = start
		}
	}
	if end < res { // 检查边界情况
		if getDays(weights, end) <= d {
			res = end
		}
	}
	return res
}

func getDays(weights []int, cap int) int {
	temp := 0
	days := 0
	for _, w := range weights {
		if temp+w > cap { // 加上之后，超出载重量，就要多运一趟
			days++
			temp = w
		} else {
			temp += w
		}
	}

	days++ // 最后一趟要加上
	return days
}

// @lc code=end

