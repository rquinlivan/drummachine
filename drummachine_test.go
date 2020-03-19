package main

import (
	"testing"
	"time"
)

func TestGetDelay(t *testing.T) {
	expectations := map[int]time.Duration{
		128: 468 * time.Millisecond,
		120: 500 * time.Millisecond,
		60:  1000 * time.Millisecond,
	}
	for bpm, dur := range expectations {
		actual := GetDelay(bpm)
		if actual != dur {
			t.Error("Bpm ", bpm, "should equate to ", dur, "but was", actual)
		}
	}
}

func TestPlayEmptyDrums(t *testing.T) {
	emptyDrums := DrumPattern{
		Name: "test",
		Bpm:  100000,
		Instruments: map[string]Instrument{
			"foo": {Name: "foo", Symbol: "?"},
		},
	}
	drumsA := DrumPattern{
		Name: "test",
		Bpm:  100000,
		Instruments: map[string]Instrument{
			"foo": {Name: "foo", Symbol: "?"},
		},
		Patterns: map[int][]string{
			1: {"foo"},
			2: {"foo"},
			3: {"foo"},
			4: {"foo"},
			5: {"foo"},
			6: {"foo"},
			7: {"foo"},
			8: {"foo"},
			9: {"foo"},
			10: {"foo"},
			11: {"foo"},
			12: {"foo"},
			13: {"foo"},
			14: {"foo"},
			15: {"foo"},
			16: {"foo"},
		},
	}
	drumsAPattern := "????????????????"
	_TestPlayConfiguration(t, drumsA, 1, drumsAPattern, 0, 0)
	_TestPlayConfiguration(t, emptyDrums,1, "", 160, 0)
}

func _TestPlayConfiguration(t *testing.T, pattern DrumPattern, measures int, out string, rest int, measure int) {
	output := ""
	restCount := 0
	measureCount := 0
	Play(pattern, measures, func(instrument Instrument) {
		output += instrument.Symbol
	}, func() {
		output += " "
		restCount++
	}, func() {
		output += "END"
		measureCount++
	})

	if restCount != rest ||
		measureCount != measure ||
		output != out {
		t.Error("Actual:     restCount ==", restCount, "measureCount ==", measureCount, "output ==", output)
		t.Fatal("Expected:   restCount ==", rest, "measureCount ==", measure, "output ==", out)
	}

}
