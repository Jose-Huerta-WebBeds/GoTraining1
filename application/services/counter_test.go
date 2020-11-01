package services

import "testing"

func TestCounter(t *testing.T) {
	//we reset the counter just in case this test is executed after another.
	counter = 0
	for expected := 1; expected < 100; expected++ {
		actual := Counter()
		if actual != expected {
			t.Errorf("The %dth execution of Counter has returned %d.", expected, actual)
		}
	}
}
