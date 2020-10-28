package golang

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	stack := make([]rune, 0)
	for _, c := range s {
		switch c {
		case ']':
			stringStack := make([]rune, 0)
			countStack := make([]rune, 0)
			var stringCompelete bool
			var idx int
			for idx = len(stack) - 1; idx >= 0; idx-- {
				v := stack[idx]
				if v == '[' {
					if stringCompelete {
						break
					}
					stringCompelete = true
					continue
				}
				if !stringCompelete {
					stringStack = append(stringStack, v)
					continue
				}
				if stringCompelete && v <= '9' && v >= '0' {
					countStack = append(countStack, v)
					continue
				}
				break
			}
			stack = stack[:idx+1]
			fmt.Println("xx", string(stack))
			numSlice := make([]rune, len(countStack))
			for i := 0; i <= len(countStack)-1; i++ {
				numSlice[i] = countStack[len(countStack)-1-i]
			}
			count, _ := strconv.Atoi(string(numSlice))
			for i := 0; i < count; i++ {
				for j := len(stringStack) - 1; j >= 0; j-- {
					stack = append(stack, stringStack[j])
				}
			}
			fmt.Println("yy", string(stack))
		default:
			stack = append(stack, c)
		}

	}
	return string(stack)

}

// @lc code=end
