# usestdlibvars

A linter that detect the possibility to use variables/constants from the Go standard library.

## Install

```bash
go install github.com/sashamelentyev/usestdlibvars
```

## Usage

```console
$ usestdlibvars -h
usestdlibvars: A linter that detect the possibility to use variables/constants from the Go standard library.

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
        suggest the use of crypto.Hash
  -debug string
        debug flags, any subset of "fpstv"
  -default-rpc-path
        suggest the use of rpc.DefaultXXPath
  -fix
        apply all suggested fixes
  -flags
        print analyzer flags in JSON
  -http-method
        suggest the use of http.MethodXX (default true)
  -http-no-body
        suggest the use of http.NoBody
  -http-status-code
        suggest the use of http.StatusXX (default true)
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
        suggest the use of time.Layout
  -time-month
        suggest the use of time.Month
  -time-weekday
        suggest the use of time.Weekday
  -trace string
        write trace log to this file
  -v    no effect (deprecated)
```

## Examples

```go
package response

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// JSON marshals 'v' to JSON, automatically escaping HTML and setting the
// Content-Type as application/json.
func JSON(w http.ResponseWriter, statusCode int, v any) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(buf.Bytes()); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
```

```bash
usestdlibvars ./...
```

```console
response.go:16:30: "500" can be replaced by http.StatusInternalServerError
response.go:22:30: "500" can be replaced by http.StatusInternalServerError
```

## Sponsors

[<img src="https://evrone.com/logo/evrone-sponsored-logo.png">](https://evrone.com/?utm_source=usestdlibvars)
