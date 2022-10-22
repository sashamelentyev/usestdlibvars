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
	analyzer.OSDevNullFlag,
	analyzer.SQLIsolationLevelFlag,
	analyzer.TLSSignatureSchemeFlag,
	analyzer.ConstantKindFlag,
}

var pkgs = []string{
	"a/crypto",
	"a/http",
	"a/rpc",
	"a/time",
	"a/os",
	"a/sql",
	"a/tls",
	"a/constant",
}

func TestUseStdlibVars(t *testing.T) {
	a := analyzer.New()

	t.Log(flags)

	for _, flag := range flags {
		if err := a.Flags.Set(flag, "true"); err != nil {
			t.Fatal(err)
		}
	}

	analysistest.Run(t, analysistest.TestData(), a, pkgs...)
}
