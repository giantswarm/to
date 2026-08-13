[![CircleCI](https://dl.circleci.com/status-badge/img/gh/giantswarm/to/tree/main.svg?style=shield)](https://dl.circleci.com/status-badge/redirect/gh/giantswarm/to/tree/main)
[![Go Reference](https://pkg.go.dev/badge/github.com/giantswarm/to.svg)](https://pkg.go.dev/github.com/giantswarm/to)

# to

Library for convenient type dereferencing in Go.

```go
import "github.com/giantswarm/to"

spec := Spec{
    Enabled: to.BoolP(true),   // value -> pointer
    Retries: to.IntP(3),
}

enabled := to.Bool(spec.Enabled) // pointer -> value
```

The `X(*T) T` functions dereference and will panic on a nil pointer. The
`XP(T) *T` functions return the address of a fresh copy, so they are safe to
call in a loop.
