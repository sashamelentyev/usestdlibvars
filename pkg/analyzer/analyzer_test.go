package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestUseStdlibVars(t *testing.T) {
	pkgs := []string{
		"a/crypto",
		"a/http",
		"a/time",
	}

	analyzer := New()
	_ = analyzer.Flags.Set("time-weekday", "true")
	_ = analyzer.Flags.Set("time-month", "true")
	_ = analyzer.Flags.Set("time-layout", "true")
	_ = analyzer.Flags.Set("crypto-hash", "true")

	analysistest.Run(t, analysistest.TestData(), analyzer, pkgs...)
}
