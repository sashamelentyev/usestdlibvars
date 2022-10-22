//go:build unix || (js && wasm)

package analyzer_test

import "github.com/sashamelentyev/usestdlibvars/pkg/analyzer"

func init() {
	flags = append(flags, analyzer.OSDevNullFlag)
	pkgs = append(pkgs, "a/os")
}
