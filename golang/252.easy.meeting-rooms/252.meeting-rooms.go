// intervals = [[0,30],[5,10],[15,20]]

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	//fmt.Println(intervals)
	maxEnd := 0
	for _, interval := range intervals {
		s, e := interval[0], interval[1]
		if s < maxEnd {
			return false
		}
		maxEnd = e
	}
	return true
}