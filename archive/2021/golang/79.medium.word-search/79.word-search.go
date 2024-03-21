/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (36.77%)
 * Likes:    5491
 * Dislikes: 242
 * Total Accepted:    644K
 * Total Submissions: 1.7M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given an m x n grid of characters board and a string word, return true if
 * word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cells,
 * where adjacent cells are horizontally or vertically neighboring. The same
 * letter cell may not be used more than once.
 *
 *
 * Example 1:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCCED"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "SEE"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word
 * = "ABCB"
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * m == board.length
 * n = board[i].length
 * 1 <= m, n <= 6
 * 1 <= word.length <= 15
 * board and word consists of only lowercase and uppercase English letters.
 *
 *
 *
 * Follow up: Could you use search pruning to make your solution faster with a
 * larger board?
 *
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	/*
		choose: 选 board 的某一个点出发; 如果匹配 word 的第一个字符，就进行
		constrain: 前进的方向不能超出board 边界; 不能重复使用字符
		capture: word 长度 == 0 时，就找到了答案
		unchoose
	*/
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if backtrack(board, x, y, word, visited) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, x, y int, word string, visited [][]bool) bool {
	if len(word) == 0 {
		return true
	}

	if y >= len(board) || x >= len(board[0]) || x < 0 || y < 0 { // out of bound
		return false
	}

	if board[y][x] != word[0] {
		return false
	}
	if visited[y][x] {
		return false
	}
	visited[y][x] = true
	defer func() { visited[y][x] = false }()
	// 如果 开始节点 == word[0]
	// 下面有四个方向可以走:
	// x+1,y; x-1,y;x,y+1;x,y-1
	/* x
	   1 2 3 y
	   4 5 6
	   7 8 9
	*/
	subWord := word[1:]
	return backtrack(board, x-1, y, subWord, visited) ||
		backtrack(board, x+1, y, subWord, visited) ||
		backtrack(board, x, y-1, subWord, visited) ||
		backtrack(board, x, y+1, subWord, visited)
}

// @lc code=end

