package snippets

import (
	"fmt"
)

/*
	冒泡排序
    O(n^2),O(n^2),O(n),O(1),稳定
	挨个比较相邻的两个元素，如果顺序错了就交换。循环再来一遍。
	每次循环，都找出了最大的数,放到后面去。所以叫冒泡。
*/
func bubbleSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

/*
	选择排序
    O(n^2),O(n^2),O(n^2),O(1),不稳定
	把列表分为排序和未排序两部分: nums[:i] 和 nums[i:]
	每次循环，都从为排序里面，找一个最小的，放到为排序的最前面。并把已排序的往后挪一位。
*/
func selectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

/*
  插入排序
  O(n^2),O(n^2),O(n),O(1),稳定
  列表的 [:i] 是已经排序的。
  每次循环，都从剩下的里面拿一个元素，找插入位置
  for each loop, consider [:i] is sorted.
  for each element of [i:], find the pos to insert in [:i] backward.
*/
func insertionSort(nums []int) []int {
	for i := 1; i < len(nums)-1; i++ {
		preIndex := i - 1
		current := nums[i]
		// nums[:i] 前面还有数，并且前面的数大于当前数；继续往前比较
		for preIndex >= 0 && nums[preIndex] > current {
			// 并且，把 preIndex 往后挪
			nums[preIndex+1] = nums[preIndex]
			preIndex = preIndex - 1
		}
		// 否则，就把当前值放在 preIndex 后面
		nums[preIndex+1] = current
	}
	return nums
}

/*
  希尔排序: 适合中等体量的数据.
  选一个 interval/gap，每隔 gap 的数分为一组，每组排序(粗排)
  再对粗排后的结果，更新（缩小 gap），再来一次，直到 gap 为 1，就是插入排序
  重点在于 interval 的选择: Knuth's Formula
*/
func shellSort(nums []int) []int {
	gap := 1
	for gap < len(nums) {
		gap = gap*3 + 1
	}
	for gap > 0 {
		// 从第 gap 个数开始（第一组）; 一开始 gap 的值最大
		for i := gap; i < len(nums); i++ {
			j := i
			current := nums[i]
			// j-gap 是这一组前一个数的下标
			// 如果 前一个数比当前数大，就交换
			for j-gap >= 0 && nums[j-gap] > current {
				nums[j] = nums[j-gap]
				j = j - gap
			}
			nums[j] = current
		}
		gap = gap / 3
	}
	return nums
}

/*
  O(nlogn),O(nlogn),O(nlogn),O(n),稳定
  2-路归并：把两个有序子序列合并；
  先把序列分成两个，分别对两个序列做, 各自做拆分-归并排序。（递归）
*/
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	return merge(
		mergeSort(nums[:len(nums)/2]),
		mergeSort(nums[len(nums)/2:]),
	)
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	lp, rp := 0, 0
	for lp < len(left) && rp < len(right) {
		if left[lp] <= right[rp] {
			result = append(result, left[lp])
			lp = lp + 1
		} else {
			result = append(result, right[rp])
			rp = rp + 1
		}
	}
	if lp < len(left) {
		result = append(result, left[lp:]...)
	}
	if rp < len(right) {
		result = append(result, right[rp:]...)
	}
	return result
}

/*
  O(nlogn),O(n^2),O(nlogn),O(nlogn),不稳定
  找到一个基准数；小的数都在左边，大的都在右边。递归。
*/
func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	l, r := 0, len(nums)-1
	pivot := (len(nums) - 1) / 2                // 找中间的或者随机一个
	nums[pivot], nums[r] = nums[r], nums[pivot] // 把基准数放到最后
	for i := range nums {
		if nums[i] < nums[r] { // 跟基准数比,如果比基准数小，就放到最左边
			nums[l], nums[i] = nums[i], nums[l]
			l = l + 1 // 并且把最左的坐标右移一个
		}
	}
	nums[l], nums[r] = nums[r], nums[l] // 把当前的最左坐标值，跟最右坐标值换一下（把基准数放到中间）
	quickSort(nums[:l])                 // 左边是都比基准数小的
	quickSort(nums[l+1:])               // 右边是大的
	return nums
}

/*
  用一个 stack 存放每个 partition 的开始结束, for-loop 处理这个 stack
*/
func quickSortIter(nums []int) []int {
	stack := make([]int, 0)
	// 放入第一个分片的下标
	stack = append(stack, 0, len(nums)-1)

	for len(stack) > 0 {
		temp := stack[len(stack)-2:]
		// 出栈，拿到这次要处理的
		l, r := temp[0], temp[1]
		stack = stack[:len(stack)-2]
		// 执行分片逻辑, 得到基准数的下标
		pivot := partition(nums, l, r)
		if pivot-1 > l { // elements in left side
			stack = append(stack, l, pivot-1)
		}
		if pivot+1 < r { // elements in right side
			stack = append(stack, pivot+1, r)
		}
	}
	return nums
}

// 分片，跟递归的分片流程一样
func partition(nums []int, l, r int) int {
	pivot := l + (r-l)/2
	nums[pivot], nums[r] = nums[r], nums[pivot]
	for i := l; i < r; i++ {
		if nums[i] < nums[r] {
			nums[i], nums[l] = nums[l], nums[i]
			l = l + 1
		}
	}
	nums[l], nums[r] = nums[r], nums[l] // swap current left with pivot value
	return l
}

/*
  构造一个小顶堆；将第一个元素放到末尾；无序 [:n-1], 有序 [n];
  对 [:n-1] 再次构造小顶堆，将第一个元素放到末尾：无序 [:i], 有序 [i:]
  直到 i=0;
*/
func heapSort(nums []int) []int {
	h := buildMinHeap(nums)
	fmt.Println(h.nums)
	for i := len(h.nums) - 1; i > 0; i-- {
		h.nums[0], h.nums[i] = h.nums[i], h.nums[0]
		h.size = h.size - 1
		h.maxHeapify(0)
	}
	return h.nums
}

type heap struct {
	nums []int
	size int
}

func (h *heap) maxHeapify(current int) {
	left, right := 2*current+1, 2*current+2
	max := current
	if left < h.size && h.nums[left] > h.nums[max] {
		max = left
	}
	if right < h.size && h.nums[right] > h.nums[max] {
		max = right
	}
	if max != current {
		h.nums[current], h.nums[max] = h.nums[max], h.nums[current]
		h.maxHeapify(max)
	}
}

func buildMinHeap(nums []int) heap {
	h := heap{nums: nums, size: len(nums)}
	for i := len(nums) / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
	return h
}
