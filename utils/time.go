package utils

import "time"

/**
 * 时间的格式化操作
 */
const TIME_FORMAR_ONE = "2006年01月2日 15:04:05"
const TIME_FORMAR_TWO = "2006.01.02 15:04:05"
const TIME_FORMAR_THREE = "2006-01-02 15:04:05"
const TIME_FORMAR_FOUR = "2006/01/02 15:04:05"

func TimeFormat(t int64, format string) string {
	return time.Unix(t, 0).Format(format)
}
