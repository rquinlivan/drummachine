# Drum machine

## Requirements

Requires go version `go1.13.8`.

## Build it

```
cd path/to/drummachine
go build
```

## Run it

`./drummachine [pattern_file] [measure_count]`

The pattern file path provided must be a valid JSON format. See 
examples in the `patterns` directory for the structure.

Pattern files define instruments and the structure of the pattern,
as well as BPM and name (metadata).

###  Sample patterns

`patterns/four_on_the_floor.json`

Four on the floor pattern.

`patterns/we_will_rock_you.json`

A classic Queen jock jam!

## Test it

`go test`

## Developer's Guide

### Pattern file JSON schema

A pattern file must include the following fields:

- `Name` (`string`):  the name of this pattern. (Metadata)
- `Bpm` (`integer`): the tempo to play at.
- `Instruments` (`map`): a map of instrument name to instrument definition.
- `Patterns` (`map`): a map of position to an array of instrument names. Defines the drum pattern. Keys used in the array must correspond to keys in the instruments map.

Instrument definition must include the following fields:
- `Name`: a *unique* name. Used as a key  in the instrument map definition.
- `Symbol`: a *unique* symbol. May be used as a key in the `Player`.

### Implementing a player

The project can be extended through implementing three functions:

#### Player

Provides a function that play an instrument. In the sample `ConsolePlayer` this entails printing to console. 
This is the main interface for implementation.

#### Rest

Provides a function to implement a rest. In audio `Player`s, this will likely be a no-op.

#### Measure

Provides a function to implement an end of a measure. In audio `Player`s, this will likely be a no-op.

