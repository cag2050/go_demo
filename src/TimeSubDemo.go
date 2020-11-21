package main

import "time"

// 获取两个日期相差天数
func TimeSub(t1, t2 time.Time) int {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}

func main() {
	// 设备激活时间戳
	var devActTime int64 = 1537977600
	// 设备激活天数
	devActDays := TimeSub(time.Now(), time.Unix(devActTime, 0))
	println(devActDays)
}
