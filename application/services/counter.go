package services

import (
	"log"
	"sync/atomic"
)

var counter int64 = 0

//Counter returns a number that will increase in 1 by 1 from 0.
//It's concurrently safe
func Counter() int {
	log.Println("Counter called")
	return int(atomic.AddInt64(&counter, 1))
}
