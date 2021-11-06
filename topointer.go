package gowxpay

import "time"

func Time(t time.Time) *time.Time {
	return &t
}

func String(s string) *string {
	return &s
}

func Bool(b bool) *bool {
	return &b
}

func Float64(f float64) *float64 {
	return &f
}

func Float32(f float32) *float32 {
	return &f
}

func Int64(i int64) *int64 {
	return &i
}

func Int32(i int32) *int32 {
	return &i
}

