import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode id=51 lang=java
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (50.36%)
 * Likes:    2860
 * Dislikes: 106
 * Total Accepted:    251.2K
 * Total Submissions: 498.6K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n x n
 * chessboard such that no two queens attack each other.
 * 
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 * 
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space,
 * respectively.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: n = 4
 * Output: [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: n = 1
 * Output: [["Q"]]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= n <= 9
 * 
 * 
 */

// @lc code=start
class Solution {
    public List<List<String>> solveNQueens(int n) {
        List<List<String>> answers = new ArrayList<>(); 
        this.backtrack(n, 0, new ArrayList<>(), answers);
        return answers;
    }

    public void backtrack(int n, int row, List<Integer> path, List<List<String>> answers) {
        if (row == n) {
            answers.add(this.pathToAnswer(path));
            return;
        }
        for (int col = 0; col < n; col++) {
            path.add(col);
            if (this.isValid(path)) {
                this.backtrack(n, row+1, path, answers);
            }
            path.remove(path.size()-1);
        }
    }

    public Boolean isValid(List<Integer> path) {
        int currRow = path.size()-1;
        int currCol = path.get(currRow);
        for (int row = 0; row < currRow; row++) {
            int col = path.get(row);
            if (currCol == col) {
                return false;
            }
            if (Math.abs(currCol - col) == Math.abs(currRow - row)) {
                return false;
            }
        }
        return true;
    }

    public List<String> pathToAnswer(List<Integer> path) {
        int len = path.size();
        List<String> answer = new ArrayList<>(len);
        for (int row = 0; row < len; row++) {
            int col = path.get(row);
            answer.add(".".repeat(col) + "Q" + ".".repeat(len - col - 1));
        }
        return answer;
    }
}
// @lc code=end
