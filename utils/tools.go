package utils

import "time"

// 计算两个日期之间的时间差
func TimeSub(t1, t2 string) int {
	tf, _ := time.Parse("2006-01-02 15:04:05", t1)
	ts, _ := time.Parse("2006-01-02 15:04:05", t2)
	return int(tf.Sub(ts).Seconds())
}
