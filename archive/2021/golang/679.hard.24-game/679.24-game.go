/*
 * @lc app=leetcode id=679 lang=golang
 *
 * [679] 24 Game
 *
 * https://leetcode.com/problems/24-game/description/
 *
 * algorithms
 * Hard (47.32%)
 * Likes:    932
 * Dislikes: 186
 * Total Accepted:    50.6K
 * Total Submissions: 106.8K
 * Testcase Example:  '[4,1,8,7]'
 *
 * You are given an integer array cards of length 4. You have four cards, each
 * containing a number in the range [1, 9]. You should arrange the numbers on
 * these cards in a mathematical expression using the operators ['+', '-', '*',
 * '/'] and the parentheses '(' and ')' to get the value 24.
 *
 * You are restricted with the following rules:
 *
 *
 * The division operator '/' represents real division, not integer
 * division.
 *
 *
 * For example, 4 / (1 - 2 / 3) = 4 / (1 / 3) = 12.
 *
 *
 * Every operation done is between two numbers. In particular, we cannot use
 * '-' as a unary operator.
 *
 * For example, if cards = [1, 1, 1, 1], the expression "-1 - 1 - 1 - 1" is not
 * allowed.
 *
 *
 * You cannot concatenate numbers together
 *
 * For example, if cards = [1, 2, 1, 2], the expression "12 + 12" is not
 * valid.
 *
 *
 *
 *
 * Return true if you can get such expression that evaluates to 24, and false
 * otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: cards = [4,1,8,7]
 * Output: true
 * Explanation: (8-4) * (7-1) = 24
 *
 *
 * Example 2:
 *
 *
 * Input: cards = [1,2,1,2]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * cards.length == 4
 * 1 <= cards[i] <= 9
 *
 *
 */

// @lc code=start
func judgePoint24(cards []int) bool {
	/*
		用回溯做：
		- 取两个数，有 + - * / 四种运算，得到一个数，替换原来两个数
		- 对三个数做计算...
		- 还有一个数的时候，检查
	*/
	floatCards := make([]float32, len(cards))
	for i, card := range cards {
		floatCards[i] = float32(card)
	}
	return backtrack(floatCards)
}

func backtrack(cards []float32) bool {
	// capture
	if len(cards) == 1 {
		return cards[0] >= 23.99 && cards[0] <= 24.01
	}

	// 挑两个数
	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			tempCards := make([]float32, 0, len(cards)-1)
			tempCards = append(tempCards, cards[:i]...)    // (:i)
			tempCards = append(tempCards, cards[i+1:j]...) // (i:j)
			tempCards = append(tempCards, cards[j+1:]...)  // (j:)

			ni := cards[i]
			nj := cards[j]
			//            fmt.Println(tempCards, ni, nj)
			isValid := backtrack(append(tempCards, ni+nj)) ||
				backtrack(append(tempCards, ni*nj)) ||
				backtrack(append(tempCards, ni/nj)) ||
				backtrack(append(tempCards, nj/ni)) ||
				backtrack(append(tempCards, ni-nj)) ||
				backtrack(append(tempCards, nj-ni))

			if isValid {
				return true
			}
		}
	}
	return false
}

// @lc code=end

