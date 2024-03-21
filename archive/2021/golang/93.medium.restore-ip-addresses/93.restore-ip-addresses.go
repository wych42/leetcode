/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (37.49%)
 * Likes:    1718
 * Dislikes: 559
 * Total Accepted:    231.1K
 * Total Submissions: 610.3K
 * Testcase Example:  '"25525511135"'
 *
 * Given a string s containing only digits, return all possible valid IP
 * addresses that can be obtained from s. You can return them in any order.
 *
 * A valid IP address consists of exactly four integers, each integer is
 * between 0 and 255, separated by single dots and cannot have leading zeros.
 * For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses and
 * "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP
 * addresses.
 *
 *
 * Example 1:
 * Input: s = "25525511135"
 * Output: ["255.255.11.135","255.255.111.35"]
 * Example 2:
 * Input: s = "0000"
 * Output: ["0.0.0.0"]
 * Example 3:
 * Input: s = "1111"
 * Output: ["1.1.1.1"]
 * Example 4:
 * Input: s = "010010"
 * Output: ["0.10.0.10","0.100.1.0"]
 * Example 5:
 * Input: s = "101023"
 * Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
 *
 *
 * Constraints:
 *
 *
 * 0 <= s.length <= 3000
 * s consists of digits only.
 *
 *
 */

// @lc code=start

/*
回溯套路：
- choice：选一个开头的片头
- constrain: 检查约束条件
- goal: 检查是否满足目标条件; 满足条件之后就写入结果
*/
var res []string

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 { // 不可能
		return nil
	}
	res = make([]string, 0)
	// 从 index 0 开始尝试，计算第一个片段
	restore(s, 0, 0, make([]string, 4))
	return res
}

func restore(s string, buildIdx int, segment int, path []string) {
	if segment == 4 && buildIdx == len(s) { // 找到一个
		res = append(res, strings.Join(path, "."))
		return
	} else if segment == 4 || buildIdx == len(s) { // 拼了4个，字符串还有剩余，不处理了；到结尾了，还没够4个，也不处理了
		return
	}
	// 每一段最多有三个字符
	for l := 1; l <= 3; l++ {
		if buildIdx+l > len(s) { // 判断超出
			break
		}
		// 找到字串
		subStr := s[buildIdx : buildIdx+l]
		subInt, err := strconv.Atoi(subStr)
		if err != nil {
			panic(err)
		}
		// 判断约束条件
		if subInt > 255 || (l >= 2 && s[buildIdx] == '0') {
			break
		}
		// 符合条件，加到path里
		path[segment] = subStr
		// 算下一段
		restore(s, buildIdx+l, segment+1, path)
		// 取出数据，方便下一个循环用
		path[segment] = ""
	}
}

// @lc code=end

