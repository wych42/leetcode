/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] Valid Sudoku
 *
 * https://leetcode.cn/problems/valid-sudoku/description/
 *
 * algorithms
 * Medium (63.34%)
 * Likes:    1221
 * Dislikes: 0
 * Total Accepted:    426.6K
 * Total Submissions: 673.5K
 * Testcase Example:  '[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]'
 *
 * Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be
 * validated according to the following rules:
 *
 *
 * Each row must contain the digits 1-9 without repetition.
 * Each column must contain the digits 1-9 without repetition.
 * Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9
 * without repetition.
 *
 *
 * Note:
 *
 *
 * A Sudoku board (partially filled) could be valid but is not necessarily
 * solvable.
 * Only the filled cells need to be validated according to the mentioned
 * rules.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: board =
 * [["5","3",".",".","7",".",".",".","."]
 * ,["6",".",".","1","9","5",".",".","."]
 * ,[".","9","8",".",".",".",".","6","."]
 * ,["8",".",".",".","6",".",".",".","3"]
 * ,["4",".",".","8",".","3",".",".","1"]
 * ,["7",".",".",".","2",".",".",".","6"]
 * ,[".","6",".",".",".",".","2","8","."]
 * ,[".",".",".","4","1","9",".",".","5"]
 * ,[".",".",".",".","8",".",".","7","9"]]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: board =
 * [["8","3",".",".","7",".",".",".","."]
 * ,["6",".",".","1","9","5",".",".","."]
 * ,[".","9","8",".",".",".",".","6","."]
 * ,["8",".",".",".","6",".",".",".","3"]
 * ,["4",".",".","8",".","3",".",".","1"]
 * ,["7",".",".",".","2",".",".",".","6"]
 * ,[".","6",".",".",".",".","2","8","."]
 * ,[".",".",".","4","1","9",".",".","5"]
 * ,[".",".",".",".","8",".",".","7","9"]]
 * Output: false
 * Explanation: Same as Example 1, except with the 5 in the top left corner
 * being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it
 * is invalid.
 *
 *
 *
 * Constraints:
 *
 *
 * board.length == 9
 * board[i].length == 9
 * board[i][j] is a digit 1-9 or '.'.

 use hashmap check repetition.
 [int]map[byte]int: key: row number, value: map of numbers in row

 [int]map[byte]int: key: col number, value: map of numbers in row

 [int]map[byte]int: key: sub-box number, value: map of numbers in row
 0,0,   0,0 .  0,1 .  0,1 . 0,2 0,2
 [0,0 - 2,2], [0,3, - 2,5],[0,6-2,8]
 1,0 .  1,0 . 1,1 .    1,1, 1,2, 1,2
 [3,0 - 5,2], [3,3, - 5,5],[3,6-5,8]
 2,0,2,0 .    2,1 2,1       2,2 2,2
 [6,0 - 8,2], [6,3, - 8,5],[6,6-8,8]
 *



 [
[".",".",".",".","5",".",".","1","."],
[".","4",".","3",".",".",".",".","."],
[".",".",".",".",".","3",".",".","1"],
["8",".",".",".",".",".",".","2","."],
[".",".","2",".","7",".",".",".","."],
[".","1","5",".",".",".",".",".","."],
[".",".",".",".",".","2",".",".","."],[".","2",".","9",".",".",".",".","."],[".",".","4",".",".",".",".",".","."]]

 *
*/

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	rowMap := make(map[int]map[byte]int)
	colMap := make(map[int]map[byte]int)
	boxMap := make(map[int]map[byte]int)
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[byte]int)
		colMap[i] = make(map[byte]int)
		boxMap[i] = make(map[byte]int)

	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			value := board[r][c]
			if value < '0' || value > '9' {
				continue
			}
			if _, ok := rowMap[r][value]; ok {
				return false
			}
			rowMap[r][value] = 1
			if _, ok := colMap[c][value]; ok {
				return false
			}
			colMap[c][value] = 1

			boxIdx := (r/3)*3 + (c / 3)
			if _, ok := boxMap[boxIdx][value]; ok {
				return false
			}
			boxMap[boxIdx][value] = 1
		}
	}
	return true

}

// @lc code=end

