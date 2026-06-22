package clock

import "time"

type Clock interface {
	Now() time.Time
}

type WallClock struct{}

func Wall() Clock {
	return WallClock{}
}

func (WallClock) Now() time.Time {
	return time.Now()
}

type FixedClock struct {
	Value time.Time
}

func Fixed(value time.Time) Clock {
	return FixedClock{Value: value}
}

func (c FixedClock) Now() time.Time {
	return c.Value
}
