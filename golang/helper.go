package golang

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	root := &TreeNode{Val: nums[0].(int)}
	queue = append(queue, root)
	i := 1
	for i < len(nums) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if nums[i] != nil {
			node.Left = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i = i + 1

		if len(nums) > i && nums[i] != nil {
			node.Right = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i = i + 1
	}
	return root
}
