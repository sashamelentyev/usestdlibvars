package analyzer_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/sashamelentyev/usestdlibvars/pkg/analyzer"
)

func TestUseStdlibVars(t *testing.T) {
	pkgs := []string{
		"a/crypto",
		"a/http",
		"a/rpc",
		"a/time",
		"a/os",
	}

	a := analyzer.New()

	if err := a.Flags.Set(analyzer.TimeWeekdayFlag, "true"); err != nil {
		t.Error(err)
	}
	if err := a.Flags.Set(analyzer.TimeMonthFlag, "true"); err != nil {
		t.Error(err)
	}
	if err := a.Flags.Set(analyzer.TimeLayoutFlag, "true"); err != nil {
		t.Error(err)
	}
	if err := a.Flags.Set(analyzer.CryptoHashFlag, "true"); err != nil {
		t.Error(err)
	}
	if err := a.Flags.Set(analyzer.DefaultRPCPathFlag, "true"); err != nil {
		t.Error(err)
	}
	if err := a.Flags.Set(analyzer.OSDevNullFlag, "true"); err != nil {
		t.Error(err)
	}

	analysistest.Run(t, analysistest.TestData(), a, pkgs...)
}
