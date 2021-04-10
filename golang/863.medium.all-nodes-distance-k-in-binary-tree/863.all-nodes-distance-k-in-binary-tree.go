/*
 * @lc app=leetcode id=863 lang=golang
 *
 * [863] All Nodes Distance K in Binary Tree
 *
 * https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/description/
 *
 * algorithms
 * Medium (57.82%)
 * Likes:    3422
 * Dislikes: 70
 * Total Accepted:    128.8K
 * Total Submissions: 221.4K
 * Testcase Example:  '[3,5,1,6,2,0,8,null,null,7,4]\n5\n2'
 *
 * We are given a binary tree (with root node root), a target node, and an
 * integer value K.
 *
 * Return a list of the values of all nodes that have a distance K from the
 * target node.  The answer can be returned in any order.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
 *
 * Output: [7,4,1]
 *
 * Explanation:
 * The nodes that are a distance 2 from the target node (with value 5)
 * have values 7, 4, and 1.
 *
 *
 *
 * Note that the inputs "root" and "target" are actually TreeNodes.
 * The descriptions of the inputs above are just serializations of these
 * objects.
 *
 *
 *
 *
 * Note:
 *
 *
 * The given tree is non-empty.
 * Each node in the tree has unique values 0 <= node.val <= 500.
 * The target node is a node in the tree.
 * 0 <= K <= 1000.
 *
 *
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	/*
		先前序遍历数，直到 target 位置，中间记录下来子节点的parent
		然后从 target 开始，分别访问它的左、右、父节点，距离+1；直到距离k位置
	*/

	if k == 0 {
		return []int{target.Val}
	}
	ans := make([]int, 0)

	m := make(map[*TreeNode]*TreeNode, 0)
	traversal(root, target, &m)

	visited := make(map[*TreeNode]bool)

	queue := make([]*TreeNode, 0)
	dis := 0
	queue = append(queue, target)
	visited[target] = true
	layerLen := len(queue)
	for len(queue) > 0 {
		if layerLen == 0 {
			layerLen = len(queue)
			dis++
		}

		if dis == k {
			for _, n := range queue {
				ans = append(ans, n.Val)
			}
			return ans
		}
		node := queue[0]
		queue = queue[1:]
		if node.Left != nil && !visited[node.Left] {
			queue = append(queue, node.Left)
			visited[node.Left] = true
		}
		if node.Right != nil && !visited[node.Right] {
			queue = append(queue, node.Right)
			visited[node.Right] = true
		}
		if parent, ok := m[node]; ok && !visited[parent] {
			fmt.Println("xx", parent.Val, visited[parent])
			queue = append(queue, parent)
			visited[parent] = true
		}
		/*
				for _, n := range queue {
					fmt.Print(n.Val)
				}
		        fmt.Print("\n")
				fmt.Println("layerLen", layerLen)
		*/
		layerLen--

	}
	return ans

}

func traversal(root, target *TreeNode, m *map[*TreeNode]*TreeNode) {
	// 问题是会多记录
	if root == target { // 找到 target 了
		return
	}
	if root.Left != nil {
		(*m)[root.Left] = root
		traversal(root.Left, target, m)
	}
	if root.Right != nil {
		(*m)[root.Right] = root
		traversal(root.Right, target, m)
	}
}

// @lc code=end

