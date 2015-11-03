package Periods

import (
	"time"
)

type MonthPeriod struct {
	code uint32
}

func NewMonthPeriod(code uint32) *MonthPeriod {
	return &MonthPeriod{code}
}

func (p MonthPeriod) Month() uint32 {
	return (p.code % 100)
}

func (p MonthPeriod) Year() uint32 {
	return (p.code / 100)
}

func (p MonthPeriod) MonthInt() int32 {
	return int32(p.code % 100)
}

func (p MonthPeriod) YearInt() int32 {
	return int32(p.code / 100)
}

func (p MonthPeriod) Description() string {
	var datePeriod = time.Date(int(p.Year()), time.Month(p.Month()), 1, 0, 0, 0, 0, time.UTC)
	return datePeriod.Format("January 2006")
}

func (s MonthPeriod) Equals(other MonthPeriod) bool {
	return (s.code == other.code)
}

func (s MonthPeriod) CompareTo(other MonthPeriod) int {
	return (int(s.code) - int(other.code))
}

func (s MonthPeriod) OpGt(other MonthPeriod) bool {
	return (s.CompareTo(other) > 0)
}

func (s MonthPeriod) OpLt(other MonthPeriod) bool {
	return (s.CompareTo(other) < 0)
}

