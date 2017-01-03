package gigasecond //main // gigasecond

import (
	//"fmt"
	"time"
)

const testVersion = 4

const gigasecond = 1000000000

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	//fmt.Println("input :", t, time.Second*gigasecond)
	return t.Add(time.Second * gigasecond)
}

/*
func main() {
	in, err := time.Parse("2006-01-02", "2011-04-25")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(AddGigasecond(in))
}
*/
