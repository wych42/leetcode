/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 *
 * https://leetcode.com/problems/lru-cache/description/
 *
 * algorithms
 * Medium (35.68%)
 * Likes:    8156
 * Dislikes: 334
 * Total Accepted:    736.8K
 * Total Submissions: 2M
 * Testcase Example:  '["LRUCache","put","put","get","put","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]'
 *
 * Design a data structure that follows the constraints of a Least Recently
 * Used (LRU) cache.
 *
 * Implement the LRUCache class:
 *
 *
 * LRUCache(int capacity) Initialize the LRU cache with positive size
 * capacity.
 * int get(int key) Return the value of the key if the key exists, otherwise
 * return -1.
 * void put(int key, int value) Update the value of the key if the key exists.
 * Otherwise, add the key-value pair to the cache. If the number of keys
 * exceeds the capacity from this operation, evict the least recently used
 * key.
 *
 *
 * Follow up:
 * Could you do get and put in O(1) time complexity?
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
 * Output
 * [null, null, null, 1, null, -1, null, -1, 3, 4]
 *
 * Explanation
 * LRUCache lRUCache = new LRUCache(2);
 * lRUCache.put(1, 1); // cache is {1=1}
 * lRUCache.put(2, 2); // cache is {1=1, 2=2}
 * lRUCache.get(1);    // return 1
 * lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
 * lRUCache.get(2);    // returns -1 (not found)
 * lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
 * lRUCache.get(1);    // return -1 (not found)
 * lRUCache.get(3);    // return 3
 * lRUCache.get(4);    // return 4
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= capacity <= 3000
 * 0 <= key <= 3000
 * 0 <= value <= 10^4
 * At most 3 * 10^4 calls will be made to get and put.
 *
 *
*/

// @lc code=start
type Node struct {
	Key  int
	Val  int // 存储的值
	Prev *Node
	Next *Node
}
type LRUCache struct {
	// 记录 capacity
	cap int

	// 需要一个 map 存数据已满足 O(1) 的读/写
	m map[int]*Node

	// 记录最近访问顺序，最近访问的case有
	// - 读一个存在的 key, key 提前
	// - 读一个不存在的 key，key 不变
	// - 写一个 key，key 提前
	// 不考虑并发
	// 用一个双向链表来存储；不能用 array，因为挪动的成本高
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		cap:  capacity,
		m:    make(map[int]*Node, capacity),
		head: nil,
	}
	return c
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		// 需要把 node 提前到 head
		this.moveNodeToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	var (
		node *Node
		ok   bool
	)
	if node, ok = this.m[key]; ok {
		// 需要把 node 提前到 head
		// fmt.Printf("put key=%d exists, update\n", key)
		node.Val = value
		this.moveNodeToHead(node)
	} else {
		// fmt.Printf("put key=%d not exists, add\n", key)
		if len(this.m) >= this.cap {
			// fmt.Printf("cache length=%d exceed cap=%d\n", len(this.m), this.cap)
			this.deleteTail()
			// 再加就超了，需要删除
		}
		node = &Node{Key: key, Val: value}
		// 把 node 挪到头部
		this.addToHead(node)
		this.m[key] = node
	}
}

func (this *LRUCache) moveNodeToHead(node *Node) {
	// this.head..prev..[node]..next..this.tail
	prev := node.Prev
	if prev == nil { // node 就是 head
		return
	}
	node.Prev = nil

	next := node.Next
	if next == nil { // node 是 tail
		this.tail = prev
	}
	// fmt.Printf("moveNodeToHead: node:%v, prev:%v, next:%v\n", node, prev, next)

	// 先把前后两个连起来
	prev.Next = next
	if next != nil {
		next.Prev = prev
	}

	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	// fmt.Printf("moveNodeToHead: current head:%v\n", this.head)
}

func (this *LRUCache) deleteTail() {
	// this.head..prev..tail
	if this.tail == nil {
		// fmt.Printf("delete tail: tail is nil\n")
		return
	}
	// fmt.Printf("delete tail: delete from map,key=%d\n", this.tail.Key)
	delete(this.m, this.tail.Key)
	if this.head == this.tail {
		// 只有一个节点
		this.head = nil
		this.tail = nil
		return
	}
	if this.tail.Prev != nil {
		// 记得把 this.tail 指向前一个
		prev := this.tail.Prev
		prev.Next = nil
		this.tail.Prev = nil
		this.tail = prev
		return
	}
	// fmt.Printf("delete tail: current tail,key=%d\n", this.tail.Key)
}

func (this *LRUCache) addToHead(node *Node) {
	// [node] this.head...
	if this.head == nil {
		// 还没有节点
		this.head = node
		this.tail = node
		return
	}
	node.Next = this.head
	this.head.Prev = node
	this.head = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

