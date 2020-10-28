/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 *
 * https://leetcode-cn.com/problems/validate-binary-search-tree/description/
 *
 * algorithms
 * Medium (32.63%)
 * Likes:    817
 * Dislikes: 0
 * Total Accepted:    184.4K
 * Total Submissions: 564.8K
 * Testcase Example:  '[2,1,3]'
 *
 * 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
 *
 * 假设一个二叉搜索树具有如下特征：
 *
 *
 * 节点的左子树只包含小于当前节点的数。
 * 节点的右子树只包含大于当前节点的数。
 * 所有左子树和右子树自身必须也是二叉搜索树。
 *
 *
 * 示例 1:
 *
 * 输入:
 * ⁠   2
 * ⁠  / \
 * ⁠ 1   3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 1   4
 * / \
 * 3   6
 * 输出: false
 * 解释: 输入为: [5,1,4,null,null,3,6]。
 * 根节点的值为 5 ，但是其右子节点值为 4 。
 *
 *
 */
package golang

import (
	"testing"

	"github.com/k0kubun/pp"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "default",
			args: args{root: buildTree([]interface{}{5, 1, 4, nil, nil, 3, 6})},
			want: false,
		},
		{
			name: "default-2",
			args: args{root: buildTree([]interface{}{10, 5, 15, nil, nil, 6, 20})},
			want: false,
		},
		{
			name: "default-3",
			args: args{root: buildTree([]interface{}{3, nil, 30, 10, nil, nil, 15, nil, 45})},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pp.Println(tt.args.root)
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
