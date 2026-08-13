package to

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	for _, v := range []time.Duration{0, time.Second, -5 * time.Hour} {
		p := DurationP(v)
		if *p != v {
			t.Errorf("DurationP(%v) points at %v", v, *p)
		}
		if got := Duration(p); got != v {
			t.Errorf("Duration(DurationP(%v)) == %v", v, got)
		}
	}
}

func TestInt(t *testing.T) {
	for _, v := range []int{0, 42, -7} {
		p := IntP(v)
		if *p != v {
			t.Errorf("IntP(%v) points at %v", v, *p)
		}
		if got := Int(p); got != v {
			t.Errorf("Int(IntP(%v)) == %v", v, got)
		}
	}
}

func TestInt64(t *testing.T) {
	for _, v := range []int64{0, 42, -7, 1 << 40} {
		p := Int64P(v)
		if *p != v {
			t.Errorf("Int64P(%v) points at %v", v, *p)
		}
		if got := Int64(p); got != v {
			t.Errorf("Int64(Int64P(%v)) == %v", v, got)
		}
	}
}

func TestInt32(t *testing.T) {
	for _, v := range []int32{0, 42, -7, 1 << 30} {
		p := Int32P(v)
		if *p != v {
			t.Errorf("Int32P(%v) points at %v", v, *p)
		}
		if got := Int32(p); got != v {
			t.Errorf("Int32(Int32P(%v)) == %v", v, got)
		}
	}
}

func TestFloat64(t *testing.T) {
	for _, v := range []float64{0, 3.5, -0.25} {
		p := Float64P(v)
		if *p != v {
			t.Errorf("Float64P(%v) points at %v", v, *p)
		}
		if got := Float64(p); got != v {
			t.Errorf("Float64(Float64P(%v)) == %v", v, got)
		}
	}
}

func TestFloat32(t *testing.T) {
	for _, v := range []float32{0, 3.5, -0.25} {
		p := Float32P(v)
		if *p != v {
			t.Errorf("Float32P(%v) points at %v", v, *p)
		}
		if got := Float32(p); got != v {
			t.Errorf("Float32(Float32P(%v)) == %v", v, got)
		}
	}
}

func TestString(t *testing.T) {
	for _, v := range []string{"", "giantswarm"} {
		p := StringP(v)
		if *p != v {
			t.Errorf("StringP(%q) points at %q", v, *p)
		}
		if got := String(p); got != v {
			t.Errorf("String(StringP(%q)) == %q", v, got)
		}
	}
}

func TestBool(t *testing.T) {
	for _, v := range []bool{true, false} {
		p := BoolP(v)
		if *p != v {
			t.Errorf("BoolP(%v) points at %v", v, *p)
		}
		if got := Bool(p); got != v {
			t.Errorf("Bool(BoolP(%v)) == %v", v, got)
		}
	}
}

// Each XP returns a distinct pointer, so callers can take addresses of
// literals in a loop without aliasing.
func TestPointersAreDistinct(t *testing.T) {
	if IntP(1) == IntP(1) {
		t.Error("IntP returned the same pointer twice")
	}
}
