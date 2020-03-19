package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	fmt.Println("####		 DRUM MACHINE		 ###")
	drumPattern := ReadFromFile("we_will_rock_you")
	Play(drumPattern, 10, ConsolePlayer, ConsoleRest, ConsoleMeasure)
}

// Represents an instrument. E.g., snare drum, cymbals, etc.
// Provide a unique Symbol and unique Name.
type Instrument struct {
	Name   string
	Symbol string
}

// Represents a pattern to be played, encapsulating both the set of Instruments
// to be applied as well as which Instruments to play at each beat.
// Must also provide a Name (merely metadata) and a Bpm setting.
type DrumPattern struct {
	Name        string
	Bpm         int
	Instruments map[string]Instrument
	Patterns    map[int][]string
}

// Given a pattern Name, return the DrumPattern.
// This is read from the file at ./patterns/[patternName].json
func ReadFromFile(patternName string) DrumPattern {
	bytes, err := ioutil.ReadFile("patterns/" + patternName + ".json")
	if err != nil {
		fmt.Println(err)
		panic("Can't read file " + patternName + ".json")
	}

	var drumPattern DrumPattern
	err = json.Unmarshal(bytes, &drumPattern)
	if err != nil {
		panic("Can't deserialize file " + patternName + ".json")
	}
	return drumPattern
}

func GetDelay(bpm int) time.Duration {
	beatsPerSec := float32(bpm) / 60
	delayInSec := 1 / beatsPerSec
	delayInMillis := int(delayInSec * 1000.0)
	return time.Duration(delayInMillis) * time.Millisecond
}

// Player types

// Instructions for how to "play" an instrument.
type Player func(instrument Instrument)

// Instructions for how a "play" a rest.
type Rest func()

// Instructions for how to "play" the end of a measure.
// Note: this should likely be a no-op, but is a useful hook for
// testing and console printing.
type Measure func()


// Console player

// Write instrument symbols to console.
func ConsolePlayer(instrument Instrument) {
	fmt.Print(instrument.Symbol)
}

// Write rest to console.
func ConsoleRest() {
	fmt.Print("_")
}

// Write end of measure to console.
func ConsoleMeasure() {
	fmt.Println("")
}

/* Play a drum pattern, defined in `drum DrumPattern`, for `measures` measures.
 * The `Play` function itself just Instruments the playing, it doesn't define
 * player behavior.
 *
 * That behavior is determined by injected functions for easy implemenation!
 * To implement your own player, provide these three functions (defined above):
 * - `Player` itself defines playing an `Instrument`.
 * - `Rest` defines behavior if an `Instrument` is *not* played.
 * - `Measure` defines optional behavior when the end of a measure is reached.
 */
func Play(drum DrumPattern, measures int, player Player, rest Rest, measure Measure) {
	fmt.Println("Playing pattern '", drum.Name, "' at", drum.Bpm, "beats per minute")
	delay := GetDelay(drum.Bpm)
	for i := 0; i < measures; i++ {
		for j := 1; j <= 16; j++ {
			time.Sleep(delay)
			patterns, present := drum.Patterns[j]
			switch present {
			case true:
				for _, patternKey := range patterns {
					inst, instExists := drum.Instruments[patternKey]
					if instExists {
						player(inst)
					} else {
						rest()
					}
				}
			case false:
				rest()
			}
		}
		if i != measures - 1 {
			measure()
		}
	}
}
