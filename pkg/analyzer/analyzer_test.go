package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestReuseStdlibVars(t *testing.T) {
	pkgs := []string{
		"_http",
		"_time",
	}
	analysistest.Run(t, analysistest.TestData(), New(), pkgs...)
}
