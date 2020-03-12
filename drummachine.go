package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(5*time.Second)
	fmt.Println("Hi")

	drumPattern := Drums {name:"Killin it", bpm: 128}
	play(drumPattern)
}

type Instrument struct {
	name string
	symbol string
	sound string // Optional
}

type Drums struct {
	name string
	bpm int
}

// Play a drum pattern
func play(drum Drums) {
	delayMillis := int(drum.bpm / 60 * 1000)
	time.Sleep(delayMillis * time.Millisecond)
}
