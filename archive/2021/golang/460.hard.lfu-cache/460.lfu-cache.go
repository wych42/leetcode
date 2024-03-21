/*
 * @lc app=leetcode id=460 lang=golang
 *
 * [460] LFU Cache
 *
 * https://leetcode.com/problems/lfu-cache/description/
 *
 * algorithms
 * Hard (36.01%)
 * Likes:    1927
 * Dislikes: 150
 * Total Accepted:    103.1K
 * Total Submissions: 283.4K
 * Testcase Example:  '["LFUCache","put","put","get","put","get","get","put","get","get","get"]\n' +
  '[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]'
 *
 * Design and implement a data structure for a Least Frequently Used (LFU)
 * cache.
 *
 * Implement the LFUCache class:
 *
 *
 * LFUCache(int capacity) Initializes the object with the capacity of the data
 * structure.
 * int get(int key) Gets the value of the key if the key exists in the cache.
 * Otherwise, returns -1.
 * void put(int key, int value) Update the value of the key if present, or
 * inserts the key if not already present. When the cache reaches its capacity,
 * it should invalidate and remove the least frequently used key before
 * inserting a new item. For this problem, when there is a tie (i.e., two or
 * more keys with the same frequency), the least recently used key would be
 * invalidated.
 *
 *
 * To determine the least frequently used key, a use counter is maintained for
 * each key in the cache. The key with the smallest use counter is the least
 * frequently used key.
 *
 * When a key is first inserted into the cache, its use counter is set to 1
 * (due to the put operation). The use counter for a key in the cache is
 * incremented either a get or put operation is called on it.
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get",
 * "get"]
 * [[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
 * Output
 * [null, null, null, 1, null, -1, 3, null, -1, 3, 4]
 *
 * Explanation
 * // cnt(x) = the use counter for key x
 * // cache=[] will show the last used order for tiebreakers (leftmost element
 * is  most recent)
 * LFUCache lfu = new LFUCache(2);
 * lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
 * lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
 * lfu.get(1);      // return 1
 * ⁠                // cache=[1,2], cnt(2)=1, cnt(1)=2
 * lfu.put(3, 3);   // 2 is the LFU key because cnt(2)=1 is the smallest,
 * invalidate 2.
 * // cache=[3,1], cnt(3)=1, cnt(1)=2
 * lfu.get(2);      // return -1 (not found)
 * lfu.get(3);      // return 3
 * ⁠                // cache=[3,1], cnt(3)=2, cnt(1)=2
 * lfu.put(4, 4);   // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate
 * 1.
 * ⁠                // cache=[4,3], cnt(4)=1, cnt(3)=2
 * lfu.get(1);      // return -1 (not found)
 * lfu.get(3);      // return 3
 * ⁠                // cache=[3,4], cnt(4)=1, cnt(3)=3
 * lfu.get(4);      // return 4
 * ⁠                // cache=[3,4], cnt(4)=2, cnt(3)=3
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= capacity, key, value <= 10^4
 * At most 10^5 calls will be made to get and put.
 *
 *
 *
 * Follow up: Could you do both operations in O(1) time complexity?
*/

// @lc start
type Node struct {
	Key  int
	Val  int // 存储的值
	Prev *Node
	Next *Node
	Freq int
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (list *DoubleLinkedList) Insert(node *Node) {
	node.Next = list.head
	if list.head != nil {
		list.head.Prev = node
	}
	list.head = node
	if list.tail == nil {
		list.tail = node
	}
	list.size++
}

func (list *DoubleLinkedList) Remove(node *Node) {
	if node == nil {
		return
	}
	prev := node.Prev
	next := node.Next
	if prev == nil { // node 是 head
		list.head = next
	} else {
		prev.Next = next
	}
	if next == nil { // node 是 tail
		list.tail = prev
	} else {
		next.Prev = prev
	}
	list.size--
}

func (list *DoubleLinkedList) Size() int {
	return list.size
}

type LFUCache struct {
	// 记录 capacity
	cap int

	// nodeMap 需要一个 map 存数据已满足 O(1) 的读/写
	nm map[int]*Node
	/*
	   把每一个频率都维护成一个链表，则对于每一种频率，都是 LRU 问题
	   3 a, b, c
	   2 d, f
	   1 f
	   case:
	   - 读一个 key，根据 node.Freq 找到当前频率链表，删除；加到新的频率链表
	   - 加一个新key: 加到 Freq=1 的频率链表
	   - 更新一个key: 从当前频率链表删除，加到 Freq=1 的频率链表
	   - 删除过期的：用 minFreq 变量记录当前最小频率，需要删除时从这里删除
	*/

	// freq map
	fm map[int]*DoubleLinkedList

	minFreq int
}

func Constructor(capacity int) LFUCache {
	c := LFUCache{
		cap: capacity,
		nm:  make(map[int]*Node, capacity),
		fm:  make(map[int]*DoubleLinkedList),
	}
	return c
}

func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}
	node, ok := this.nm[key]
	if !ok {
		return -1
	}
	// 当前频率
	freq := node.Freq

	// 从当前频率的链表里删除节点
	fl := this.fm[freq]
	fl.Remove(node)
	// 如果当前频率的链表空了，从频率map里删除
	if fl.Size() == 0 {
		delete(this.fm, freq)
		// 更新 minFreq
		if freq == this.minFreq {
			this.minFreq++
		}
	}

	// 新节点
	node = &Node{Key: node.Key, Val: node.Val, Freq: freq + 1}
	this.addNode(node)
	return node.Val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	node, ok := this.nm[key]
	//fmt.Println("put: try get key:", key, node)
	if !ok {
		// 不存在，要加新的节点
		if len(this.nm) >= this.cap {

			// 满了，先删除
			// 拿到 minFreq 对应的链表
			list := this.fm[this.minFreq]
			// 找到要删除的节点
			node := list.tail
			//fmt.Printf("put: cache is full, delete freq=%d, node=%v", this.minFreq, node)
			// 从 map 里删除
			delete(this.nm, node.Key)
			// 从 list 里移除
			list.Remove(node)

			// 如果链表空了，把链表删除
			if list.Size() == 0 {
				delete(this.fm, this.minFreq)
			}
		}
		this.minFreq = 1
		// 新节点的 freq = 1
		node = &Node{Key: key, Val: value, Freq: 1}
		this.nm[key] = node
		nfl, ok := this.fm[1]
		if !ok {
			nfl = &DoubleLinkedList{}
		}
		nfl.Insert(node)
		this.fm[1] = nfl
		return
	} else {
		// 更新旧节点
		// 当前频率
		//fmt.Printf("put: key exists, updating node=%v", node)
		freq := node.Freq

		// 从当前频率的链表里删除节点
		fl := this.fm[freq]
		fl.Remove(node)
		// 如果当前频率的链表空了，从map里删除
		if fl.Size() == 0 {
			delete(this.fm, freq)
			if freq == this.minFreq {
				this.minFreq++
			}
		}

		// 新节点
		node = &Node{Key: node.Key, Val: value, Freq: freq + 1}
		this.addNode(node)
	}
}

func (this *LFUCache) addNode(node *Node) {
	this.nm[node.Key] = node

	// 新节点要加入到频率1的链表里
	nfl, ok := this.fm[node.Freq]
	if !ok {
		nfl = &DoubleLinkedList{}
	}
	nfl.Insert(node)
	if !ok {
		this.fm[node.Freq] = nfl
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

