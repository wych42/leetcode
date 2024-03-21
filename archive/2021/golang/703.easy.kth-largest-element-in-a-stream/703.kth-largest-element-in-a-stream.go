/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 *
 * https://leetcode.com/problems/kth-largest-element-in-a-stream/description/
 *
 * algorithms
 * Easy (50.66%)
 * Likes:    1174
 * Dislikes: 716
 * Total Accepted:    127.8K
 * Total Submissions: 251K
 * Testcase Example:  '["KthLargest","add","add","add","add","add"]\n' +
  '[[3,[4,5,8,2]],[3],[5],[10],[9],[4]]'
 *
 * Design a class to find the k^th largest element in a stream. Note that it is
 * the k^th largest element in the sorted order, not the k^th distinct
 * element.
 *
 * Implement KthLargest class:
 *
 *
 * KthLargest(int k, int[] nums) Initializes the object with the integer k and
 * the stream of integers nums.
 * int add(int val) Returns the element representing the k^th largest element
 * in the stream.
 *
 *
 *
 * Example 1:
 *
 *
 * Input
 * ["KthLargest", "add", "add", "add", "add", "add"]
 * [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
 * Output
 * [null, 4, 5, 5, 8, 8]
 *
 * Explanation
 * KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
 * kthLargest.add(3);   // return 4
 * kthLargest.add(5);   // return 5
 * kthLargest.add(10);  // return 5
 * kthLargest.add(9);   // return 8
 * kthLargest.add(4);   // return 8
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= 10^4
 * 0 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * -10^4 <= val <= 10^4
 * At most 10^4 calls will be made to add.
 * It is guaranteed that there will be at least k elements in the array when
 * you search for the k^th element.
 *
 *
*/

// @lc code=start
type KthLargest struct {
	/*
	   这就是个 minHeap
	*/
	nums []int
	cap  int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		nums: make([]int, 0, k),
		cap:  k,
	}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (h *KthLargest) Add(num int) int {
	if len(h.nums) < h.cap {
		h.nums = append(h.nums, num)
		h.bubbleUp(len(h.nums) - 1)
		return h.nums[0]
	}
	// heap 已经满了, 新加的数比顶部的小，不用加进来计算
	if h.nums[0] > num {
		return h.nums[0]
	}
	h.nums[0] = num // replace top, as top is min
	h.bubbleDown(0)
	return h.nums[0]
}

func (h *KthLargest) bubbleUp(idx int) {
	// 当前数与parent比较，如果小就交换
	parent := (idx - 1) / 2
	for parent >= 0 {
		if h.nums[parent] <= h.nums[idx] { // parent 比自己小，是个合格的 heap
			return
		}
		h.nums[parent], h.nums[idx] = h.nums[idx], h.nums[parent]
		idx = parent
		parent = (idx - 1) / 2
	}
}

func (h *KthLargest) bubbleDown(idx int) {
	// 当前数与左右子节点比较
	// 找到左右子节点里小的那个，换上来
	l := idx*2 + 1
	r := idx*2 + 2
	minIdx := idx
	if l < len(h.nums) && h.nums[l] < h.nums[minIdx] {
		minIdx = l
	}
	if r < len(h.nums) && h.nums[r] < h.nums[minIdx] {
		minIdx = r
	}
	if minIdx == idx { // idx 比左右都小，就放在这里
		return
	}
	// 否则就交换
	h.nums[idx], h.nums[minIdx] = h.nums[minIdx], h.nums[idx]
	h.bubbleDown(minIdx)
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

