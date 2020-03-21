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
	drumsB := DrumPattern{
		Name: "test",
		Bpm:  100000,
		Instruments: map[string]Instrument{
			"foo": {Name: "foo", Symbol: "?"},
			"bar": {Name: "bar", Symbol: "("},
		},
		Patterns: map[int][]string{
			1: {"foo", "bar"},
			2: {"foo", "bar"},
			3: {"foo", "bar"},
			4: {"foo", "bar"},
			5: {"foo", "bar"},
			6: {"foo", "bar"},
			7: {"foo", "bar"},
			8: {"foo", "bar"},
			9: {"foo", "bar"},
			10: {"foo", "bar"},
			11: {"foo", "bar"},
			12: {"foo", "bar"},
			13: {"foo", "bar"},
			14: {"foo", "bar"},
			15: {"foo", "bar"},
			16: {"foo", "bar"},
		},
	}
	drumsC := DrumPattern{
		Name: "test",
		Bpm:  100000,
		Instruments: map[string]Instrument{
			"foo": {Name: "foo", Symbol: "?"},
			"bar": {Name: "bar", Symbol: "("},
			"baz": {Name: "baz", Symbol: ")"},
		},
		Patterns: map[int][]string{
			1: {"foo", "bar", "baz"},
			2: {"foo", "bar", "baz"},
			3: {"foo", "bar", "baz"},
			4: {"foo", "bar", "baz"},
			5: {"foo", "bar", "baz"},
			6: {"foo", "bar", "baz"},
			7: {"foo", "bar", "baz"},
			8: {"foo", "bar", "baz"},
			9: {"foo", "bar", "baz"},
			10: {"foo", "bar", "baz"},
			11: {"foo", "bar", "baz"},
			12: {"foo", "bar", "baz"},
			13: {"foo", "bar", "baz"},
			14: {"foo", "bar", "baz"},
			15: {"foo", "bar", "baz"},
			16: {"foo", "bar", "baz"},
		},
	}
	drumsD := DrumPattern{
		Name: "test",
		Bpm:  100000,
		Instruments: map[string]Instrument{
			"foo": {Name: "foo", Symbol: "?"},
			"bar": {Name: "bar", Symbol: "("},
			"baz": {Name: "baz", Symbol: ")"},
		},
		Patterns: map[int][]string{
			1: {"foo"},
		},
	}
	badDrums := DrumPattern{
		Name: "test",
		Bpm:  100000,
		Instruments: map[string]Instrument{
			"foo": {Name: "foo", Symbol: "?"},
			"bar": {Name: "bar", Symbol: "("},
			"baz": {Name: "baz", Symbol: ")"},
		},
		Patterns: map[int][]string{
			1: {"qux"},
		},
	}
	_TestPlayConfiguration(t, emptyDrums,1, 0, 16, 0)
	_TestPlayConfiguration(t, drumsA, 1, 16, 0, 0)
	_TestPlayConfiguration(t, drumsB, 1, 32, 0, 0)
	_TestPlayConfiguration(t, drumsC, 1, 48, 0, 0)
	_TestPlayConfiguration(t, drumsD, 10, 10, 150, 9)
	_TestPlayConfiguration(t, badDrums, 1, 0, 16, 0)
	_TestPlayConfiguration(t, badDrums, 10, 0, 160, 9)
}

// Test a configured DrumPattern
// Params:
// - measures: number of measures to play
// - playCount: asserted play count
// - restCount: asserted rest count
// - measureCount: asserted measure break count
func _TestPlayConfiguration(t *testing.T, pattern DrumPattern, measures int, playCount int, restCount int, measureCount int) {
	restCalls := 0
	measureCalls := 0
	playCalls := 0
	Play(pattern, measures, func(instrument Instrument) {
		playCalls++
	}, func() {
		restCalls++
	}, func() {
		measureCalls++
	})

	if restCalls != restCount ||
		measureCalls != measureCount ||
		playCalls != playCount {
		t.Error("Actual:     restCalls ==", restCalls, "measureCalls ==", measureCalls, "playCount ==", playCalls)
		t.Fatal("Expected:   restCalls ==", restCount, "measureCalls ==", measureCount, "playCount ==", playCount)
	}

}
