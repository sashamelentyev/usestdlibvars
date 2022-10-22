//go:build !windows && !plan9

package analyzer_test

import "github.com/sashamelentyev/usestdlibvars/pkg/analyzer"

func init() {
	flags = append(flags, analyzer.SyslogPriorityFlag)
	pkgs = append(pkgs, "a/syslog")
}
