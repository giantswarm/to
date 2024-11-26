// Package go provides functions to return the value of a pointer,
// and vice versa.
package to

import "time"

func Duration(d *time.Duration) time.Duration {
	return *d
}

func DurationP(d time.Duration) *time.Duration {
	return &d
}

func Int(i *int) int {
	return *i
}

func IntP(i int) *int {
	return &i
}

func Int64(i *int64) int64 {
	return *i
}

func Int64P(i int64) *int64 {
	return &i
}

func Int32(i *int32) int32 {
	return *i
}

func Int32P(i int32) *int32 {
	return &i
}

func Float64(i *float64) float64 {
	return *i
}

func Float64P(i float64) *float64 {
	return &i
}

func Float32(i *float32) float32 {
	return *i
}

func Float32P(i float32) *float32 {
	return &i
}

func String(s *string) string {
	return *s
}

func StringP(s string) *string {
	return &s
}

func Bool(b *bool) bool {
	return *b
}

func BoolP(b bool) *bool {
	return &b
}
