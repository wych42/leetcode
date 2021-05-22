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

func minimumTimeRequired(jobs []int, k int) int {
	/*
		结果取决于工作时间最长的那个工人

		对于每一个工作，可以选择交给其中的一个工人，然后继续计算，直到工作量为0; 再回头，看把工作交给下一个人会怎么样
	*/
	ans := 1<<31 - 1
	backtrack(jobs, 0, make([]int, k), 0, &ans)
	return ans
}

func backtrack(jobs []int, idx int, workers []int, max int, ans *int) {
	// capture
	//fmt.Println("bt", jobs, idx, workers, *ans)
	if idx == len(jobs) { // 所有的工作都分配了
		// find max in workers
		//fmt.Println("before", workers)
		if max < *ans {
			*ans = max
		}
		//fmt.Println("after", workers)
		return
	}
	job := jobs[idx]
	for i, d := range workers {
		total := d + job
		// constrain
		if total > *ans {
			continue
		}
		temp := max
		// choose
		workers[i] = total
		if total > max {
			max = total
		}
		// explore
		backtrack(jobs, idx+1, workers, max, ans)
		// unchoose
		workers[i] = d
		max = temp
	}
}

// @lc code=end

