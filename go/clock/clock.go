package clock // main //clock //

import (
	"fmt"
)

const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	c := Clock{hour: hour, minute: minute}
	standardize(&c)
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	standardize(&c)
	return c
}

func standardize(c *Clock) {
	rolloverHours := c.minute / 60
	//fmt.Println("rolloverHours=", rolloverHours)
	c.minute %= 60
	if c.minute < 0 {
		rolloverHours--
		c.minute += 60
	}
	c.hour += rolloverHours
	c.hour %= 24
	if c.hour < 0 {
		c.hour += 24
	}
}

/*
func main() {

	c := Clock{1, -60}
	standardize(&c)
	fmt.Println(c.String())

}
*/
