// Package go provides functions to return the value of a pointer.
//
// Deprecated: Use idiomatic Go instead
package to

import "time"

// Deprecated: Use idiomatic Go instead
func Duration(d *time.Duration) time.Duration {
	return *d
}

// Deprecated: Use idiomatic Go instead
func DurationP(d time.Duration) *time.Duration {
	return &d
}

// Deprecated: Use idiomatic Go instead
func Int(i *int) int {
	return *i
}

// Deprecated: Use idiomatic Go instead
func IntP(i int) *int {
	return &i
}

// Deprecated: Use idiomatic Go instead
func Int64(i *int64) int64 {
	return *i
}

// Deprecated: Use idiomatic Go instead
func Int64P(i int64) *int64 {
	return &i
}

// Deprecated: Use idiomatic Go instead
func Int32(i *int32) int32 {
	return *i
}

// Deprecated: Use idiomatic Go instead
func Int32P(i int32) *int32 {
	return &i
}

// Deprecated: Use idiomatic Go instead
func Float64(i *float64) float64 {
	return *i
}

// Deprecated: Use idiomatic Go instead
func Float64P(i float64) *float64 {
	return &i
}

// Deprecated: Use idiomatic Go instead
func Float32(i *float32) float32 {
	return *i
}

// Deprecated: Use idiomatic Go instead
func Float32P(i float32) *float32 {
	return &i
}

// Deprecated: Use idiomatic Go instead
func String(s *string) string {
	return *s
}

// Deprecated: Use idiomatic Go instead
func StringP(s string) *string {
	return &s
}

// Deprecated: Use idiomatic Go instead
func Bool(b *bool) bool {
	return *b
}

// Deprecated: Use idiomatic Go instead
func BoolP(b bool) *bool {
	return &b
}
