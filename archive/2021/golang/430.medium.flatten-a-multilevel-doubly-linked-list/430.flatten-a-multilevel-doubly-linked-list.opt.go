/*
 * @lc app=leetcode id=430 lang=golang
 *
 * [430] Flatten a Multilevel Doubly Linked List
 *
 * https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list/description/
 *
 * algorithms
 * Medium (56.78%)
 * Likes:    2155
 * Dislikes: 187
 * Total Accepted:    153.3K
 * Total Submissions: 269.4K
 * Testcase Example:  '[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]'
 *
 * You are given a doubly linked list which in addition to the next and
 * previous pointers, it could have a child pointer, which may or may not point
 * to a separate doubly linked list. These child lists may have one or more
 * children of their own, and so on, to produce a multilevel data structure, as
 * shown in the example below.
 *
 * Flatten the list so that all the nodes appear in a single-level, doubly
 * linked list. You are given the head of the first level of the list.
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
 * Output: [1,2,3,7,8,11,12,9,10,4,5,6]
 * Explanation:
 *
 * The multilevel linked list in the input is as follows:
 *
 *
 *
 * After flattening the multilevel linked list it becomes:
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: head = [1,2,null,3]
 * Output: [1,3,2]
 * Explanation:
 *
 * The input multilevel linked list is as follows:
 *
 * ⁠ 1---2---NULL
 * ⁠ |
 * ⁠ 3---NULL
 *
 *
 * Example 3:
 *
 *
 * Input: head = []
 * Output: []
 *
 *
 *
 *
 * How multilevel linked list is represented in test case:
 *
 * We use the multilevel linked list from Example 1 above:
 *
 *
 * ⁠1---2---3---4---5---6--NULL
 * ⁠        |
 * ⁠        7---8---9---10--NULL
 * ⁠            |
 * ⁠            11--12--NULL
 *
 * The serialization of each level is as follows:
 *
 *
 * [1,2,3,4,5,6,null]
 * [7,8,9,10,null]
 * [11,12,null]
 *
 *
 * To serialize all levels together we will add nulls in each level to signify
 * no node connects to the upper node of the previous level. The serialization
 * becomes:
 *
 *
 * [1,2,3,4,5,6,null]
 * [null,null,7,8,9,10,null]
 * [null,11,12,null]
 *
 *
 * Merging the serialization of each level and removing trailing nulls we
 * obtain:
 *
 *
 * [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
 *
 *
 * Constraints:
 *
 *
 * The number of Nodes will not exceed 1000.
 * 1 <= Node.val <= 10^5
 *
 *
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	/*
		一次挪一个节点;
		如果发现 curr.Child != nil, 就找到 Child 的尾部，与 curr.Next 相连;
		curr.Child 变成 curr.Next, 并且把 curr.Child 设置为 nil
	*/
	if root == nil {
		return nil
	}
	curr := root
	for curr != nil {
		if curr.Child == nil {
			curr = curr.Next // move
			continue
		}

		child := curr.Child
		for child.Next != nil {
			child = child.Next // find tail of child
		}
		// 1 - 2 - 3
		//     |
		//     4 - 5
		//         |
		//         6
		// to
		// 1 - 2    3
		//     |	|
		//     4  - 5
		//          |
		//			6
		child.Next = curr.Next
		if curr.Next != nil {
			curr.Next.Prev = child
		}

		// to
		// 1 - 2 - 4 - 5 - 3
		//             |
		//             6
		curr.Next = curr.Child
		curr.Child.Prev = curr
		curr.Child = nil
	}
	return head
}

// @lc code=end

