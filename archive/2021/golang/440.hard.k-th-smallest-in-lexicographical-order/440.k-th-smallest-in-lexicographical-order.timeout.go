/*
 * @lc app=leetcode id=440 lang=golang
 *
 * [440] K-th Smallest in Lexicographical Order
 *
 * https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/description/
 *
 * algorithms
 * Hard (29.79%)
 * Likes:    415
 * Dislikes: 60
 * Total Accepted:    15K
 * Total Submissions: 50.2K
 * Testcase Example:  '13\n2'
 *
 * Given integers n and k, find the lexicographically k-th smallest integer in
 * the range from 1 to n.
 *
 * Note: 1 ≤ k ≤ n ≤ 10^9.
 *
 * Example:
 *
 * Input:
 * n: 13   k: 2
 *
 * Output:
 * 10
 *
 * Explanation:
 * The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so
 * the second smallest number is 10.
 *
 *
 *
 */

// @lc code=start
/*
找最小用 maxHeap
*/
import "strconv"

func findKthNumber(n int, k int) int {
	heap := NewMaxHeap(k)
	for i := 1; i <= n; i++ {
		heap.Add(strconv.Itoa(i))
	}
	ans, _ := strconv.Atoi(heap.nums[0])
	return ans
}

/*
  用 []string{} 存heap
  对于 i，parent= (i-1)/2 left/right children: 2*i+1,2*i+2

  heapify的过程是：
  - 把新的数加到末尾，向上冒泡（是不是比parent大，是的话就替换)
  - 替换顶部的数：向下冒泡，跟左右的比，把左右较大的换上来
*/
type MaxHeap struct {
	nums []string
	cap  int
}

func NewMaxHeap(cap int) *MaxHeap {
	heap := &MaxHeap{
		nums: make([]string, 0, cap),
		cap:  cap,
	}
	return heap
}

func (h *MaxHeap) Add(num string) {
	if len(h.nums) < h.cap {
		h.nums = append(h.nums, num)
		h.bubbleUp(len(h.nums) - 1)
		return
	}
	// heap 已经满了, 新加的数比顶部的大，不用加进来计算
	if h.nums[0] <= num {
		return
	}
	h.nums[0] = num // replace top, as top is min
	h.bubbleDown(0)
}

func (h *MaxHeap) bubbleUp(idx int) {
	// 当前数与parent比较，如果小就交换
	parent := (idx - 1) / 2
	for parent >= 0 {
		if h.nums[parent] >= h.nums[idx] { // parent 比自己大，是个合格的 heap
			return
		}
		h.nums[parent], h.nums[idx] = h.nums[idx], h.nums[parent]
		idx = parent
		parent = (idx - 1) / 2
	}
}

func (h *MaxHeap) bubbleDown(idx int) {
	// 当前数与左右子节点比较
	// 找到左右子节点里大的那个，换上来
	l := idx*2 + 1
	r := idx*2 + 2
	maxIdx := idx
	if l < len(h.nums) && h.nums[l] > h.nums[maxIdx] {
		maxIdx = l
	}
	if r < len(h.nums) && h.nums[r] > h.nums[maxIdx] {
		maxIdx = r
	}
	if maxIdx == idx { // idx 比左右都小，就放在这里
		return
	}
	// 否则就交换
	h.nums[idx], h.nums[maxIdx] = h.nums[maxIdx], h.nums[idx]
	h.bubbleDown(maxIdx)
}

// @lc code=end

