package snippets

/*
heap 是个完全二叉树;
MaxHeap 的最大元素在堆顶;
新加一个数，加到末尾，往上冒泡到合适的位置;
删除堆顶，把末尾的数放到堆顶，往下冒泡到合适的位置

对于一个节点:
parent: (i-1)/2
left child: 2*i+1
right child 2*i+2
*/
type MaxHeap struct {
	nums []int
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

/*
Remove: 删除堆顶时，把末尾放到堆顶，然后向下冒泡：
- 把当前节点小于左右子节点，与大的那个子节点交换
- 直到大于子节点或者达到尾部
*/
func (h *MaxHeap) Remove() {
	if len(h.nums) == 0 { // nothing to remove
		return
	}
	// 用末尾替代头部
	h.nums[0] = h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(nums)-1] // 缩减
	// 向下冒泡
	h.bubbleDown(0)

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

func (h *MaxHeap) Peak() int {
	if len(h.nums) == 0 {
		return -1
	}
	return h.nums[0]

}

func (h *MaxHeap) Build(nums []int) {
	for _, num := range nums {
		h.Insert(num)
	}
}
