/*
 * @lc app=leetcode.cn id=853 lang=golang
 *
 * [853] Car Fleet
 *
 * https://leetcode.cn/problems/car-fleet/description/
 *
 * algorithms
 * Medium (42.49%)
 * Likes:    177
 * Dislikes: 0
 * Total Accepted:    20.5K
 * Total Submissions: 48.2K
 * Testcase Example:  '12\n[10,8,0,5,3]\n[2,4,1,1,3]'
 *
 * There are n cars going to the same destination along a one-lane road. The
 * destination is target miles away.
 *
 * You are given two integer array position and speed, both of length n, where
 * position[i] is the position of the i^th car and speed[i] is the speed of the
 * i^th car (in miles per hour).
 *
 * A car can never pass another car ahead of it, but it can catch up to it and
 * drive bumper to bumper at the same speed. The faster car will slow down to
 * match the slower car's speed. The distance between these two cars is ignored
 * (i.e., they are assumed to have the same position).
 *
 * A car fleet is some non-empty set of cars driving at the same position and
 * same speed. Note that a single car is also a car fleet.
 *
 * If a car catches up to a car fleet right at the destination point, it will
 * still be considered as one car fleet.
 *
 * Return the number of car fleets that will arrive at the destination.
 *
 *
 * Example 1:
 *
 *
 * Input: target = 12, position = [10,8,0,5,3], speed = [2,4,1,1,3]
 * Output: 3
 * Explanation:
 * The cars starting at 10 (speed 2) and 8 (speed 4) become a fleet, meeting
 * each other at 12.
 * The car starting at 0 does not catch up to any other car, so it is a fleet
 * by itself.
 * The cars starting at 5 (speed 1) and 3 (speed 3) become a fleet, meeting
 * each other at 6. The fleet moves at speed 1 until it reaches target.
 * Note that no other cars meet these fleets before the destination, so the
 * answer is 3.
 *
 *
 * Example 2:
 *
 *
 * Input: target = 10, position = [3], speed = [3]
 * Output: 1
 * Explanation: There is only one car, hence there is only one fleet.
 *
 *
 * Example 3:
 *
 *
 * Input: target = 100, position = [0,2,4], speed = [4,2,1]
 * Output: 1
 * Explanation:
 * The cars starting at 0 (speed 4) and 2 (speed 2) become a fleet, meeting
 * each other at 4. The fleet moves at speed 2.
 * Then, the fleet (speed 2) and the car starting at 4 (speed 1) become one
 * fleet, meeting each other at 6. The fleet moves at speed 1 until it reaches
 * target.
 *
 *
 *
 * Constraints:
 *
 *
 * n == position.length == speed.length
 * 1 <= n <= 10^5
 * 0 < target <= 10^6
 * 0 <= position[i] < target
 * All the values of position are unique.
 * 0 < speed[i] <= 10^6
 *
 *
 sort the car by position(far to near), then by speed(fast to slow):
 if a car is farer and faster than another, then they cannot merge

start from nearest car, check wether it can catch previous car and merge into one fleet
 if it can catch, consider the fleet as a single car
 if it can not catch, put into stack
*/

package main

import (
	"fmt"
	"sort"
)

// @lc code=start

type fleet struct {
	pos   int
	speed int
}

func canMerge(f1, f2 fleet, target int) bool {
	if f1.speed == f2.speed {
		return f1.pos == f2.pos
	}

	// all positions are unique

	// f1.pos + f1.speed * x = f2.pos + f2.speed * x
	// x = (f2.pos - f1.pos) / (f1.speed - f2.speed)
	time2catch := float32((f2.pos - f1.pos)) / float32((f1.speed - f2.speed))
	// fmt.Println("catch", f1, f2, time2catch)
	if time2catch < 0 {
		return false
	}
	return float32(f2.pos)+time2catch*float32(f2.speed) <= float32(target)
}

func carFleet(target int, positions []int, speeds []int) int {
	fleets := make([]fleet, len(positions))
	for i := 0; i < len(positions); i++ {
		fleets[i] = fleet{pos: positions[i], speed: speeds[i]}
	}
	sort.Slice(fleets, func(i, j int) bool {
		fi := fleets[i]
		fj := fleets[j]
		return fi.pos > fj.pos // not consider same position&& fi.speed > fj.speed
	})
	stack := make([]fleet, 0)
	// fmt.Println(fleets)
	for _, fleet := range fleets {
		// fmt.Println("stack", fleet, stack)
		if len(stack) <= 0 {
			stack = append(stack, fleet)
			continue
		}

		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if canMerge(fleet, top, target) {
				break
			} else {
				stack = append(stack, fleet)
			}
		}
	}
	return len(stack)

}

// 0 2 4
// 2 3 1

// @lc code=end

func main() {
	// fmt.Println(carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
	// fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	// fmt.Println(carFleet(10, []int{0, 4, 2}, []int{2, 1, 3}))
	fmt.Println(carFleet(12, []int{4, 0, 5, 3, 1, 2}, []int{6, 10, 9, 6, 7, 2}))
}
