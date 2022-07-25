package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestUseStdlibVars(t *testing.T) {
	pkgs := []string{
		"_crypto",
		"_http",
		"_time",
	}
	analysistest.Run(t, analysistest.TestData(), New(), pkgs...)
}
