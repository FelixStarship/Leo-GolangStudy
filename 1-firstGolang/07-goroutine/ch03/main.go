package main

import (
	"fmt"
	"time"
)

// 计算时间段的交集
func calculateIntersection(start1, end1, start2, end2 time.Time) (time.Time, time.Time) {
	// 计算交集的起始时间
	intersectionStart := start1
	if start2.After(intersectionStart) {
		intersectionStart = start2
	}

	// 计算交集的结束时间
	intersectionEnd := end1
	if end2.Before(intersectionEnd) {
		intersectionEnd = end2
	}

	// 检查是否存在交集
	if intersectionStart.After(intersectionEnd) {
		return time.Time{}, time.Time{} // 返回空时间表示无交集
	}

	return intersectionStart, intersectionEnd
}

// 计算时间段的差集
func calculateDifference(start1, end1, start2, end2 time.Time) []time.Time {
	var difference []time.Time

	// 检查第一个时间段是否在第二个时间段之前
	if end1.Before(start2) || end1.Equal(start2) {
		difference = append(difference, start1, end1)
		return difference
	}

	// 检查第二个时间段是否在第一个时间段之前
	if end2.Before(start1) || end2.Equal(start1) {
		difference = append(difference, start1, end1)
		return difference
	}

	// 第一个时间段在第二个时间段的前半部分
	if start1.Before(start2) {
		difference = append(difference, start1, start2)
	}

	// 第一个时间段在第二个时间段的后半部分
	if end1.After(end2) {
		difference = append(difference, end2, end1)
	}

	return difference
}

func main() {
	// 创建两个时间段
	time1Start := time.Date(2023, time.May, 1, 0, 0, 0, 0, time.UTC)
	time1End := time.Date(2023, time.May, 10, 0, 0, 0, 0, time.UTC)
	time2Start := time.Date(2023, time.May, 5, 0, 0, 0, 0, time.UTC)
	time2End := time.Date(2023, time.May, 15, 0, 0, 0, 0, time.UTC)

	// 计算交集
	intersectionStart, intersectionEnd := calculateIntersection(time1Start, time1End, time2Start, time2End)
	fmt.Println("Intersection:", intersectionStart, "-", intersectionEnd)

	// 计算差集
	difference := calculateDifference(time1Start, time1End, time2Start, time2End)
	fmt.Println("Difference:", difference)
}
