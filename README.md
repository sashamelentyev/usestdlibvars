# usestdlibvars

A linter that detect the possibility to use variables/constants from the Go standard library.

## Install

```
go install github.com/sashamelentyev/usestdlibvars
```

## Usage

```console
$ usestdlibvars -h                                                     
usestdlibvars: Detect the possibility to use constants/variables from the stdlib.

Usage: usestdlibvars [-flag] [package]


Flags:
  -V    print version and exit
  -all
        no effect (deprecated)
  -c int
        display offending line with this many lines of context (default -1)
  -cpuprofile string
        write CPU profile to this file
  -crypto-hash
        
  -debug string
        debug flags, any subset of "fpstv"
  -fix
        apply all suggested fixes
  -flags
        print analyzer flags in JSON
  -http-method
         (default true)
  -http-status-code
         (default true)
  -json
        emit JSON output
  -memprofile string
        write memory profile to this file
  -source
        no effect (deprecated)
  -tags string
        no effect (deprecated)
  -test
        indicates whether test files should be analyzed, too (default true)
  -time-layout
        
  -time-month
        
  -time-weekday
        
  -trace string
        write trace log to this file
  -v    no effect (deprecated)

```

## Examples

```bash
usestdlibvars ./...
```
