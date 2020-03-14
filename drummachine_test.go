package main

import (
	"testing"
	"time"
)

func TestGetDelay(t *testing.T) {
	expectations := map[int]time.Duration {
		128: 468  * time.Millisecond,
		120: 500  * time.Millisecond,
		60:  1000 * time.Millisecond,
	}
	for bpm, dur := range expectations {
		actual := GetDelay(bpm)
		if actual != dur {
			t.Error("bpm ", bpm, "should equate to ", dur, "but was", actual)
		}
	}
}

func TestPlay(t *testing.T) {

}
