package utils

import (
	"testing"

	"github.com/Giulianos/scheduler"
)

var cases = map[string]scheduler.CronRule{
	"4 * 10 * 45":  {4, scheduler.EveryMonth, 10, scheduler.EveryHour, 45},
	"6 5 10 12 45": {6, 5, 10, 12, 45},
	"* 4 * 10 *":   {scheduler.EveryWeekday, 4, scheduler.EveryDay, 10, scheduler.EveryMinute},
}

func TestStringToCronRule(t *testing.T) {
	for i, expected := range cases {
		got, err := StringToCronRule(i)
		if err != nil {
			t.Errorf("test for %s failed: %s", i, err)
		}
		if got != expected {
			t.Errorf("test for %s failed", i)
		}
	}
}
