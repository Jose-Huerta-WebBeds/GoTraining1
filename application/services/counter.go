package services

import (
	"sync/atomic"
	"time"
)

var counter int64 = 0

//Counter returns a number that will increase in 1 by 1 from 0.
//It's concurrently safe
func Counter() int {
	time.Sleep(time.Second)
	return int(atomic.AddInt64(&counter, 1))
}
