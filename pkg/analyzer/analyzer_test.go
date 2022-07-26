package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestUseStdlibVars(t *testing.T) {
	pkgs := []string{
		// "a/crypto",
		"a/http",
		// "a/time",
	}
	analysistest.Run(t, analysistest.TestData(), New(), pkgs...)
}
