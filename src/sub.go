package main

import "time"

// 获取两个日期相差天数
func TimeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

func main() {
	var activeTimestamp int64 = 1605608500
	sub := TimeSub(time.Now(), time.Unix(activeTimestamp, 0))
	println(sub)
}
