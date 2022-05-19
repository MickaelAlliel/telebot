package utils

import "time"

func ToTimePtr(t time.Time) *time.Time {
	return &t
}
