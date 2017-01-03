package clock

import (
	"fmt"
)

const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	return Clock{hour: hour, minute: minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%v:%v", c.hour, c.minute)
}

func (Clock) Add(minutes int) Clock {
	return Clock{}
}
