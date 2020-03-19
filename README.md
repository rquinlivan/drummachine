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
