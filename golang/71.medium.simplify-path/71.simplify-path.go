/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 *
 * https://leetcode.com/problems/simplify-path/description/
 *
 * algorithms
 * Medium (34.99%)
 * Likes:    360
 * Dislikes: 105
 * Total Accepted:    277.3K
 * Total Submissions: 790.9K
 * Testcase Example:  '"/home/"'
 *
 * Given a string path, which is an absolute path (starting with a slash '/')
 * to a file or directory in a Unix-style file system, convert it to the
 * simplified canonical path.
 *
 * In a Unix-style file system, a period '.' refers to the current directory, a
 * double period '..' refers to the directory up a level, and any multiple
 * consecutive slashes (i.e. '//') are treated as a single slash '/'. For this
 * problem, any other format of periods such as '...' are treated as
 * file/directory names.
 *
 * The canonical path should have the following format:
 *
 *
 * The path starts with a single slash '/'.
 * Any two directories are separated by a single slash '/'.
 * The path does not end with a trailing '/'.
 * The path only contains the directories on the path from the root directory
 * to the target file or directory (i.e., no period '.' or double period
 * '..')
 *
 *
 * Return the simplified canonical path.
 *
 *
 * Example 1:
 *
 *
 * Input: path = "/home/"
 * Output: "/home"
 * Explanation: Note that there is no trailing slash after the last directory
 * name.
 *
 *
 * Example 2:
 *
 *
 * Input: path = "/../"
 * Output: "/"
 * Explanation: Going one level up from the root directory is a no-op, as the
 * root level is the highest level you can go.
 *
 *
 * Example 3:
 *
 *
 * Input: path = "/home//foo/"
 * Output: "/home/foo"
 * Explanation: In the canonical path, multiple consecutive slashes are
 * replaced by a single one.
 *
 *
 * Example 4:
 *
 *
 * Input: path = "/a/./b/../../c/"
 * Output: "/c"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= path.length <= 3000
 * path consists of English letters, digits, period '.', slash '/' or '_'.
 * path is a valid absolute Unix path.
 *
 *
 */

// @lc code=start
import "strings"

func simplifyPath(path string) string {
	/*
		用 / split 输入;
		用栈保存遍历的结果；
		碰到：
		/ . 就跳过
		碰到 .. 就出栈
	*/
	segs := strings.Split(path, "/", -1)
	stack := make([]string, 0)
	for _, seg := range segs {
		switch seg {
		case "", ".", "/":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, seg)
		}
	}
	return "/" + strings.Join(stack, "/")
}

// @lc code=end

