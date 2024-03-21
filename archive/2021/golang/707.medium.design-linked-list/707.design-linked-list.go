/*
 * @lc app=leetcode id=707 lang=golang
 *
 * [707] Design Linked List
 *
 * https://leetcode.com/problems/design-linked-list/description/
 *
 * algorithms
 * Medium (25.86%)
 * Likes:    809
 * Dislikes: 858
 * Total Accepted:    104.8K
 * Total Submissions: 403.1K
 * Testcase Example:  '["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]\n' +
  '[[],[1],[3],[1,2],[1],[1],[1]]'
 *
 * Design your implementation of the linked list. You can choose to use a
 * singly or doubly linked list.
 * A node in a singly linked list should have two attributes: val and next. val
 * is the value of the current node, and next is a pointer/reference to the
 * next node.
 * If you want to use the doubly linked list, you will need one more attribute
 * prev to indicate the previous node in the linked list. Assume all nodes in
 * the linked list are 0-indexed.
 *
 * Implement the MyLinkedList class:
 *
 *
 * MyLinkedList() Initializes the MyLinkedList object.
 * int get(int index) Get the value of the index^th node in the linked list. If
 * the index is invalid, return -1.
 * void addAtHead(int val) Add a node of value val before the first element of
 * the linked list. After the insertion, the new node will be the first node of
 * the linked list.
 * void addAtTail(int val) Append a node of value val as the last element of
 * the linked list.
 * void addAtIndex(int index, int val) Add a node of value val before the
 * index^th node in the linked list. If index equals the length of the linked
 * list, the node will be appended to the end of the linked list. If index is
 * greater than the length, the node will not be inserted.
 * void deleteAtIndex(int index) Delete the index^th node in the linked list,
 * if the index is valid.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get",
 * "deleteAtIndex", "get"]
 * [[], [1], [3], [1, 2], [1], [1], [1]]
 * Output
 * [null, null, null, null, 2, null, 3]
 *
 * Explanation
 * MyLinkedList myLinkedList = new MyLinkedList();
 * myLinkedList.addAtHead(1);
 * myLinkedList.addAtTail(3);
 * myLinkedList.addAtIndex(1, 2);    // linked list becomes 1->2->3
 * myLinkedList.get(1);              // return 2
 * myLinkedList.deleteAtIndex(1);    // now the linked list is 1->3
 * myLinkedList.get(1);              // return 3
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= index, val <= 1000
 * Please do not use the built-in LinkedList library.
 * At most 2000 calls will be made to get, addAtHead, addAtTail, addAtIndex and
 * deleteAtIndex.
 *
 *
*/

// @lc code=start
type Node struct {
	val  int
	next *Node
}
type MyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) count() {
	curr := this.head
	c := 1
	for curr.next != nil {
		c++
		curr = curr.next
	}
	fmt.Println("total", c, "tail value", curr.val, curr.next, this.tail)
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	// fmt.Println("get", index, this.length)

	if index > this.length-1 {
		return -1
	}

	curr := this.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	return curr.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	n := &Node{val: val}
	if this.length == 0 {
		this.head = n
		this.tail = n
	} else {
		n.next = this.head
		this.head = n
	}
	this.length++
	//     fmt.Println("add at head", val, this.length)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	n := &Node{val: val}
	if this.length == 0 {
		this.head = n
		this.tail = n
	} else {

		this.tail.next = n
		this.tail = this.tail.next
	}
	this.length++
	// fmt.Println("add at tail", val , this.length, this.tail.val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.length {
		return
	}

	// index 和 length 相等，加在最好
	if index == this.length {
		this.AddAtTail(val)
		return
	}

	dummy := &Node{next: this.head}
	prev, curr := dummy, this.head
	for i := 0; i < index; i++ {
		prev = prev.next
		curr = curr.next
	}
	n := &Node{val: val, next: curr}
	prev.next = n
	// 注意重置 head
	this.head = dummy.next
	this.length++
	// fmt.Println("add at index",index, val, this.length)
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index > this.length-1 {
		return
	}
	dummy := &Node{next: this.head}
	prev := dummy // 要删除的是 prev.next
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	if prev.next != nil {
		prev.next = prev.next.next
	}
	// 如果删除的是最后一个，要重置 tail
	if index == this.length-1 {
		this.tail = prev
	}

	// 注意重置 head
	this.head = dummy.next
	this.length--
	//    fmt.Println("delete at index", index, this.length)

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end

