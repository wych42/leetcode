/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 *
 * https://leetcode.com/problems/edit-distance/description/
 *
 * algorithms
 * Hard (46.68%)
 * Likes:    5434
 * Dislikes: 67
 * Total Accepted:    347.6K
 * Total Submissions: 737.7K
 * Testcase Example:  '"horse"\n"ros"'
 *
 * Given two strings word1 and word2, return the minimum number of operations
 * required to convert word1 to word2.
 *
 * You have the following three operations permitted on a word:
 *
 *
 * Insert a character
 * Delete a character
 * Replace a character
 *
 *
 *
 * Example 1:
 *
 *
 * Input: word1 = "horse", word2 = "ros"
 * Output: 3
 * Explanation:
 * horse -> rorse (replace 'h' with 'r')
 * rorse -> rose (remove 'r')
 * rose -> ros (remove 'e')
 *
 *
 * Example 2:
 *
 *
 * Input: word1 = "intention", word2 = "execution"
 * Output: 5
 * Explanation:
 * intention -> inention (remove 't')
 * inention -> enention (replace 'i' with 'e')
 * enention -> exention (replace 'n' with 'x')
 * exention -> exection (replace 'n' with 'c')
 * exection -> execution (insert 'u')
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= word1.length, word2.length <= 500
 * word1 and word2 consist of lowercase English letters.
 *
 *
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	/*
		操作:
			replace;
			delete;
			insert
		检查字符串的末尾元素:
		    exp: beny, eph
		    - 如果相等，不用操作, 计算 word1[0:i-1], word2[0:j-1]
		    - 不相等:
			  - replace: 1+ word1[0:i-1],word2[0:j-1]
			      计算两个字串，做替换; 计算 ben,ep
			  - insert:  1+ word1[0:i], word2[0:j-1]
			      计算 1 和 2的字串，再插入2的最后一位; 计算 beny,ep
			  - delete: 1 + word1[0:i-1], word2[0:j]
			      把word1的最后一位删除,计算 word1字串和word2; 计算 ben;eph
			dp[i][j] = min(dp[i-1][j-1], 1+dp[i-1][j-1], 1+dp[i][j-1], 1+dp[i-1][j])
			               相等           replace         插入          删除

		    |replace| insert |
			|delete | current|
			     "" b e n y
			""   0  1 2 3 4      b,e 不相等,此处做替换，并加上 """" 的结果;
			e    1  1 1 2        e,e 相等，用b,e结果;
			p    2               e,n不相等, replace: be,""; insert: ben,"",delete:be,e;
			h    3


	*/
	dp := make([][]int, len(word2)+1) // word2 竖着排, 有 len(word2)+1 行
	for i := 0; i <= len(word2); i++ {
		dp[i] = make([]int, len(word1)+1) // 每行有 len(word1)+1 列
		dp[i][0] = i                      // 每行第一个，等于行数
	}
	for i := 0; i <= len(word1); i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len(word2); i++ { // 从第一行开始计算
		for j := 1; j <= len(word1); j++ { // 计算每一列
			if word2[i-1] == word1[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// replace
				replace := 1 + dp[i-1][j-1]
				// insert
				insert := 1 + dp[i][j-1]
				// delete
				delete := 1 + dp[i-1][j]
				dp[i][j] = min(replace, min(insert, delete))
			}
		}
	}
	return dp[len(word2)][len(word1)]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

