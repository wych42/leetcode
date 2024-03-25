/*
 * @lc app=leetcode.cn id=218 lang=golang
 *
 * [218] The Skyline Problem
 *
 * https://leetcode.cn/problems/the-skyline-problem/description/
 *
 * algorithms
 * Hard (55.24%)
 * Likes:    823
 * Dislikes: 0
 * Total Accepted:    53.4K
 * Total Submissions: 96.7K
 * Testcase Example:  '[[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]'
 *
 * A city's skyline is the outer contour of the silhouette formed by all the
 * buildings in that city when viewed from a distance. Given the locations and
 * heights of all the buildings, return the skyline formed by these buildings
 * collectively.
 *
 * The geometric information of each building is given in the array buildings
 * where buildings[i] = [lefti, righti, heighti]:
 *
 *
 * lefti is the x coordinate of the left edge of the i^th building.
 * righti is the x coordinate of the right edge of the i^th building.
 * heighti is the height of the i^th building.
 *
 *
 * You may assume all buildings are perfect rectangles grounded on an
 * absolutely flat surface at height 0.
 *
 * The skyline should be represented as a list of "key points" sorted by their
 * x-coordinate in the form [[x1,y1],[x2,y2],...]. Each key point is the left
 * endpoint of some horizontal segment in the skyline except the last point in
 * the list, which always has a y-coordinate 0 and is used to mark the
 * skyline's termination where the rightmost building ends. Any ground between
 * the leftmost and rightmost buildings should be part of the skyline's
 * contour.
 *
 * Note: There must be no consecutive horizontal lines of equal height in the
 * output skyline. For instance, [...,[2 3],[4 5],[7 5],[11 5],[12 7],...] is
 * not acceptable; the three lines of height 5 should be merged into one in the
 * final output as such: [...,[2 3],[4 5],[12 7],...]
 *
 *
 * Example 1:
 *
 *
 * Input: buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
 * Output: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]
 * Explanation:
 * Figure A shows the buildings of the input.
 * Figure B shows the skyline formed by those buildings. The red points in
 * figure B represent the key points in the output list.
 *
 *
 * Example 2:
 *
 *
 * Input: buildings = [[0,2,3],[2,5,3]]
 * Output: [[0,3],[5,0]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= buildings.length <= 10^4
 * 0 <= lefti < righti <= 2^31 - 1
 * 1 <= heighti <= 2^31 - 1
 * buildings is sorted by lefti in non-decreasing order.

 check this: https://www.youtube.com/watch?v=GSBLe8cKu0s
 *
 *
*/

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// @lc code=start
func getSkyline(buildings [][]int) [][]int {
	points := make([]*Point, len(buildings)*2)
	for i, building := range buildings {
		points[2*i] = &Point{x: building[0], height: building[2], typ: PointType_Start, index: i}
		points[2*i+1] = &Point{x: building[1], height: building[2], typ: PointType_End, index: i}
	}
	sort.Slice(points, func(i, j int) bool {
		pi := points[i]
		pj := points[j]
		if pi.typ == PointType_Start && pj.typ == PointType_Start {
			if pi.x < pj.x {
				return true
			}
			if pi.x == pj.x {
				return pi.height >= pj.height
			}
			return false
		}
		if pi.typ == PointType_End && pj.typ == PointType_End {
			if pi.x < pj.x {
				return true
			}
			if pi.x == pj.x {
				return pi.height <= pj.height
			}
			return false
		}
		if pi.x < pj.x {
			return true
		}
		if pi.x == pj.x {
			return pi.typ == PointType_Start
		}
		return false
	})

	// fmt.Println("sorted points")
	// for _, point := range points {
	// 	fmt.Println(point)
	// }

	marked2RemoveMap := make(map[int]bool)

	result := make([][]int, 0)
	prevMax := &Point{height: 0}
	h := &PointHeap{}
	heap.Init(h)
	for _, point := range points {
		if point.typ == PointType_Start {
			// if start of building, should Push point to heap and check
			heap.Push(h, point)
			currentMax := h.Peek().(*Point)
			// fmt.Println("push point to heap", point, currentMax, prevMax)
			// h.Print()
			if currentMax.height > prevMax.height {
				// if max height changes, the new point should be added
				result = append(result, []int{point.x, currentMax.height})
				// and replace prevMax point
				prevMax = currentMax
			}
		}
		if point.typ == PointType_End {
			// if end of building, should Remove point from heap and check

			// h.Print()
			currentMax := h.Peek().(*Point)
			// fmt.Println("end point, currentHeapRoot", currentMax, point, prevMax)
			if currentMax.height == point.height {
				// heapRoot equals current point,then pop it out.
				// keep pop if next root is marked to remove
				// fmt.Println("end height matchesd heap root, remove top", point)
				heap.Pop(h)
				// h.Print()
				for {
					p := h.Peek()
					if p != nil && marked2RemoveMap[p.(*Point).index] == true {
						// fmt.Println("heap root is marked to remove", p, marked2RemoveMap)
						heap.Pop(h)
						// h.Print()
					} else {
						break
					}
				}
				// fmt.Println("after pop root,should check heaproot changed")
				// h.Print()
				if root := h.Peek(); root == nil {
					// fmt.Println("heap root changed to nil, add height 0 to result")
					result = append(result, []int{point.x, 0})
					prevMax = &Point{height: 0}
				} else {
					if prevMax.height != root.(*Point).height {
						// fmt.Println("heap root changed, add to result", prevMax, root)
						prevMax = root.(*Point)
						result = append(result, []int{point.x, prevMax.height})
					}
				}
			} else {
				// heapRoot not each current Point, marked to remove
				// Heap's removal of non-root element is O(n)
				marked2RemoveMap[point.index] = true
			}
		}
	}

	return result
}

type Point struct {
	x      int
	height int
	typ    PointType // start or end
	index  int       // in origin building index
}

type PointType string

const (
	PointType_Start PointType = "start"
	PointType_End   PointType = "end"
)

type PointHeap []*Point

func (h PointHeap) Len() int {
	return len(h)
}

func (h PointHeap) Less(i, j int) bool {
	pi := h[i]
	pj := h[j]
	return pi.height > pj.height
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(p interface{}) {
	*h = append(*h, p.(*Point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *PointHeap) Peek() interface{} {
	if h.Len() > 0 {
		return (*h)[0]
	}
	return nil
}

func (h *PointHeap) Print() {
	fmt.Println("print heap...")
	for _, e := range *h {
		fmt.Println(e)
	}
}

// func main() {

// 	result := getSkyline([][]int{
// 		[]int{2, 9, 10},
// 		[]int{3, 7, 15},
// 		[]int{5, 12, 12},
// 		[]int{15, 20, 10},
// 		[]int{19, 24, 8},
// 	})
// 	fmt.Println(result)

// 	result = getSkyline([][]int{
// 		[]int{0, 2, 3},
// 		[]int{2, 5, 3},
// 	})
// 	fmt.Println(result)

// result := getSkyline([][]int{
// 	{5, 10, 12},
// 	{10, 15, 7},
// })
// fmt.Println(result)

// // debug heap
// // h := &PointHeap{}
// // heap.Init(h)
// // for _, point := range []*Point{
// // 	&Point{x: 2, height: 10, typ: PointType_Start},
// // 	&Point{x: 9, height: 10, typ: PointType_End},
// // 	&Point{x: 3, height: 15, typ: PointType_Start},
// // 	&Point{x: 7, height: 15, typ: PointType_End},
// // 	&Point{x: 7, height: 8, typ: PointType_Start},
// // } {
// // 	heap.Push(h, point)
// // }
// // fmt.Println("peek", (*h)[0])
// // fmt.Println("pop", heap.Pop(h))
// // for h.Len() > 0 {
// // 	fmt.Println(heap.Pop(h))
// // }
// }

// @lc code=end
