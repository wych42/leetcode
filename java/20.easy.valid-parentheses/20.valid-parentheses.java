/*
 * @lc app=leetcode.cn id=20 lang=java
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.cn/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (44.03%)
 * Likes:    4003
 * Dislikes: 0
 * Total Accepted:    1.5M
 * Total Submissions: 3.4M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * Every close bracket has a corresponding open bracket of the same type.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "()"
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "()[]{}"
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "(]"
 * Output: false
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 * 
 * 
 */

// @lc code=start
class Solution {
    // use stack
    // put open parenthes into stack
    // when meets close parenthes, pop stack top element, check if they match
    // if not match, return false
    // else continue
    public boolean isValid(String s) {
        char[] stack = new char[s.length()];
        int stackCur = 0;
        for (int i=0;i<s.length();i++) {
            char c = s.charAt(i);
            if (c=='(' || c=='[' || c=='{') {
                stack[stackCur] = c;
                stackCur++;
                continue;
            }
            if (c==')' || c==']' || c=='}') {
                // stack has no element, a close parenthes won't match
                if (stackCur==0) {
                    return false;
                }
                char stackTop = stack[stackCur-1];
                if (!isMatchParenthes(stackTop , c)) {
                    return false;
                }
                stackCur--;
            }
        }
        // stack should be empty if all matches
        return stackCur==0;
    }
    private boolean isMatchParenthes(char open, char close) {
        return (open == '(' && close == ')') 
        || (open == '[' && close == ']')
        || (open == '{' && close == '}');
    }
}
// @lc code=end

