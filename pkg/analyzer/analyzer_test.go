package analyzer_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/sashamelentyev/usestdlibvars/pkg/analyzer"
)

var flags = []string{
	analyzer.TimeWeekdayFlag,
	analyzer.TimeMonthFlag,
	analyzer.TimeLayoutFlag,
	analyzer.CryptoHashFlag,
	analyzer.RPCDefaultPathFlag,
	analyzer.SQLIsolationLevelFlag,
	analyzer.TLSSignatureSchemeFlag,
	analyzer.ConstantKindFlag,
	analyzer.TimeDateMonthFlag,
}

func TestUseStdlibVars(t *testing.T) {
	a := analyzer.New()

	for _, flag := range flags {
		if err := a.Flags.Set(flag, "true"); err != nil {
			t.Fatal(err)
		}
	}

	testCases := []struct {
		dir string
	}{
		{dir: "a/crypto"},
		{dir: "a/http"},
		{dir: "a/rpc"},
		{dir: "a/time"},
		{dir: "a/sql"},
		{dir: "a/tls"},
		{dir: "a/constant"},
	}

	for _, test := range testCases {
		t.Run(test.dir, func(t *testing.T) {
			analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), a, test.dir)
		})
	}
}
