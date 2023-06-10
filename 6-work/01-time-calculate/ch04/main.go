package main

import (
	"fmt"
	"time"
)

type ServiceTime struct {
	Name      string
	StartTime string
	EndTime   string
}

func main() {
	serviceTimes := []ServiceTime{
		{"周一", "00:00:00", "23:59:59"},
		{"周二", "09:00:00", "18:00:00"},
		{"周三", "09:00:00", "18:00:00"},
		{"周四", "00:00:00", "23:59:59"},
		{"周五", "00:00:00", "23:59:59"},
		{"周六", "00:00:00", "23:59:59"},
		{"周日", "00:00:00", "23:59:59"},
	}

	var availableTimes []string
	start := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(0, 1, 2, 0, 0, 0, 0, time.UTC)

	for _, st := range serviceTimes {
		startTime, _ := time.Parse("15:04:05", st.StartTime)
		endTime, _ := time.Parse("15:04:05", st.EndTime)

		// 如果开始时间大于结束时间，表示跨过了一天，需要分两段进行计算
		if startTime.After(endTime) {
			startTime2, _ := time.Parse("15:04:05", "00:00:00")
			endTime2, _ := time.Parse("15:04:05", st.EndTime)

			availableTimes = append(availableTimes, getTimeRange(start, startTime2)...)
			availableTimes = append(availableTimes, getTimeRange(endTime2, end)...)
		} else {
			availableTimes = append(availableTimes, getTimeRange(startTime, endTime)...)
		}
	}

	for _, t := range availableTimes {
		fmt.Println(t)
	}
}

func getTimeRange(startTime time.Time, endTime time.Time) []string {
	var times []string

	for !startTime.Equal(endTime) {
		times = append(times, startTime.Format("15:04:05"))
		startTime = startTime.Add(time.Second)
	}

	return times
}
