/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (58.33%)
 * Likes:    5412
 * Dislikes: 351
 * Total Accepted:    869K
 * Total Submissions: 1.5M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Given an integer array nums and an integer k, return the k^th largest
 * element in the array.
 *
 * Note that it is the k^th largest element in the sorted order, not the k^th
 * distinct element.
 *
 *
 * Example 1:
 * Input: nums = [3,2,1,5,6,4], k = 2
 * Output: 5
 * Example 2:
 * Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
 * Output: 4
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	/*
		用一个 minHeap 存 k 个数;
		每次新加数到 heap 时，如果容量满了，就把顶部的去掉（顶部最小）；
		此时 heap 里就是截止当前最大的 K 个数
	*/
	heap := NewMinHeap(k)
	for _, num := range nums {
		heap.Add(num)
	}
	return heap.nums[0]
}

/*
  用 []int{} 存heap
  对于 i，parent= (i-1)/2 left/right children: 2*i+1,2*i+2

  heapify的过程是：
  - 把新的数加到末尾，向上冒泡（是不是比parent小，是的话就替换)
  - 替换顶部的数：向下冒泡，跟左右的比，把左右较小的换上来
*/
type MinHeap struct {
	nums []int
	cap  int
}

func NewMinHeap(cap int) *MinHeap {
	heap := &MinHeap{
		nums: make([]int, 0, cap),
		cap:  cap,
	}
	return heap
}

func (h *MinHeap) Add(num int) {
	if len(h.nums) < h.cap {
		h.nums = append(h.nums, num)
		h.bubbleUp(len(h.nums) - 1)
		return
	}
	// heap 已经满了, 新加的数比顶部的小，不用加进来计算
	if h.nums[0] > num {
		return
	}
	h.nums[0] = num // replace top, as top is min
	h.bubbleDown(0)
}

func (h *MinHeap) bubbleUp(idx int) {
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

func (h *MinHeap) bubbleDown(idx int) {
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

// @lc code=end

