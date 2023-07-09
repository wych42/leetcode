/*
 * @lc app=leetcode.cn id=1 lang=java
 *
 * [1] Two Sum
 *
 * https://leetcode.cn/problems/two-sum/description/
 *
 * algorithms
 * Easy (52.82%)
 * Likes:    17281
 * Dislikes: 0
 * Total Accepted:    4.6M
 * Total Submissions: 8.7M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers nums and an integer target, return indices of the
 * two numbers such that they add up to target.
 * 
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 * 
 * You can return the answer in any order.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [2,7,11,15], target = 9
 * Output: [0,1]
 * Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [3,2,4], target = 6
 * Output: [1,2]
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: nums = [3,3], target = 6
 * Output: [0,1]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 2 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * Only one valid answer exists.
 * 
 * 
 * 
 * Follow-up: Can you come up with an algorithm that is less than O(n^2) time
 * complexity?
 */

// @lc code=start

import java.util.HashMap;
import java.util.Map;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        // use a map to store diff between target and current number as key, current idx as value
        // go through the nums again, if num exists in the diffMap, return the idx of current number and map value.
        Map<Integer, Integer> diffMap = new HashMap<>();
        for (int i=0;i<nums.length; i++) {
            diffMap.put(target-nums[i], i);
        }
        for (int i=0;i<nums.length;i++) {
            Integer pos = diffMap.get(nums[i]);
            // can't use same num twice
            if (pos != null && !pos.equals(i)) {
                return new int[]{i, diffMap.get(nums[i])};
            }
        }
        return new int[2];
    }
}
// @lc code=end

