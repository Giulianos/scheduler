package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Giulianos/scheduler"
)

func StringToCronRule(rs string) (scheduler.CronRule, error) {
	r := scheduler.CronRule{}
	var weekday, month, day, hour, minute string
	_, err := fmt.Sscanf(rs, "%s %s %s %s %s", &weekday, &month, &day, &hour, &minute)
	if err != nil {
		return scheduler.CronRule{}, err
	}

	if weekday == "*" {
		r.Weekday = scheduler.EveryWeekday
	} else {
		v, err := strconv.ParseInt(weekday, 10, 64)
		if err != nil {
			return scheduler.CronRule{}, err
		}

		r.Weekday = time.Weekday(v)
	}

	if month == "*" {
		r.Month = scheduler.EveryMonth
	} else {
		v, err := strconv.ParseInt(month, 10, 64)
		if err != nil {
			return scheduler.CronRule{}, err
		}

		r.Month = time.Month(v)
	}

	if day == "*" {
		r.Day = scheduler.EveryDay
	} else {
		v, err := strconv.ParseInt(day, 10, 64)
		if err != nil {
			return scheduler.CronRule{}, err
		}

		r.Day = int(v)
	}

	if hour == "*" {
		r.Hour = scheduler.EveryHour
	} else {
		v, err := strconv.ParseInt(hour, 10, 64)
		if err != nil {
			return scheduler.CronRule{}, err
		}

		r.Hour = int(v)
	}

	if minute == "*" {
		r.Minute = scheduler.EveryMinute
	} else {
		v, err := strconv.ParseInt(minute, 10, 64)
		if err != nil {
			return scheduler.CronRule{}, err
		}

		r.Minute = int(v)
	}

	return r, nil

}
