/*
 * @lc app=leetcode id=51 lang=golang
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
import "strings"
import "math"

var answerSets [][]int

func solveNQueens(n int) [][]string {
	/*
		backtrack 问题：
		1. 一行一行的摆
		2. 对每一行，选择一列放上；
		3. 尝试摆下一行；
	*/
	// 每个Queen放一行，每个 []int 记录每一行上 Queen 所在列
	// cols 就是结果集合
	//tmpl = strings.Repeat(".", n)
	answerSets = make([][]int, 0)
	backtrack(n, 0, []int{})
	result := make([][]string, 0, len(answerSets))
	for _, answer := range answerSets {
		matrix := make([]string, n)
		for row, col := range answer {
			matrix[row] = strings.Repeat(".", col) + "Q" + strings.Repeat(".", n-col-1)
		}
		result = append(result, matrix)
	}
	return result
}

func backtrack(n int, row int, path []int) {
	/*
		n: 总行数
		row: 当前处理哪一行
		path: 这一次递归里，每一行已经摆放的 Queen 所在列
		ans: 结果
	*/
	// Goal
	if row == n { // 所有行都摆上了
		//fmt.Println("found answer", path)
		answer := make([]int, len(path))
		// golang 语法的坑，这里不能直接加到结果里
		for i := 0; i < len(path); i++ {
			answer[i] = path[i]
		}
		answerSets = append(answerSets, answer)
		//fmt.Println(answerSets)
		return
	}
	// 在当前行上，尝试在每一列上摆 Queen
	for col := 0; col < n; col++ {
		// choose
		path = append(path, col)
		// constrain
		if isValid(path) {
			// explore
			backtrack(n, row+1, path)
		}
		// unchoose
		path = path[:len(path)-1]
	}
}

func isValid(path []int) bool {
	nextRow := len(path) - 1 // 当前要摆放的行; path 里有两个元素，则现在要摆放第三行(idx=2)
	nextCol := path[nextRow]
	for row := 0; row < nextRow; row++ {
		col := path[row]
		if nextCol == col {
			return false // 同一列
		}
		// 斜线
		if math.Abs(float64(nextCol-col)) == math.Abs(float64(nextRow-row)) {
			return false
		}
	}
	return true
}

// @lc code=end

