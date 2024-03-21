import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode id=6 lang=java
 *
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (37.89%)
 * Likes:    2299
 * Dislikes: 5716
 * Total Accepted:    562.2K
 * Total Submissions: 1.5M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 * 
 * 
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 * 
 * 
 * And then read line by line: "PAHNAPLSIIGYIR"
 * 
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 * 
 * 
 * string convert(string s, int numRows);
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "PAYPALISHIRING", numRows = 3
 * Output: "PAHNAPLSIIGYIR"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "PAYPALISHIRING", numRows = 4
 * Output: "PINALSIGYAHRPI"
 * Explanation:
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "A", numRows = 1
 * Output: "A"
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 1000
 * s consists of English letters (lower-case and upper-case), ',' and '.'.
 * 1 <= numRows <= 1000
 * 
 * 
 */

// @lc code=start
class Solution {
    public String convert(String s, int numRows) {
        if (numRows == 1) {
            return s;
        }
        List<StringBuilder> grid = new ArrayList<>();
        for (int i=0;i<numRows;i++){
            grid.add(new StringBuilder());
        }
        boolean downwards = false;
        int row = 0;
        for (char ch: s.toCharArray()) {
            grid.get(row).append(ch);
            if (row==numRows-1) {
                downwards = false;
            } else if (row==0) {
                downwards = true;
            }
            if (downwards) {
                row++;
            } else {
                row--;
            }
        }
        StringBuilder ans = new StringBuilder();
        for (StringBuilder r: grid) {
            ans.append(r);
        }
        return ans.toString();
    }
}
// @lc code=end

