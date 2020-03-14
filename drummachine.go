package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("####		 DRUM MACHINE		 ###")
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
	instruments map[string]Instrument
	patterns    map[int][]string
}

// Given a pattern name, return the DrumPattern.
// This is read from the file at ./patterns/[patternName].json
// TODO: read file
func readFromFile(patternName string) DrumPattern {
	instruments := map[string]Instrument{
		"hi_hat": {
			name:   "hi_hat",
			symbol: "^",
		},
		"bass_drum": {
			name:   "bass_drum",
			symbol: "&",
		},
		"snare_drum": {
			name:   "snare_drum",
			symbol: "*",
		},
	}
	patterns := map[int][]string{
		1:  {"bass_drum"},
		3:  {"hi_hat"},
		5:  {"snare_drum", "bass_drum"},
		7:  {"hi_hat"},
		9:  {"bass_drum"},
		11: {"hi_hat"},
		13: {"snare_drum","bass_drum"},
		15: {"hi_hat"},
	}
	return DrumPattern{
		name:        "Four on the floor",
		bpm:         120,
		instruments: instruments,
		patterns:    patterns,
	}
}

func getDelay(bpm int) time.Duration {
	beatsPerSec := float32(bpm) / 60
	delayInSec := 1 / beatsPerSec
	delayInMillis := int(delayInSec * 1000.0)
	return time.Duration(delayInMillis) * time.Millisecond
}

// Play a drum pattern
func play(drum DrumPattern) {
	fmt.Println("Playing pattern '", drum.name, "' at", drum.bpm, "beats per minute")
	measures := 10
	delay := getDelay(drum.bpm)
	fmt.Println("Delay is ", delay)
	for i := 0; i < measures; i++ {
		for j := 1; j <= 16; j++ {
			time.Sleep(delay)
			patterns, present := drum.patterns[j]
			switch present {
			case true:
				for _, patternKey := range patterns {
					inst, instExists := drum.instruments[patternKey]
					if instExists {
						fmt.Print(inst.symbol)
					} else {
						fmt.Print("?")
					}
				}
			case false:
				fmt.Print("_")
			}
		}
		fmt.Println("")
	}
}
