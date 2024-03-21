/*
 * @lc app=leetcode id=1723 lang=golang
 *
 * [1723] Find Minimum Time to Finish All Jobs
 *
 * https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs/description/
 *
 * algorithms
 * Hard (43.77%)
 * Likes:    232
 * Dislikes: 11
 * Total Accepted:    8.2K
 * Total Submissions: 18.7K
 * Testcase Example:  '[3,2,3]\n3'
 *
 * You are given an integer array jobs, where jobs[i] is the amount of time it
 * takes to complete the i^th job.
 *
 * There are k workers that you can assign jobs to. Each job should be assigned
 * to exactly one worker. The working time of a worker is the sum of the time
 * it takes to complete all jobs assigned to them. Your goal is to devise an
 * optimal assignment such that the maximum working time of any worker is
 * minimized.
 *
 * Return the minimum possible maximum working time of any assignment.
 *
 *
 * Example 1:
 *
 *
 * Input: jobs = [3,2,3], k = 3
 * Output: 3
 * Explanation: By assigning each person one job, the maximum time is 3.
 *
 *
 * Example 2:
 *
 *
 * Input: jobs = [1,2,4,7,8], k = 2
 * Output: 11
 * Explanation: Assign the jobs the following way:
 * Worker 1: 1, 2, 8 (working time = 1 + 2 + 8 = 11)
 * Worker 2: 4, 7 (working time = 4 + 7 = 11)
 * The maximum working time is 11.
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= jobs.length <= 12
 * 1 <= jobs[i] <= 10^7
 *
 *
 */

// @lc code=start

import "sort"

func minimumTimeRequired(jobs []int, k int) int {
	/*
		结果落在一个区间里：jobs 里最大值(即 k==len(jobs)和 jobs 之和(k==1)
		用二分法找到这个值
	*/

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i] > jobs[j]
	})
	max := jobs[0]
	sum := 0
	for _, job := range jobs {
		sum += job
	}

	// 二分
	l, r := max, sum
	for l < r {
		mid := l + (r-l)/2
		workers := make([]int, k)
		if backtrack(jobs, 0, workers, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// 检查 jobs 分配到 k 组里，并且上限是 limit，能否满足
func backtrack(jobs []int, jobIdx int, workers []int, limit int) bool {
	//fmt.Println("bt", jobs, jobIdx, workers, limit)
	// capture
	//fmt.Println("bt", jobs, idx, workers, *ans)
	if jobIdx == len(jobs) { // 所有的工作都分配了
		return true
	}
	job := jobs[jobIdx]
	// 尝试把任务分配给每个人，看分配完了之后能不能满足条件
	for i := range workers {
		if workers[i]+job <= limit { // 可以分配给工人 i
			// choose
			workers[i] += job
			// explore
			if backtrack(jobs, jobIdx+1, workers, limit) {
				return true
			}
			// unchoose
			workers[i] -= job

			// constrain
			// 走到这里，说明把 jobIdx 分配给第 i 个工人是不行的
			// 如果unchoose 之后，第 i 个工人的工作量是0，说明分配给其他工作量 >= 0 的工人也不行，即不需要再继续检查了
			if workers[i] == 0 {
				break
			}
		}
	}
	return false
}

// @lc code=end

