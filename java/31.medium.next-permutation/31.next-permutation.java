/*
 * @lc app=leetcode id=31 lang=java
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (33.77%)
 * Likes:    5325
 * Dislikes: 1828
 * Total Accepted:    501K
 * Total Submissions: 1.5M
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 * 
 * If such an arrangement is not possible, it must rearrange it as the lowest
 * possible order (i.e., sorted in ascending order).
 * 
 * The replacement must be in place and use only constant extra memory.
 * 
 * 
 * Example 1:
 * Input: nums = [1,2,3]
 * Output: [1,3,2]
 * Example 2:
 * Input: nums = [3,2,1]
 * Output: [1,2,3]
 * Example 3:
 * Input: nums = [1,1,5]
 * Output: [1,5,1]
 * Example 4:
 * Input: nums = [1]
 * Output: [1]
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 100
 * 
 * 
 */

// @lc code=start
class Solution {
    public void nextPermutation(int[] nums) {
        int idx  = nums.length-2;
        for (;idx >= 0 && nums[idx] >= nums[idx+1];idx--) {
        }
        if (idx >= 0) {
            for (int j = nums.length-1;j>idx;j--) {
                if (nums[j] > nums[idx]) {
                    int temp = nums[j];
                    nums[j] = nums[idx];
                    nums[idx] = temp;
                    break;
                } 
            }
        }
        for (int l=idx+1, r=nums.length-1;l<r;l++,r--) {
            int temp = nums[l];
            nums[l] = nums[r];
            nums[r] = temp;
        }
    }
}
// @lc code=end

