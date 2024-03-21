/*
 * @lc app=leetcode id=54 lang=java
 *
 * [54] Spiral Matrix
 *
 * https://leetcode.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (35.93%)
 * Likes:    3717
 * Dislikes: 660
 * Total Accepted:    488.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given an m x n matrix, return all elements of the matrix in spiral order.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [1,2,3,6,9,8,7,4,5]
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * Output: [1,2,3,4,8,12,11,10,9,5,6,7]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * m == matrix.length
 * n == matrix[i].length
 * 1 <= m, n <= 10
 * -100 <= matrix[i][j] <= 100
 * 
 * 
 */

// @lc code=start
class Solution {
    public List<Integer> spiralOrder(int[][] matrix) {
        int rows = matrix.length;
        int cols = matrix[0].length;
        List<Integer> ans = new ArrayList<Integer>(cols*rows);
        int left = 0;
        int right = cols;
        int top = 0;
        int bottom = rows;
        while (left<right && top<bottom){
            for (int i=left;i<right;i++) {
                ans.add(matrix[top][i]); // top row
            }
            top++;
            for (int i=top;i<bottom;i++) {
                ans.add(matrix[i][right-1]); // right column
            }
            right--;
            if (left>=right||top>=bottom) {
                break;
            }
            for (int i=right-1;i>=left;i--) {
                ans.add(matrix[bottom-1][i]); // bottom row
            }
            bottom--;
            for (int i=bottom-1;i>=top;i--) {
                ans.add(matrix[i][left]); // left row
            }
            left++;
        }
        return ans;
    }
}
// @lc code=end

