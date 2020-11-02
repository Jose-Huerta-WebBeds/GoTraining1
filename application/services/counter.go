package services

import (
	"sync/atomic"
)

var counter int64 = 0

//Counter returns a number that will increase in 1 by 1 from 0.
//It's concurrently safe
func Counter() int {

	return int(atomic.AddInt64(&counter, 1))
}
