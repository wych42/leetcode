package snippets

import (
	"fmt"
)

/*
    O(n^2),O(n^2),O(n),O(1),稳定
	walk through the slice, compare element next to current one.
	swap them if in wrong order.
	repeat.
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
    O(n^2),O(n^2),O(n^2),O(1),不稳定
	array contains sorted section: nums[:i] and unsorted section nums[i:]
	for each loop(in unsoerted section), find minimal element, swap with first element.
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
  O(n^2),O(n^2),O(n),O(1),稳定
  for each loop, consider [:i] is sorted.
  for each element of [i:], find the pos to insert in [:i] backward.
*/
func insertionSort(nums []int) []int {
	for i := 1; i < len(nums)-1; i++ {
		preIndex := i - 1
		current := nums[i]
		// has prv-element; pre element > current value,compare foward one.
		for preIndex >= 0 && nums[preIndex] > current {
			// move pre-element back for 1 step.(in-place swap)
			nums[preIndex+1] = nums[preIndex]
			preIndex = preIndex - 1
		}
		// pre-element < current, insert current behind pre-element.
		nums[preIndex+1] = current
	}
	return nums
}

/*
  适合中等体量的数据.
  选一个 interval(gap)，把相隔 interval 的数据分为一组，每组排序(粗排)
  再对粗排后的结果，进行插入排序
  重点在于 interval 的选择: Knuth's Formula
*/
func shellSort(nums []int) []int {
	gap := 1
	for gap < len(nums) {
		gap = gap*3 + 1
	}
	for gap > 0 {
		// 从第 gap 个数开始（第一组）
		for i := gap; i < len(nums); i++ {
			j := i
			current := nums[i]
			// j-gap is previous element.
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
	pivot := (len(nums) - 1) / 2 // or random
	nums[pivot], nums[r] = nums[r], nums[pivot]
	for i := range nums {
		if nums[i] < nums[r] { //compare with pivot
			nums[l], nums[i] = nums[i], nums[l] // swap i to most left
			l = l + 1                           // move left forward
		}
	}
	nums[l], nums[r] = nums[r], nums[l] // swap current left with pivot value
	quickSort(nums[:l])
	quickSort(nums[l+1:])
	return nums
}

/*
  用一个 stack 存放每个 partition 的开始结束, for-loop 处理这个 stack
*/
func quickSortIter(nums []int) []int {
	stack := make([]int, 0)
	stack = append(stack, 0, len(nums)-1)

	for len(stack) > 0 {
		temp := stack[len(stack)-2:]
		l, r := temp[0], temp[1]
		stack = stack[:len(stack)-2]
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
