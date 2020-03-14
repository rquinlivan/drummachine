package main

import (
	"fmt"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)

	drumPattern := readFromFile("four_on_the_floor")
	play(drumPattern)
}

type Instrument struct {
	name   string
	symbol string
}

type DrumPattern struct {
	name        string
	bpm         int
	instruments []Instrument
	patterns    map[int][]string
}

// Given a pattern name, return the DrumPattern.
// This is read from the file at ./patterns/[patternName].json
// TODO: read file
func readFromFile(patternName string) DrumPattern {
	instruments := []Instrument {
		Instrument{
			name:   "",
			symbol: "",
		},
	}
	patterns := map[int][]string {
		1: []string {"hi_hat"},
	}
	return DrumPattern{
		name: "Four on the floor",
		bpm:  120,
		instruments: instruments,
		patterns: patterns,
	}
}

// Play a drum pattern
func play(drum DrumPattern) {
	fmt.Println("Playing pattern '", drum.name, "' at", drum.bpm, "beats per minute")
	measures := 10
	delay := time.Duration(int(drum.bpm / 60 * 1000)) * time.Millisecond
	for i := 0; i < measures; i++ {
		for j := 1; j <= 16; j++ {
			time.Sleep(delay)
			fmt.Print("_")
		}
		fmt.Println("")
	}
}
