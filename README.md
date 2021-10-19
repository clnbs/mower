Mower for Jellysmack
===

For this project, I wanted to keep things really simple so I did not design a complex parallel execution process with a worker manager. Movement processing for each mower are done in routines but step by step.

## Tests
Tests are partials and do not cover all corner cases, for example : given two mowers A{1, 1, N} and B{2, 2, W}, if they move forward in the same round, they'll be on the same case

To run the tests, simply run `go test ./... -v` in the root directory of this project

## Usage
You can run it easily without compiling it if you have go installed with `go run cmd/mower/main.go <path to file>`