package gst

// #include "gst.go.h"
import "C"
import "time"

// Clock is a go wrapper around a GstClock.
type Clock struct{ *Object }

// Instance returns the underlying GstClock instance.
func (c *Clock) Instance() *C.GstClock { return C.toGstClock(c.unsafe()) }

// IsSynced returns true if the clock is synced.
func (c *Clock) IsSynced() bool { return gobool(C.gst_clock_is_synced(c.Instance())) }

// Time gets the current time of the given clock in nanoseconds or ClockTimeNone if invalid.
// The time is always monotonically increasing and adjusted according to the current offset and rate.
func (c *Clock) Time() uint64 {
	res := C.gst_clock_get_time(c.Instance())
	if uint64(res) == ClockTimeNone {
		return ClockTimeNone
	}
	return uint64(res)
}

// InternalTime gets the current internal time of the given clock in nanoseconds
// or ClockTimeNone if invalid. The time is returned unadjusted for the offset and the rate.
func (c *Clock) InternalTime() uint64 {
	res := C.gst_clock_get_internal_time(c.Instance())
	if uint64(res) == ClockTimeNone {
		return ClockTimeNone
	}
	return uint64(res)
}

// Duration returns the time.Duration equivalent of this clock time.
func (c *Clock) Duration() time.Duration {
	tm := c.Time()
	if tm == ClockTimeNone {
		return time.Duration(-1)
	}
	return nanosecondsToDuration(tm)
}

// InternalDuration returns the time.Duration equivalent of this clock's internal time.
func (c *Clock) InternalDuration() time.Duration {
	tm := c.InternalTime()
	if tm == ClockTimeNone {
		return time.Duration(-1)
	}
	return nanosecondsToDuration(tm)
}

// String returns the string representation of this clock value.
func (c *Clock) String() string { return c.Duration().String() }

// InternalString returns the string representation of this clock's internal value.
func (c *Clock) InternalString() string { return c.InternalDuration().String() }

func nanosecondsToDuration(n uint64) time.Duration { return time.Duration(n) * time.Nanosecond }
