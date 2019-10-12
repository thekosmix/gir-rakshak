package utils

import (
	"time"
)

func NowAsUnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}
