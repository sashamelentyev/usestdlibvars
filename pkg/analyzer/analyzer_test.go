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
		"a/sql",
		"a/tls",
		"a/constant",
	}

	a := analyzer.New()

	mustNil(t, a.Flags.Set(analyzer.TimeWeekdayFlag, "true"))
	mustNil(t, a.Flags.Set(analyzer.TimeMonthFlag, "true"))
	mustNil(t, a.Flags.Set(analyzer.TimeLayoutFlag, "true"))
	mustNil(t, a.Flags.Set(analyzer.CryptoHashFlag, "true"))
	mustNil(t, a.Flags.Set(analyzer.RPCDefaultPathFlag, "true"))
	mustNil(t, a.Flags.Set(analyzer.OSDevNullFlag, "true"))
	mustNil(t, a.Flags.Set(analyzer.SQLIsolationLevelFlag, "true"))
	mustNil(t, a.Flags.Set(analyzer.TLSSignatureSchemeFlag, "true"))
	mustNil(t, a.Flags.Set(analyzer.ConstantKindFlag, "true"))

	analysistest.Run(t, analysistest.TestData(), a, pkgs...)
}

func mustNil(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error(err)
	}
}
