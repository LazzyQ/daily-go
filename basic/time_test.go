package basic

import (
	"testing"
	"time"
)

const (
	TIME_LAYOUT = "2006-01-02 15:04:05"
)

func TestTimeParse(t *testing.T) {
	now := time.Now()
	t.Logf("now: %v", now)

	formatNow := now.Format(TIME_LAYOUT)
	t.Logf("formatNow: %v", formatNow)

	parseTime, _ := time.Parse(TIME_LAYOUT, formatNow)
	t.Logf("parseTime: %v", parseTime)

	l, _ := time.LoadLocation("Asia/Shanghai")
	parseLoctionTime, _ := time.ParseInLocation(TIME_LAYOUT, formatNow, l)
	t.Logf("parseLoctionTime: %v", parseLoctionTime)
}
