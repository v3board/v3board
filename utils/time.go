package utils

import "github.com/golang-module/carbon/v2"

// TodayStartTime 获取今天的开始时间
func TodayStartTime() int64 {
	return carbon.Now().StartOfDay().Timestamp()
}

// YesterdayStartTime 获取昨天的开始时间
func YesterdayStartTime() int64 {
	return carbon.Yesterday().StartOfDay().Timestamp()
}
