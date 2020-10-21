/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+","-","*","/":
			if len(stack) < 2 {
				return -1
			}
			operator := stack[len(stack) - 1]
			operand := stack[len(stack) - 2]
			var v int
			switch token {
			case "+":
				v = operand + operator
			case "-":
				v = operand - operator
			case "*":
				v = operand * operator
			case "/":
				v = operand / operator
			default:
				return -1
			}
			stack = append(stack[:len(stack)-2], v)
		default:
			v, _ := strconv.Atoi(token)
			stack = append(stack, v)
		}
	}
	return stack[0]
}
// @lc code=end

