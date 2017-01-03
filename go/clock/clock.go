package clock

import (
	"fmt"
)

const testVersion = 4

type Clock int

func New(hour, minute int) Clock {
	return Clock(0).Add((hour * 60) + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

func (c Clock) Add(minutes int) Clock {
	return Clock(((c+Clock(minutes))%1440 + 1440) % 1440)
}
