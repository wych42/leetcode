import "sort"
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Peek() int {
	return (*h)[0]
}
func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	/*
		用贪心的思路：越早开始的会议需要优先安排.
		1. 按照开始时间排序
		2. 记录第一个会议的结束时间(即该会议室的最早可用时间)；
			- 如果下一个会议的开始时间早于结束时间，就冲突了，需要安排新的会议室
			- 如果下一个会议的开始时间晚于结束时间，就不用安排新的会议室，但是需要更新会议室的最早可用时间

		3. 每新开一个会议室，都要追踪该会议室的最早可用时间
		4. 当有新的会议需要安排时，从追踪列表里找一个最早可用的，如果找不到，就安排新的
	*/
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	h := &IntHeap{}
	for _, interval := range intervals {
		if h.Len() > 0 && h.Peek() <= interval[0] {
			heap.Pop(h)
		}
		heap.Push(h, interval[1])
	}
	return h.Len()
}