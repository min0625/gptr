package gptr

import "time"

func NewTime(v time.Time) *time.Time {
	return &v
}

func ToTime(p *time.Time, d time.Time) time.Time {
	if p != nil {
		return *p
	}
	return d
}
