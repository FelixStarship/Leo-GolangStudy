package main

import (
	"fmt"
	"time"
)

type TimeInterval struct {
	Day     time.Weekday
	Start   time.Time
	End     time.Time
	IsWork  bool
	Minutes int
}

func main() {
	// 定义工作时间段
	workIntervals := []TimeInterval{
		{time.Monday, parseTime("2023-05-29 09:00"), parseTime("2023-05-29 18:00"), true, 0},
		{time.Tuesday, parseTime("2023-05-30 20:00"), parseTime("2023-05-30 21:00"), true, 0},
		{time.Wednesday, parseTime("2023-05-31 09:00"), parseTime("2023-05-31 18:00"), true, 0},

		{time.Thursday, parseTime("2023-06-01 06:00"), parseTime("2023-06-01 08:00"), true, 0},

		{time.Friday, parseTime("2023-06-02 09:00"), parseTime("2023-06-02 18:00"), true, 0},
		{time.Saturday, parseTime("2023-06-03 20:00"), parseTime("2023-06-03 21:00"), true, 0},
		{time.Sunday, parseTime("2023-06-04 20:00"), parseTime("2023-06-04 21:00"), true, 0},
	}

	// 计算每个工作时间段的分钟数
	for i := range workIntervals {
		workIntervals[i].Minutes = int(workIntervals[i].End.Sub(workIntervals[i].Start).Minutes())
	}

	// 假设告警发生时间为非工作时间
	alertTime := parseTime("2023-06-01 19:30")

	// 判断告警发生时间是否在工作时间段内
	isWorkTime := false
	var currentWorkInterval TimeInterval
	for _, interval := range workIntervals {
		if alertTime.Weekday() == interval.Day && isTimeBetween(alertTime, interval.Start, interval.End) {
			isWorkTime = true
			currentWorkInterval = interval
			break
		}
	}

	// 计算超时时间
	timeoutDuration := 2 * time.Hour
	var timeoutTime time.Time
	if isWorkTime {
		// 计算告警超时时间，扣除非工作时间
		remainingMinutes := currentWorkInterval.Minutes - int(alertTime.Sub(currentWorkInterval.Start).Minutes())
		if remainingMinutes <= 0 {
			timeoutTime = alertTime.Add(timeoutDuration)
		} else {
			timeoutTime = alertTime.Add(time.Duration(remainingMinutes) * time.Minute)
		}
	} else {
		// 寻找下一个工作时间段
		nextWorkInterval := getNextWorkInterval(alertTime, workIntervals)
		if nextWorkInterval != nil {
			if nextWorkInterval.Start.After(alertTime) {
				// 计算告警超时时间，顺延到下一个工作时间段开始并加上2小时
				timeoutTime = nextWorkInterval.Start.Add(timeoutDuration)
			} else if nextWorkInterval.End.Before(alertTime) {
				// 计算告警超时时间，扣除工作时间
				remainingMinutes := nextWorkInterval.Minutes + int(alertTime.Sub(nextWorkInterval.End).Minutes())
				timeoutTime = alertTime.Add(time.Duration(remainingMinutes) * time.Minute)
			} else {
				// 在工作时间段内，直接将告警超时时间设置为工作时间段结束时间
				timeoutTime = nextWorkInterval.End
			}
		} else {
			// 没有下一个工作时间段，无法计算超时时间
			timeoutTime = time.Time{}
		}
	}

	// 打印超时时间
	if timeoutTime.IsZero() {
		fmt.Println("无法计算超时时间")
	} else {
		fmt.Println("告警超时时间:", timeoutTime.Format("2006-01-02 15:04:05"))
	}
}

// 解析时间字符串
func parseTime(timeStr string) time.Time {
	layout := "2006-01-02 15:04"
	t, _ := time.Parse(layout, timeStr)
	return t
}

// 判断时间是否在给定的时间范围内
func isTimeBetween(t, start, end time.Time) bool {
	return t.After(start) && t.Before(end)
}

// 获取告警发生时间之后的下一个工作时间段
func getNextWorkInterval(alertTime time.Time, workIntervals []TimeInterval) *TimeInterval {
	nextDay := alertTime.Weekday()
	nextTime := alertTime
	for {
		nextDay = (nextDay + 1) % 7
		nextTime = nextTime.AddDate(0, 0, 1)
		for _, interval := range workIntervals {
			if nextDay == interval.Day {
				nextTime = time.Date(nextTime.Year(), nextTime.Month(), nextTime.Day(), interval.Start.Hour(), interval.Start.Minute(), 0, 0, nextTime.Location())
				return &TimeInterval{
					Day:     nextDay,
					Start:   interval.Start,
					End:     interval.End,
					IsWork:  interval.IsWork,
					Minutes: interval.Minutes,
				}
			}
		}
	}
}
