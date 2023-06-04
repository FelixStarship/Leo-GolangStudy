package main

import (
	"fmt"
	"time"
)

type TimeInterval struct {
	Day    time.Weekday
	Start  time.Time
	End    time.Time
	IsWork bool
}

func main() {
	//  定义工作时间段
	workIntervals := []TimeInterval{
		{time.Monday, parseTime("2023-05-29  09:00"), parseTime("2023-05-29  18:00"), true},
		{time.Tuesday, parseTime("2023-05-30  20:00"), parseTime("2023-05-30  21:00"), true},
		{time.Wednesday, parseTime("2023-05-31  09:00"), parseTime("2023-05-31  18:00"), true},
		{time.Thursday, parseTime("2023-06-01  20:00"), parseTime("2023-06-01  21:00"), true},
		{time.Friday, parseTime("2023-06-02  09:00"), parseTime("2023-06-02  18:00"), true},
		{time.Saturday, parseTime("2023-06-03  20:00"), parseTime("2023-06-03  21:00"), true},
		{time.Sunday, parseTime("2023-06-04  20:00"), parseTime("2023-06-04  21:00"), true},
	}

	//  假设告警发生时间为非工作时间
	alertTime := parseTime("2023-06-01  05:00")

	//  计算超时时间
	timeoutDuration := 2 * time.Hour

	timeoutTime := calculateTimeoutTime(alertTime, workIntervals, timeoutDuration)

	//  打印超时时间
	if !timeoutTime.IsZero() {
		fmt.Println("告警超时时间:", timeoutTime.Format("2006-01-02  15:04:05"))
	} else {
		fmt.Println("无法计算超时时间")
	}
}

// 解析时间字符串
func parseTime(timeStr string) time.Time {
	layout := "2006-01-02  15:04"
	t, _ := time.Parse(layout, timeStr)
	return t
}

// 判断时间是否在给定的时间范围内
func isTimeBetween(t, start, end time.Time) bool {
	return t.After(start) && t.Before(end)
}

func getCurrentWorkInterval(alertTime time.Time, workIntervals []TimeInterval) TimeInterval {
	for _, interval := range workIntervals {
		// 判断当前是否是工作时间
		if alertTime.Weekday() == interval.Day {
			return interval
		}
	}
	return TimeInterval{}
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
					Day:    nextDay,
					Start:  interval.Start,
					End:    interval.End,
					IsWork: interval.IsWork,
				}
			}
		}
	}
}

// 计算告警超时时间
func calculateTimeoutTime(alertTime time.Time, workIntervals []TimeInterval, timeoutDuration time.Duration) time.Time {
	isWorkTime := false
	var currentWorkInterval TimeInterval
	//  判断告警发生时间是否在工作时间段内
	for _, interval := range workIntervals {
		if alertTime.Weekday() == interval.Day && isTimeBetween(alertTime, interval.Start, interval.End) {
			isWorkTime = true
			currentWorkInterval = interval
			break
		}
	}
	if isWorkTime {
		//在工作时间段内
		if alertTime.Add(timeoutDuration).After(currentWorkInterval.End) {
			//超过工作时间，需要顺延到下一个工作时间段开始
			nextWorkInterval := getNextWorkInterval(alertTime, workIntervals)
			if nextWorkInterval != nil {
				//计算非工作时间：最新的工作开始时间-上一段工作结束时间
				remainingMinutes := nextWorkInterval.Start.Sub(currentWorkInterval.End)
				//告警超时时间为下一个工作时间段开始时间加上剩余时间和2小时
				return alertTime.Add(remainingMinutes).Add(timeoutDuration)
			} else {
				return time.Time{}
			}
		} else {
			//未超过工作时间，直接加上2小时
			return alertTime.Add(timeoutDuration)
		}
	} else {
		// 获取当前工作时间
		currentWorkInterval = getCurrentWorkInterval(alertTime, workIntervals)
		//寻找下一个工作时间段
		nextWorkInterval := getNextWorkInterval(alertTime, workIntervals)
		// 告警产生的时间段整个都是非工作时间
		if !currentWorkInterval.IsWork {

		} else {
			if nextWorkInterval != nil {
				// 告警发生时间+2h判断是否在当前工作时间段内
				if alertTime.Add(timeoutDuration).After(currentWorkInterval.Start) {
					// 告警发生在上一段非工作时间内，【下一段工作时间-当前工作时间段=非工作时间段】
					nextRemaining := nextWorkInterval.Start.Sub(currentWorkInterval.End)
					// 当前工作时间段-告警发生时间=非工作时间（上一段）
					previousRemaining := currentWorkInterval.Start.Sub(alertTime)
					timeoutTime := alertTime.Add(timeoutDuration).Add(nextRemaining).Add(previousRemaining)
					return timeoutTime
				} else {
					// 告警发生在当前非工作时间段内
					// 当前工作开始时间-告警发生时间=非工作时间
					currentNonWorking := currentWorkInterval.Start.Sub(alertTime)
					// 下一段工作开始时间-当前工作结束时间
					remainingNonWorking := nextWorkInterval.Start.Sub(currentWorkInterval.End)
					timeoutTime := alertTime.Add(timeoutDuration).Add(remainingNonWorking).Add(currentNonWorking)
					return timeoutTime
				}
			} else {
				//没有下一个工作时间段，无法计算超时时间
				return time.Time{}
			}
		}
		return time.Time{}
	}
}
