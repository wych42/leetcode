/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 *
 * https://leetcode.com/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (64.55%)
 * Likes:    785
 * Dislikes: 364
 * Total Accepted:    138.6K
 * Total Submissions: 214.4K
 * Testcase Example:  '[5,2,3,1]'
 *
 * Given an array of integers nums, sort the array in ascending order.
 *
 *
 * Example 1:
 * Input: nums = [5,2,3,1]
 * Output: [1,2,3,5]
 * Example 2:
 * Input: nums = [5,1,1,2,0,0]
 * Output: [0,0,1,1,2,5]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 50000
 * -50000 <= nums[i] <= 50000
 *
 *
 */

// @lc code=start
func sortArray(nums []int) []int {
	h := NewMaxHeap()
	h.Build(nums)
	//fmt.Println(h.nums, h.size)
	for i := len(nums) - 1; i > 0; i-- {
		// heap 顶部是最大的元素，把顶部元素与末尾交换,末尾就是有序的了
		// h.size-1,再进行向下冒泡，把刚刚放到顶部的元素回答正确的位置
		h.nums[0], h.nums[i] = h.nums[i], h.nums[0]
		h.size--
		h.bubbleDown(0)
	}
	return h.nums
}

type MaxHeap struct {
	nums []int
	size int
}

func NewMaxHeap() MaxHeap {
	return MaxHeap{}
}

/*
Insert: 加数的时候，加到末尾，然后向上冒泡：
- 把当前节点与父节点比较，如果大，就交换
- 直到小于父节点或者达到头部
*/
func (h *MaxHeap) Insert(num int) {
	h.nums = append(h.nums, num)
	h.size++
	idx := len(h.nums) - 1
	parent := (idx - 1) / 2
	for parent >= 0 {
		if h.nums[parent] >= num { // parent > num, 是 max heap
			break
		}
		// parent < num; 交换
		h.nums[parent], h.nums[idx] = h.nums[idx], h.nums[parent]
		idx = parent
		parent = (idx - 1) / 2
	}
}

func (h *MaxHeap) bubbleDown(idx int) {
	left, right := 2*idx+1, 2*idx+2
	// find max from left/right child
	maxIdx := idx
	if left < h.size && h.nums[left] > h.nums[maxIdx] {
		maxIdx = left
	}
	if right < h.size && h.nums[right] > h.nums[maxIdx] {
		maxIdx = right
	}
	if maxIdx == idx { // 当前节点大于左右子节点，不用动
		return
	}
	// 交换
	h.nums[idx], h.nums[maxIdx] = h.nums[maxIdx], h.nums[idx]
	h.bubbleDown(maxIdx) // 继续进行
}

func (h *MaxHeap) Build(nums []int) {
	for _, num := range nums {
		h.Insert(num)
	}
}

// @lc code=end

